pipeline {
  agent {
    kubernetes {
      yaml '''
apiVersion: v1
kind: Pod
spec:
  securityContext:
    runAsUser: 0
    runAsGroup: 0
    fsGroup: 0
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:v1.23.2-debug
    command:
    - sleep
    - infinity
    tty: true
  - name: kubectl
    image: bitnami/kubectl:1.32
    command:
    - sleep
    - infinity
    tty: true
  - name: sonar-scanner-cli
    image: sonarsource/sonar-scanner-cli:11.2
    command:
    - sleep
    - infinity
    tty: true
        '''
    }
  }
  stages {
    stage('CI Kaniko Build Image & Push to Harbor') {
      steps {
          container('sonar-scanner-cli') {
              script {
                def containerRegistryHost = "${params.ContainerRegistryHost}"
                def containerRegistryProject = "${params.ContainerRegistryProject}"
                def containerName = "${params.ContainerImageName}"
                // def containerTag = "${env.BUILD_NUMBER}"
                // def containerTag = "${params.ContainerImageTag}"
                def containerTag = "${env.GIT_TAG_NAME}"
                sh("sleep 3000")
              }
          }
      }
    }
  }
}