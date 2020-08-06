// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = new(DataSources)
		(*in).DeepCopyInto(*out)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(Permissions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.LastConfigured.DeepCopyInto(&out.LastConfigured)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
	if in.DefaultArtifact != nil {
		in, out := &in.DefaultArtifact, &out.DefaultArtifact
		*out = new(MatchArtifact)
		(*in).DeepCopyInto(*out)
	}
	in.MatchArtifact.DeepCopyInto(&out.MatchArtifact)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactReference) DeepCopyInto(out *ArtifactReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactReference.
func (in *ArtifactReference) DeepCopy() *ArtifactReference {
	if in == nil {
		return nil
	}
	out := new(ArtifactReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BakeManifest) DeepCopyInto(out *BakeManifest) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	if in.FailPipeline != nil {
		in, out := &in.FailPipeline, &out.FailPipeline
		*out = new(bool)
		**out = **in
	}
	if in.ContinuePipeline != nil {
		in, out := &in.ContinuePipeline, &out.ContinuePipeline
		*out = new(bool)
		**out = **in
	}
	if in.CompleteOtherBranchesThenFail != nil {
		in, out := &in.CompleteOtherBranchesThenFail, &out.CompleteOtherBranchesThenFail
		*out = new(bool)
		**out = **in
	}
	if in.ExpectedArtifacts != nil {
		in, out := &in.ExpectedArtifacts, &out.ExpectedArtifacts
		*out = make([]Artifact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputArtifacts != nil {
		in, out := &in.InputArtifacts, &out.InputArtifacts
		*out = make([]*ArtifactReference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArtifactReference)
				**out = **in
			}
		}
	}
	out.InputArtifact = in.InputArtifact
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BakeManifest.
func (in *BakeManifest) DeepCopy() *BakeManifest {
	if in == nil {
		return nil
	}
	out := new(BakeManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckPreconditions) DeepCopyInto(out *CheckPreconditions) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	if in.Preconditions != nil {
		in, out := &in.Preconditions, &out.Preconditions
		*out = make([]*Precondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Precondition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckPreconditions.
func (in *CheckPreconditions) DeepCopy() *CheckPreconditions {
	if in == nil {
		return nil
	}
	out := new(CheckPreconditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomArtifact) DeepCopyInto(out *CustomArtifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomArtifact.
func (in *CustomArtifact) DeepCopy() *CustomArtifact {
	if in == nil {
		return nil
	}
	out := new(CustomArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSources) DeepCopyInto(out *DataSources) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new([]DataSource)
		if **in != nil {
			in, out := *in, *out
			*out = make([]DataSource, len(*in))
			copy(*out, *in)
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new([]DataSource)
		if **in != nil {
			in, out := *in, *out
			*out = make([]DataSource, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSources.
func (in *DataSources) DeepCopy() *DataSources {
	if in == nil {
		return nil
	}
	out := new(DataSources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteManifest) DeepCopyInto(out *DeleteManifest) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(Options)
		(*in).DeepCopyInto(*out)
	}
	if in.Kinds != nil {
		in, out := &in.Kinds, &out.Kinds
		*out = make([]KubernetesKind, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteManifest.
func (in *DeleteManifest) DeepCopy() *DeleteManifest {
	if in == nil {
		return nil
	}
	out := new(DeleteManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployManifest) DeepCopyInto(out *DeployManifest) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Moniker = in.Moniker
	if in.ManifestArtifact != nil {
		in, out := &in.ManifestArtifact, &out.ManifestArtifact
		*out = new(MatchArtifact)
		(*in).DeepCopyInto(*out)
	}
	if in.RequiredArtifacts != nil {
		in, out := &in.RequiredArtifacts, &out.RequiredArtifacts
		*out = make([]Artifact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequiredArtifactIds != nil {
		in, out := &in.RequiredArtifactIds, &out.RequiredArtifactIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TrafficManagement.DeepCopyInto(&out.TrafficManagement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployManifest.
func (in *DeployManifest) DeepCopy() *DeployManifest {
	if in == nil {
		return nil
	}
	out := new(DeployManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedArtifact) DeepCopyInto(out *EmbeddedArtifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedArtifact.
func (in *EmbeddedArtifact) DeepCopy() *EmbeddedArtifact {
	if in == nil {
		return nil
	}
	out := new(EmbeddedArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrNameUndefined) DeepCopyInto(out *ErrNameUndefined) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrNameUndefined.
func (in *ErrNameUndefined) DeepCopy() *ErrNameUndefined {
	if in == nil {
		return nil
	}
	out := new(ErrNameUndefined)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FindArtifactsFromResource) DeepCopyInto(out *FindArtifactsFromResource) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FindArtifactsFromResource.
func (in *FindArtifactsFromResource) DeepCopy() *FindArtifactsFromResource {
	if in == nil {
		return nil
	}
	out := new(FindArtifactsFromResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jitter) DeepCopyInto(out *Jitter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jitter.
func (in *Jitter) DeepCopy() *Jitter {
	if in == nil {
		return nil
	}
	out := new(Jitter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JudgmentInput) DeepCopyInto(out *JudgmentInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JudgmentInput.
func (in *JudgmentInput) DeepCopy() *JudgmentInput {
	if in == nil {
		return nil
	}
	out := new(JudgmentInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JudgmentMessage) DeepCopyInto(out *JudgmentMessage) {
	*out = *in
	if in.ManualJudgmentContinue != nil {
		in, out := &in.ManualJudgmentContinue, &out.ManualJudgmentContinue
		*out = new(JudgmentMessageValue)
		**out = **in
	}
	if in.ManualJudgmentStop != nil {
		in, out := &in.ManualJudgmentStop, &out.ManualJudgmentStop
		*out = new(JudgmentMessageValue)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JudgmentMessage.
func (in *JudgmentMessage) DeepCopy() *JudgmentMessage {
	if in == nil {
		return nil
	}
	out := new(JudgmentMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JudgmentMessageValue) DeepCopyInto(out *JudgmentMessageValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JudgmentMessageValue.
func (in *JudgmentMessageValue) DeepCopy() *JudgmentMessageValue {
	if in == nil {
		return nil
	}
	out := new(JudgmentMessageValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelector) DeepCopyInto(out *LabelSelector) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make([]Selector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelector.
func (in *LabelSelector) DeepCopy() *LabelSelector {
	if in == nil {
		return nil
	}
	out := new(LabelSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualJudgment) DeepCopyInto(out *ManualJudgment) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]*ManualJudgmentNotification, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ManualJudgmentNotification)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.JudgmentInputs != nil {
		in, out := &in.JudgmentInputs, &out.JudgmentInputs
		*out = make([]JudgmentInput, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualJudgment.
func (in *ManualJudgment) DeepCopy() *ManualJudgment {
	if in == nil {
		return nil
	}
	out := new(ManualJudgment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualJudgmentNotification) DeepCopyInto(out *ManualJudgmentNotification) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(JudgmentMessage)
		(*in).DeepCopyInto(*out)
	}
	if in.When != nil {
		in, out := &in.When, &out.When
		*out = new([]JudgmentState)
		if **in != nil {
			in, out := *in, *out
			*out = make([]JudgmentState, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualJudgmentNotification.
func (in *ManualJudgmentNotification) DeepCopy() *ManualJudgmentNotification {
	if in == nil {
		return nil
	}
	out := new(ManualJudgmentNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchArtifact) DeepCopyInto(out *MatchArtifact) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchArtifact.
func (in *MatchArtifact) DeepCopy() *MatchArtifact {
	if in == nil {
		return nil
	}
	out := new(MatchArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchStage) DeepCopyInto(out *MatchStage) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchStage.
func (in *MatchStage) DeepCopy() *MatchStage {
	if in == nil {
		return nil
	}
	out := new(MatchStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Moniker) DeepCopyInto(out *Moniker) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Moniker.
func (in *Moniker) DeepCopy() *Moniker {
	if in == nil {
		return nil
	}
	out := new(Moniker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionValue) DeepCopyInto(out *OptionValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionValue.
func (in *OptionValue) DeepCopy() *OptionValue {
	if in == nil {
		return nil
	}
	out := new(OptionValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	if in.GracePeriodSeconds != nil {
		in, out := &in.GracePeriodSeconds, &out.GracePeriodSeconds
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new([]OptionValue)
		if **in != nil {
			in, out := *in, *out
			*out = make([]OptionValue, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permissions) DeepCopyInto(out *Permissions) {
	*out = *in
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permissions.
func (in *Permissions) DeepCopy() *Permissions {
	if in == nil {
		return nil
	}
	out := new(Permissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	if in.ParameterConfig != nil {
		in, out := &in.ParameterConfig, &out.ParameterConfig
		*out = new([]Parameter)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Parameter, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ExpectedArtifacts != nil {
		in, out := &in.ExpectedArtifacts, &out.ExpectedArtifacts
		*out = new([]Artifact)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Artifact, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]MatchStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = new([]Trigger)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Trigger, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	in.LastConfigured.DeepCopyInto(&out.LastConfigured)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Precondition) DeepCopyInto(out *Precondition) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Precondition.
func (in *Precondition) DeepCopy() *Precondition {
	if in == nil {
		return nil
	}
	out := new(Precondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictedExecutionWindow) DeepCopyInto(out *RestrictedExecutionWindow) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	out.Jitter = in.Jitter
	if in.WhiteList != nil {
		in, out := &in.WhiteList, &out.WhiteList
		*out = make([]WhiteListWindow, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictedExecutionWindow.
func (in *RestrictedExecutionWindow) DeepCopy() *RestrictedExecutionWindow {
	if in == nil {
		return nil
	}
	out := new(RestrictedExecutionWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	if in.RequisiteStageRefIds != nil {
		in, out := &in.RequisiteStageRefIds, &out.RequisiteStageRefIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StageEnabled != nil {
		in, out := &in.StageEnabled, &out.StageEnabled
		*out = new(StageEnabled)
		**out = **in
	}
	in.RestrictedExecutionWindow.DeepCopyInto(&out.RestrictedExecutionWindow)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageEnabled) DeepCopyInto(out *StageEnabled) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageEnabled.
func (in *StageEnabled) DeepCopy() *StageEnabled {
	if in == nil {
		return nil
	}
	out := new(StageEnabled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagement) DeepCopyInto(out *TrafficManagement) {
	*out = *in
	in.TrafficManagementOptions.DeepCopyInto(&out.TrafficManagementOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagement.
func (in *TrafficManagement) DeepCopy() *TrafficManagement {
	if in == nil {
		return nil
	}
	out := new(TrafficManagement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagementOptions) DeepCopyInto(out *TrafficManagementOptions) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagementOptions.
func (in *TrafficManagementOptions) DeepCopy() *TrafficManagementOptions {
	if in == nil {
		return nil
	}
	out := new(TrafficManagementOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UndoRolloutManifest) DeepCopyInto(out *UndoRolloutManifest) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UndoRolloutManifest.
func (in *UndoRolloutManifest) DeepCopy() *UndoRolloutManifest {
	if in == nil {
		return nil
	}
	out := new(UndoRolloutManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	in.Stage.DeepCopyInto(&out.Stage)
	if in.ExpectedArtifacts != nil {
		in, out := &in.ExpectedArtifacts, &out.ExpectedArtifacts
		*out = make([]Artifact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.StageEnabled = in.StageEnabled
	if in.RetryStatusCodes != nil {
		in, out := &in.RetryStatusCodes, &out.RetryStatusCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.FailFastStatusCodes != nil {
		in, out := &in.FailFastStatusCodes, &out.FailFastStatusCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteListWindow) DeepCopyInto(out *WhiteListWindow) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteListWindow.
func (in *WhiteListWindow) DeepCopy() *WhiteListWindow {
	if in == nil {
		return nil
	}
	out := new(WhiteListWindow)
	in.DeepCopyInto(out)
	return out
}
