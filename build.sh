#!/bin/sh

build_push(){
  docker buildx build  --platform ${ARCHS} -t ${REGISTRY}/${NAME}:latest  --push .
}

helm_build_push(){
  FN=${NAME}-${VER}.tgz
  # rm ${FN}
  helm package ./install --version ${VER}
  curl --data-binary "@${FN}" https://helm.klogsolenopsys.org/api/charts
}

REGISTRY=registry.klogsolenopsys.org
NAME=alexstorm-hsm-dgraph
ARCHS="linux/amd64,linux/arm64"
VER=0.1.1


helm_build_push
build_push





