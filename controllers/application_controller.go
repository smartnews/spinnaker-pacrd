/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/armory/plank"
	"github.com/go-logr/logr"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pacrdv1alpha1 "github.com/armory-io/pacrd/api/v1alpha1"
)

const (
	// ControllerName represents this controller's name.
	ControllerName = "spinnaker-application-controller"
	// FinalizerName is used to mark which objects have been processed for deletion.
	FinalizerName = "app.finalizers.armory.io"
)

// SpinnakerClient negotiates interactions with a Spinnaker cluster.
type SpinnakerClient interface {
	GetApplication(string) (*plank.Application, error)
	CreateApplication(*plank.Application) error
	DeleteApplication(string) error
	GetPipelines(string) ([]plank.Pipeline, error)
	DeletePipeline(plank.Pipeline) error
	UpsertPipeline(plank.Pipeline, string) error
	ResyncFiat() error
	ArmoryEndpointsEnabled() bool
	EnableArmoryEndpoints()
}

// ApplicationReconciler reconciles a Application object
type ApplicationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	SpinnakerClient
}

// +kubebuilder:rbac:groups=pacrd.armory.spinnaker.io,resources=applications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pacrd.armory.spinnaker.io,resources=applications/status,verbs=get;update;patch

// Reconcile Application state within the cluster.
func (r *ApplicationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("application", req.NamespacedName)
	logger.Info("reconciling application")

	// Get a handle on the CRD application under scrutiny
	var app pacrdv1alpha1.Application
	if err := r.Get(ctx, req.NamespacedName, &app); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Determine if we're actively deleting the CRD and make sure we clean up
	// any Spinnaker resources that are related. Finalizers are initialized
	// further down the method.
	// TODO delete seems unreliable at this time; about 50/50 chance that
	// delete loop runs during reconcile of a delete. Potentially related to
	// port-forwarding on my local machine.
	// TODO move to DeleteApplicationFromSpinnaker fn
	if !app.GetObjectMeta().GetDeletionTimestamp().IsZero() {
		finalizers := app.GetObjectMeta().GetFinalizers()

		logger.Info("deleting application")

		if containsString(finalizers, FinalizerName) {
			if err := r.SpinnakerClient.DeleteApplication(app.Name); err != nil {
				return ctrl.Result{}, err
			}

			logger.Info("successfully deleted application from spinnaker")

			// Make sure we clean up the finalizer so k8s knows we're done with
			// this resource.
			app.SetFinalizers(removeString(finalizers, FinalizerName))
			if err := r.Update(context.Background(), &app); err != nil {
				return ctrl.Result{}, err
			}
		}

		logger.Info("done deleting application")
		return ctrl.Result{}, nil
	}

	// First, attempt to find the Spinnaker app.
	// TODO move to GetOrCreateApplicationFromSpinnaker fn
	pApp, spinErr := r.SpinnakerClient.GetApplication(req.Name)
	var ue *plank.ErrUnsupportedStatusCode
	if spinErr != nil {
		// If we didn't find the Spinnaker app, create it.
		if errors.As(spinErr, &ue) && ue.Code == 404 {
			logger.Info("did not find application in Spinnaker. Creating...")
			spinApp := toSpinApplication(app)
			spinErr = r.SpinnakerClient.CreateApplication(&spinApp)
			// If we got an error from Spinnaker we failed to create the app,
			// so requeue the event.
			if spinErr != nil {
				// TODO it's possible this will never succeed, consider backoff
				//      or general failure
				return ctrl.Result{}, spinErr
			}
			logger.Info("created spinnaker app")
			return ctrl.Result{}, nil
		}

		// Otherwise something else broke, so requeue the event.
		// TODO understand this failure case
		logger.Error(spinErr, "what does this do")
		return ctrl.Result{}, spinErr
	}

	// Register our finalizer so that we make sure to delete the Spinnaker app
	// when the CRD is deleted.
	// TODO move to RegisterDeleteFinalizer fn
	finalizers := app.GetObjectMeta().GetFinalizers()
	if !containsString(finalizers, FinalizerName) {
		app.SetFinalizers(append(finalizers, FinalizerName))
		if err := r.Update(context.Background(), &app); err != nil {
			return ctrl.Result{}, err
		}
	}

	app.Status.LastConfigured = v1.Time{Time: time.Now()}
	if err := r.Status().Update(context.Background(), &app); err != nil {
		return ctrl.Result{}, err
	}

	logger.Info(pApp.Name)

	// TODO set URL for application

	logger.Info("done reconciling application")
	return ctrl.Result{}, nil
}

func (r *ApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pacrdv1alpha1.Application{}).
		Complete(r)
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
