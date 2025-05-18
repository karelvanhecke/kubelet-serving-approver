#!/bin/bash

err() {
    echo $1 >&2
    exit 1
}

err::notInstalled() {
    err "$1 not installed"
}

err::cleanup() {
    cleanup::cluster::delete
    err "$1"
}

precheck::kind() {
    kind version &>/dev/null || err::notInstalled kind
}

precheck::kubectl() {
    kubectl version --client &>/dev/null || err::notInstalled kubectl
}

precheck::docker() {
    docker version &>/dev/null || err::notInstalled docker
}

precheck::buildx() {
    docker buildx version &>/dev/null || err::notInstalled buildx
}

setup::image::build() {
    docker buildx build -t ghcr.io/karelvanhecke/kubelet-serving-approver:v0.0.0 . &>/dev/null \
    || err "Failed to build image"
}

setup::image::load() {
    kind load docker-image --name kubelet-serving-approver ghcr.io/karelvanhecke/kubelet-serving-approver:v0.0.0 &>/dev/null \
    || err::cleanup "Failed to load image"
}

setup::image::deploy() {
    kubectl apply -k manifests &>/dev/null\
    || err::cleanup "Failed to deploy image"
    kubectl wait -n kubelet-serving-approver --for=condition=Available --timeout=1m deployment/kubelet-serving-approver &>/dev/null \
    || err::cleanup "Deployment not ready after 1 minute"
}

setup::cluster::create() {
    kind create cluster --config tools/kind/config.yaml &>/dev/null \
    || err "Failed to create cluster"
}

cleanup::cluster::delete() {
    kind delete cluster --name kubelet-serving-approver &>/dev/null \
    || echo "Failed to cleanup cluster"
}

testing::approval() {
    kubectl wait csr --timeout=1m --field-selector=spec.signerName=kubernetes.io/kubelet-serving --for=condition=Approved &>/dev/null \
    || err "Fail: no CSR's were approved in time"
}

main() {
    echo -n "Precheck: "
    for precheck in precheck::{kind,kubectl,docker,buildx}; do
        $precheck
    done
    echo "ok"

    echo -n "Setup: "
    setup::image::build
    setup::cluster::create
    setup::image::load
    setup::image::deploy
    echo "ok"

    echo "Tests: "
    echo -n "  * approval: "
    testing::approval
    echo "ok"

    echo -n "Cleanup: "
    cleanup::cluster::delete
    echo "ok"
}

main
