pipeline {
  agent {
    kubernetes {
      yaml '''
apiVersion: v1
kind: Pod
spec:
  securityContext:
    runAsUser: 1000
    runAsGroup: 1000
    fsGroup: 1000
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
        '''
    }
  }
  parameters {
    string(defaultValue: '.kubernetes-deploy-kustomize', description: 'Kustomize Path', name: 'kustomizae_path')
    string(defaultValue: 'pso_cluster_kubeconfig', description: 'KubeConfig File to do deploy step', name: 'kubeconfig_credential_id')
  }

  environment {
      KUBERNETES_KUSTOMIZE_PATH="${params.kustomizae_path}"
      KUBECONFIG_FILE = credentials("${params.kubeconfig_credential_id}")
  }

  stages {
        stage('Kubectl') {
            steps {
                script {
                    container('kubectl') {
                        dir("${KUBERNETES_KUSTOMIZE_PATH}") {
      //sh "echo ${env.KUBECONFIG_FILE}"
      sh('kubectl --kubeconfig ${KUBECONFIG_FILE} get node -o wide')
      // sh "sleep 300"
                        }
                    }
                }
            }
        }

  }
}

