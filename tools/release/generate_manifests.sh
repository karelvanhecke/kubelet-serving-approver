#!/bin/bash

if [[ -z $1 ]]; then
    echo "Tag must be specified as argument" >&2
    exit 1
fi

mkdir -p manifests/release

envsubst <<EOF > manifests/release/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

images:
- name: ghcr.io/karelvanhecke/kubelet-serving-approver
  newTag: $1

resources:
- ../base
EOF
kubectl kustomize -o manifests/release.yaml manifests/release
