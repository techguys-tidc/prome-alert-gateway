pipeline {
  agent {
    kubernetes {
      yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:v1.23.2-debug
    command:
    - sleep
    - infinity
    tty: true
        '''
    }
  }

  environment {
      APP_DOT_ENV_FILE = credentials('prome-gateway-agent-env')
      GIT_TAG="v1.1"
  }
  stages {
        stage('Load Env File') {
            steps {
                script {
                    container('kaniko') {

                    def dotEnvFile = "${env.APP_DOT_ENV_FILE}"
                    // Load the file from the credential ID
                    echo "helloworld"
                    echo "============"
                    echo "${dotEnvFile}"
                    sh('pwd')
                      withCredentials([file(credentialsId: 'prome-gateway-agent-env', variable: 'FILE')]) {
                        dir('subdir') {
                          sh 'cp $FILE .env'
                          sh 'cat .env'
                        }
                      }
                    }
                }
            }
        }

    

        stage('Generate Kustomization File') {
            steps {
                script {
                  container('kaniko') {
                    def kustomizationContent = """
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: my-app
    newTag: ${env.GIT_TAG}
"""
                    writeFile(file: 'kustomization.yaml', text: kustomizationContent)
                    echo "Generated kustomization.yaml with tag ${env.GIT_TAG}"
                    sh('pwd')
                    sh('cat kustomization.yaml')
                }
            }
        }
        }

  }
}

// withCredentials([file(credentialsId: 'PRIVATE_KEY', variable: 'my-private-key'),
//                  file(credentialsId: 'PUBLIC_KEY', variable: 'my-public-key')]) {
//    sh "cp \$my-public-key /src/main/resources/my-public-key.der"
//    sh "cp \$my-private-key /src/main/resources/my-private-key.der"
// }