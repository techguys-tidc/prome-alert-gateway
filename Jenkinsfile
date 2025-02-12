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
    string(defaultValue: 'v0.0.1', description: 'Container Registry Tag for use in container tag', name: 'ContainerImageTag')
    // CD - DEPLOY K8SKUSTOMIZE
    string(defaultValue: '.kubernetes-deploy-kustomize', description: 'Kustomize Path', name: 'kustomizae_path')
    string(defaultValue: 'base', description: 'Kustomize Folder', name: 'kustomize_folder')
    string(defaultValue: 'prome-gateway', description: 'Deploy to Target Namespace', name: 'deploy_to_namespace')
    // SONARQUBE
    string(defaultValue: 'my-sonarqube-server', description: 'SonarQube in Jenkins Manage > Global > SonarQube', name: 'sonarqube_env')
  }

  environment {
      // # HARBOR
      CI_REGISTRY_USER = credentials("${params.harbor_user_credential_id}")
      CI_REGISTRY_PASSWORD = credentials("${params.harbor_password_credential_id}")
      // # KUBERNETES
      KUBERNETES_KUSTOMIZE_PATH = "${params.kustomizae_path}"
      KUBERNETES_KUSTOMIZE_FOLDER = "${params.kustomize_folder}"
      KUBECONFIG_FILE = credentials("${params.kubeconfig_credential_id}")
      KUBERNETES_DEPLOY_TO_NAMESPACE = "${params.deploy_to_namespace}"

      // # HARBOR CONFIGURATION
      CONTAINER_REGISTRY_HOST = "${params.ContainerRegistryHost}"
      CONTAINER_REGISTRY_PROJECT = "${params.ContainerRegistryProject}"
      CONTAINER_REGISTRY_CONTAINER_NAME = "${params.ContainerImageName}"
      CONTAINER_REGISTRY_CONTAINER_TAG = "${params.ContainerImageTag}"

      // # SONARQUBE
      SONARQUBE_ENV_NAME = "${params.sonarqube_env}"

      // # GIT
      GIT_TAG_NAME = gitTagName()
      // # APPLICATION
      APP_DOT_ENV_FILE = credentials("${params.app_dot_env_credential_id}")
  }

  stages {
    stage('Code Analysis') {
      steps {
          container('sonar-scanner-cli') {
              script {
                echo "Workspace Path: ${env.WORKSPACE}"
                withSonarQubeEnv("${SONARQUBE_ENV_NAME}") {
                  echo "SONAR-URL: ${env.SONAR_HOST_URL} "
                  sh("sonar-scanner -Dsonar.sources=${env.WORKSPACE}")
                }
              }
          }
      }
    }
    stage('Create /kaniko/.docker/config.json') {
                     when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
            steps {
                        container('kaniko') {
              dir(env.WORKSPACE) {
                script {
                    env.CI_REGISTRY_TMP = env.CI_REGISTRY_USER+":"+env.CI_REGISTRY_PASSWORD
                    env.CI_REGISTRY_AUTH = sh(script: 'echo -n $CI_REGISTRY_TMP | base64', returnStdout: true).trim()
                    sh('echo "{\\"auths\\":{\\"$CONTAINER_REGISTRY_HOST\\":{\\"auth\\":\\"$CI_REGISTRY_AUTH\\"}}}" > /kaniko/.docker/config.json')
                }
          }
      }
    }
    
  }
    stage('CI Kaniko Build Image & Push to Harbor') {
            when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
      steps {
          container('kaniko') {
              script {
                def containerRegistryHost = "${params.ContainerRegistryHost}"
                def containerRegistryProject = "${params.ContainerRegistryProject}"
                def containerName = "${params.ContainerImageName}"
                // def containerTag = "${env.BUILD_NUMBER}"
                // def containerTag = "${params.ContainerImageTag}"
                def containerTag = "${env.GIT_TAG_NAME}"
                sh """
                  echo "${containerRegistryHost}/${containerRegistryProject}/${containerName}:${containerTag}"
                  /kaniko/executor --skip-tls-verify --context ./ --dockerfile ./Dockerfile --destination ${containerRegistryHost}/${containerRegistryProject}/${containerName}:${containerTag}
                """
              }
          }
      }
    }
stage('trivy registry login') {
                     when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
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
                      when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
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
        stage('Generate Kustomization File & Copy .env file') {
            when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
      steps {
        script {
          container('kubectl') {
            dir("${KUBERNETES_KUSTOMIZE_PATH}/${KUBERNETES_KUSTOMIZE_FOLDER}") {
              def kustomizationContent = """
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - deployment.yaml
  - service.yaml

namespace: ${env.KUBERNETES_DEPLOY_TO_NAMESPACE}

images:
  - name: prome-alert-gateway
    newName: k-harbor-01.server.maas/prome-gateway/prome-alert-gateway
    newTag: ${env.GIT_TAG_NAME}

# patches:
#   - target:
#       kind: Ingress
#       name: my-app-ingress
#     patch: |-
#       - op: replace
#         path: /spec/rules/0/host
#         value: my-app.example.com  # Change this host as needed

generatorOptions:
  disableNameSuffixHash: true

secretGenerator:
  - name: app-config
    files:
      - .env
    type: Opaque

"""
              writeFile(file: 'kustomization.yaml', text: kustomizationContent)
              echo "Generated kustomization.yaml with tag ${env.GIT_TAG}"
              sh('cat kustomization.yaml')
              sh('cp $APP_DOT_ENV_FILE .env')
            // sh('cat .env')
            }
          }
        }
      }
        }
        stage('Deploy') {
            when {
                not{
                environment name: 'GIT_TAG_NAME', value: 'null'
                }
            }
      steps {
        script {
          container('kubectl') {
            dir("${KUBERNETES_KUSTOMIZE_PATH}") {
              //sh "echo ${env.KUBECONFIG_FILE}"
              sh('kubectl --kubeconfig ${KUBECONFIG_FILE} get node -o wide')
              sh('ls')
              // sh('kubectl --kubeconfig ${KUBECONFIG_FILE} kustomize ${KUBERNETES_KUSTOMIZE_FOLDER}')
              sh('kubectl --kubeconfig ${KUBECONFIG_FILE} apply -k ${KUBERNETES_KUSTOMIZE_FOLDER}')
            }
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

/** @return The tag name, or `null` if the current commit isn't a tag. */
String gitTagName() {
  commit = getCommit()
  if (commit) {
try {
      desc = sh(script: "git describe --tags ${commit}", returnStdout: true).trim()
      if (isTag(desc)) {
        return desc
      }
    } catch (Exception e) {
      // If the git describe fails, return null
      return null
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