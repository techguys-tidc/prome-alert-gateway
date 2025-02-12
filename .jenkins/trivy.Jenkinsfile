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
  - name: trivy
    image: bitnami/trivy:0.59.1
    command:
    - sleep
    - infinity
    tty: true
        '''
        }
    }
  parameters {
    // CREDENTIAL NEEDS
    string(defaultValue: 'prome-gateway-agent-env', description: '.env file credentialid', name: 'app_dot_env_credential_id')
    string(defaultValue: 'pso_cluster_kubeconfig', description: 'KubeConfig File to do deploy step', name: 'kubeconfig_credential_id')
    string(defaultValue: 'harbor_k-harbor-01-username', description: 'Harbor Credential', name: 'harbor_user_credential_id')
    string(defaultValue: 'harbor_k-harbor-01-password', description: 'Harbor Credential', name: 'harbor_password_credential_id')
    // CI - HARBOR IMAGE
    string(defaultValue: 'k-harbor-01.server.maas', description: 'Container Registry Host for use in container tag', name: 'ContainerRegistryHost')
    string(defaultValue: 'prome-gateway', description: 'Container Registry Project for use in container tag', name: 'ContainerRegistryProject')
    string(defaultValue: 'prome-alert-gateway', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageName')
    string(defaultValue: 'dev-gong-v0.0.3', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageTag')
  }

  environment {
      // # HARBOR
      CI_REGISTRY_USER = credentials("${params.harbor_user_credential_id}")
      CI_REGISTRY_PASSWORD = credentials("${params.harbor_password_credential_id}")

      // # HARBOR CONFIGURATION
      CONTAINER_REGISTRY_HOST = "${params.ContainerRegistryHost}"
      CONTAINER_REGISTRY_PROJECT = "${params.ContainerRegistryProject}"
      CONTAINER_REGISTRY_CONTAINER_NAME = "${params.ContainerImageName}"
      CONTAINER_REGISTRY_CONTAINER_TAG = "${params.ContainerImageTag}"

  }
    stages {
    stage('trivy registry login') {
         
            steps {
                        container('trivy') {

                script {
                    env.CI_REGISTRY_TMP = env.CI_REGISTRY_USER+":"+env.CI_REGISTRY_PASSWORD
                    env.CI_REGISTRY_AUTH = sh(script: 'echo -n $CI_REGISTRY_TMP | base64', returnStdout: true).trim()
                    sh('trivy registry login --insecure --username \\"$CI_REGISTRY_USER\\" --password \\"$CI_REGISTRY_PASSWORD\\" $CONTAINER_REGISTRY_HOST')
                }

      }
    }
    
  }
        stage('Scan Image with Trivy') {
            steps {
                container('trivy') {
                    script {
                        //sh('trivy image --insecure --format template --template "@/opt/bitnami/trivy/contrib/html.tpl" -o trivy-report.html k-harbor-01.server.maas/prome-gateway/prome-alert-gateway:dev-gong-v0.0.3')
                        sh('trivy image --insecure --format template --template "@/opt/bitnami/trivy/contrib/html.tpl" -o trivy-report.html $CONTAINER_REGISTRY_HOST/$CONTAINER_REGISTRY_PROJECT/$CONTAINER_REGISTRY_CONTAINER_NAME:$CONTAINER_REGISTRY_CONTAINER_TAG')
                        archiveArtifacts artifacts: 'trivy-report.html', fingerprint: true
                    }
                }
            }
        }
    }
        post {
            always {
                publishHTML (target: [
                    reportDir: '',
                    reportFiles: 'trivy-report.html',
                    reportName: 'Trivy Security Report'
                ])
            }
        }
}


