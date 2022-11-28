#
## pull submodules
git submodule update --init --recursive

**or**

git submodule update --force --recursive --init --remote


## Docker
docker build . -t registry.local/hs-dgraph:latest
docker push registry.local/hs-dgraph:latest

## Helm
helm package .\install

## Registry width helm cli
helm push dgraph-0.1.0.tgz  oci://helm.local --insecure-skip-tls-verify

## Registry width direct http
curl --data-binary "@dgraph-0.1.0.tgz" http://helm.local/api/charts

## Env Vars
zmq.SocketUrl=tcp://*:5558
dgraph.Host=10.23.92.23
dgraph.Port=9080
 
