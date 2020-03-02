def repository = 'armory'
def name = 'pacrd'

pipeline {
  agent any

  stages {
   stage('Checkout from SCM') {
      steps {
        checkout scm
      }
   }

   stage('Run tests') {
      steps {
        echo 'lol, what tests (TODO)'
      }
   }

   stage('Build & Publish Docker Image') {
      steps {
        sh 'export VERSION=$(git describe --always --dirty)'
        sh "docker build . -t ${repository}/${name}:\$VERSION"
        sh "docker push ${repository}/${name}:\$VERSION"
      }
   }
  }

  post {
    failure {
      slackSend(
        color:'danger',
        message: "${env.JOB_NAME} failed to build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)",
        channel: '#eng-ico-build-notifications'
      )
    }

    success {
      slackSend(
        color:'good',
        message: "${env.JOB_NAME} was published: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)",
        channel: '#eng-ico-build-notifications'
      )
    }
  }
}
