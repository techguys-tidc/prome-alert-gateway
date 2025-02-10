
- [Custom MSTeam](#custom-msteam)
- [Go Lang Debug](#go-lang-debug)
  - [Clone git](#clone-git)
  - [Run](#run)
- [Build Docker Image](#build-docker-image)
  - [Clone git](#clone-git-1)
  - [Docker Build](#docker-build)
  - [Docker Run](#docker-run)
  - [Troubleshoot Container Log](#troubleshoot-container-log)
  - [Stop \& Terminate Container](#stop--terminate-container)
- [Call Prome Gateway API Services](#call-prome-gateway-api-services)
  - [Curl to endpoint /line-notify](#curl-to-endpoint-line-notify)
  - [Curl to endpoint /msteams-notify](#curl-to-endpoint-msteams-notify)
- [Jenkins CI/CD](#jenkins-cicd)
  - [CI using Jenkinsfile](#ci-using-jenkinsfile)
  - [CD using Kuberneetes](#cd-using-kuberneetes)


# Custom MSTeam
too custom MSTeams adaptive card go to this url

https://amdesigner.azurewebsites.net/

copy json payload in the website and put into json request under attribute content


# Go Lang Debug

## Clone git

```shell
{
git clone https://github.com/techguys-tidc/prome-alert-gateway.git
cd prome-alert-gateway
}
```

## Run

```shell
{
go run main.go
}
```

# Build Docker Image

## Clone git

```shell
{
git clone https://github.com/techguys-tidc/prome-alert-gateway.git
cd prome-alert-gateway
}
```


## Docker Build

```shell
{
CONTAINER_NAME=prome-alert-gateway
CONTAINER_TAG=0.01
docker build -t ${CONTAINER_NAME}:${CONTAINER_TAG} .
docker image ls | grep -i ${CONTAINER_NAME}
}
```

## Docker Run

```shell
{
CONTAINER_NAME=prome-alert-gateway
CONTAINER_TAG=0.01
docker run -d -p 8080:8080 --name ${CONTAINER_NAME}  \
  -e DEBUG_BODY=true \
  -e ENV_VAR2=value2 \
  ${CONTAINER_NAME}:${CONTAINER_TAG}
sleep .5
docker ps | grep ${CONTAINER_NAME}
}
```

## Troubleshoot Container Log

```shell
{
CONTAINER_NAME=prome-alert-gateway
docker logs -f --tail 100 ${CONTAINER_NAME}
}
```

## Stop & Terminate Container

```shell
{
CONTAINER_NAME=prome-alert-gateway
docker stop ${CONTAINER_NAME} && docker rm ${CONTAINER_NAME}
sleep .5
docker ps | grep ${CONTAINER_NAME}
}
```

# Call Prome Gateway API Services

## Curl to endpoint /line-notify

```shell
{
JSON_INPUT_FILENAME=example-request-test.json
PROME_GATEWAY_URL=127.0.0.1:8080
API_ENDPOINT=line-notify
LINE_API_TOKEN=1341@2342343243213123123423423423434
curl -X POST http://${PROME_GATEWAY_URL}/${API_ENDPOINT}?token=${LINE_API_TOKEN} \
     -H "Content-Type: application/json" \
     -d @${JSON_INPUT_FILENAME}
}
```

## Curl to endpoint /msteams-notify

```shell
{
JSON_INPUT_FILENAME=example-request-test.json
PROME_GATEWAY_URL=127.0.0.1:8080
API_ENDPOINT=msteams-notify
curl -X POST http://${PROME_GATEWAY_URL}/${API_ENDPOINT} \
     -H "Content-Type: application/json" \
     -d @${JSON_INPUT_FILENAME}
}
```

# Jenkins CI/CD

## CI using Jenkinsfile

```shell
{
./Jenkinsfile
}
```

## CD using Kuberneetes

```shell
{
export DOT_ENV_FILE="/somepath/.env"
cd ./.kubernetes-deploy-kustomize
cp ${DOT_ENV_FILE} base/.env
kubectl -k base
}
```