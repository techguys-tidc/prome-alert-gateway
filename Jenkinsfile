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
  parameters {
    // // CREDENTIAL NEEDS
    // string(defaultValue: 'prome-gateway-agent-env', description: '.env file credentialid', name: 'app_dot_env_credential_id')
    // string(defaultValue: 'pso_cluster_kubeconfig', description: 'KubeConfig File to do deploy step', name: 'kubeconfig_credential_id')
    // string(defaultValue: 'harbor_k-harbor-01-token', description: 'Harbor Credential', name: 'harbor_user_pass_credential_id')
    // // CI - HARBOR IMAGE
    // string(defaultValue: 'k-harbor-01.server.maas', description: 'Container Registry Host for use in container tag', name: 'ContainerRegistryHost')
    // string(defaultValue: 'prome-gateway', description: 'Container Registry Project for use in container tag', name: 'ContainerRegistryProject')
    // string(defaultValue: 'prome-alert-gateway', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageName')
    // string(defaultValue: 'v0.0.1', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageTag')
    // // CD - DEPLOY K8SKUSTOMIZE
    // string(defaultValue: '.kubernetes-deploy-kustomize', description: 'Kustomize Path', name: 'kustomizae_path')
    // string(defaultValue: 'base', description: 'Kustomize Folder', name: 'kustomize_folder')
    // string(defaultValue: 'prome-gateway', description: 'Deploy to Target Namespace', name: 'deploy_to_namespace')
  }

  environment {
    //   // # HARBOR
    //   TOKEN_CONTAINER_REGISTRY = credentials("${params.harbor_user_pass_credential_id}")
    //   // # KUBERNETES
    //   KUBERNETES_KUSTOMIZE_PATH = "${params.kustomizae_path}"
    //   KUBERNETES_KUSTOMIZE_FOLDER = "${params.kustomize_folder}"
    //   KUBECONFIG_FILE = credentials("${params.kubeconfig_credential_id}")
    //   KUBERNETES_DEPLOY_TO_NAMESPACE = "${params.deploy_to_namespace}"

    //   // # HARBOR CONFIGURATION
    //   CONTAINER_REGISTRY_HOST = "${params.ContainerRegistryHost}"
    //   CONTAINER_REGISTRY_PROJECT = "${params.ContainerRegistryProject}"
    //   CONTAINER_REGISTRY_CONTAINER_NAME = "${params.ContainerImageName}"
    //   CONTAINER_REGISTRY_CONTAINER_TAG = "${params.ContainerImageTag}"
    //   // # GIT
    //   GIT_TAG_NAME = gitTagName()
    //   // # APPLICATION
    //   APP_DOT_ENV_FILE = credentials("${params.app_dot_env_credential_id}")
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
                sh("sleep 300")
              }
          }
      }
    }
  }
}