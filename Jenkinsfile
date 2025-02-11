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
    stage('Code Analysis') {
      steps {
          container('sonar-scanner-cli') {
              script {
                echo "Workspace Path: ${env.WORKSPACE}"
                withSonarQubeEnv('my-sonarqube-server') {
                  echo "SONAR-URL: ${env.SONAR_HOST_URL} " 
                }
              }
          }
      }
    }
  }
}