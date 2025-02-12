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
    image: bitnami/kubectl
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

    environment {
        // # GIT
        GIT_TAG_NAME = "dev-gong-v0.0.1"
        //GIT_TAG_NAME = null
    }

    stages {
        stage('Always Run') {
          steps {
              container('sonar-scanner-cli') {
                  script {
                    echo "GIT_TAG_NAME: ${GIT_TAG_NAME}"
                  }
              }
          }
        }
        stage('Run if GIT_TAG_NAME is not null') {
            when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
            steps {
                container('sonar-scanner-cli') {
                    script {
                    def myVar = null // or some value
                    if (myVar == null) {
                        echo 'myVar is null'
                    } else {
                        echo 'myVar is not null'
                    }
                        echo "GIT_TAG_NAME: ${GIT_TAG_NAME}"
                        echo ' Hello Git Tag is not Null'
                    }
                }
            }
        }
    }
}
