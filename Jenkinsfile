pipeline {
  agent any
  stages {
    stage('Git push') {
      steps {
        script {
          if (currentBuild.changeSets.size() > 0) {
            echo 'Git push'
          }
          else {
            echo 'No changes'
          }
        }

      }
    }

    stage('Build') {
      steps {
        sh """
           chmod +x ./create_docker_image.sh
           ./create_docker_image.sh
        """
      }
    }

  }
}