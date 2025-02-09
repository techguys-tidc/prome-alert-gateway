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
  parameters {
    string(defaultValue: 'prome-gateway-agent-env', description: '.env file credentialid', name: 'app_dot_env_credential_id')
    string(defaultValue: '.kubernetes-deploy-kustomize', description: 'Kustomize Path', name: 'kustomizae_path')
    string(defaultValue: 'k-harbor-01.server.maas', description: 'Container Registry Host for use in container tag', name: 'ContainerRegistryHost')
    string(defaultValue: 'prome-gateway', description: 'Container Registry Project for use in container tag', name: 'ContainerRegistryProject')
    string(defaultValue: 'prome-alert-gateway', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageName')
    string(defaultValue: 'v0.0.1', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageTag')
  }

  environment {
      TOKEN_CONTAINER_REGISTRY = credentials('harbor_k-harbor-01-token')
      APP_DOT_ENV_FILE = credentials("${params.app_dot_env_credential_id}")
      KUBERNETES_KUSTOMIZE_PATH="${params.kustomizae_path}"
      CONTAINER_REGISTRY_HOST="${params.ContainerRegistryHost}"
      CONTAINER_REGISTRY_PROJECT="${params.ContainerRegistryProject}"
      CONTAINER_REGISTRY_CONTAINER_NAME="${params.ContainerImageName}"
      CONTAINER_REGISTRY_CONTAINER_TAG="${params.ContainerImageTag}"
      GIT_TAG_NAME = gitTagName()
  }

  stages {
        stage('Load App .env file from Jenkins Secret') {
            steps {
                script {
                    container('kaniko') {
                        dir("${KUBERNETES_KUSTOMIZE_PATH}") {
                          sh('pwd')
                          sh('ls')
                          sh 'cp $APP_DOT_ENV_FILE .env'
                          sh 'cat .env'
                        }
                    }
                }
            }
        }
    // stage('Prepare Container Push Token') {
    //   steps {
    //       container('kaniko') {
    //           dir ('prome-alert-gateway') {
    //             sh('echo "{\\\"auths\\\":{\\\"$CONTAINER_REGISTRY_HOST\\\":{\\\"auth\\\":\\\"$TOKEN_CONTAINER_REGISTRY\\\"}}}"  > /kaniko/.docker/config.json')
    //           }
    //       }
    //   }
    // }
    // stage('kaniko build & push') {
    //   steps {
    //       container('kaniko') {
    //         dir('prome-alert-gateway') {
    //           script {
    //             def containerRegistryHost = "${params.ContainerRegistryHost}"
    //             def containerRegistryProject = "${params.ContainerRegistryProject}"
    //             def containerName = "${params.ContainerImageName}"
    //             // def containerTag = "${env.BUILD_NUMBER}"
    //             // def containerTag = "${params.ContainerImageTag}"
    //             def containerTag = "${env.GIT_TAG_NAME}"
    //             sh """
    //               echo "${containerRegistryHost}/${containerRegistryProject}/${containerName}:${containerTag}"
    //               /kaniko/executor --skip-tls-verify --context ./ --dockerfile ./Dockerfile --destination ${containerRegistryHost}/${containerRegistryProject}/${containerName}:${containerTag}
    //             """
    //           }
    //         }
    //       }
    //   }
    // }
  }
}


/** @return The tag name, or `null` if the current commit isn't a tag. */
String gitTagName() {
    commit = getCommit()
    if (commit) {
        desc = sh(script: "git describe --tags ${commit}", returnStdout: true)?.trim()
        if (isTag(desc)) {
            return desc
        }
    }
    return null
}

String getCommit() {
    return sh(script: 'git rev-parse HEAD', returnStdout: true)?.trim()
}

@NonCPS
boolean isTag(String desc) {
    match = desc =~ /.+-[0-9]+-g[0-9A-Fa-f]{6,}$/
    result = !match
    match = null // prevent serialisation
    return result
}