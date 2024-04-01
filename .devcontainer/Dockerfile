FROM docker.io/library/golang:latest

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1

ENV GO111MODULE=off
# /workspaces is the default location for git clones, but I should be getting "first" from the `devcontainer.json` file 
# or the devcontainer creator invocation
ENV GOPATH=$GOPATH:/workspaces/first
