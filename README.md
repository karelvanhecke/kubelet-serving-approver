# Kubelet serving approver

Approves or denies certificate signing requests for the kubernetes.io/kubelet-serving signer.

## Installation

### Manifests
Default installation (basically approves anything):

`kubectl apply -f https://github.com/karelvanhecke/kubelet-serving-approver/releases/latest/download/kubelet-serving-approver.yaml`

It is recommended that you pass some options to configure the policy. Example:
```
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
- https://github.com/karelvanhecke/kubelet-serving-approver/releases/latest/download/kubelet-serving-approver.yaml

patches:
- patch: |-
    - op: add
      path: /spec/template/spec/containers/0/args
      value:
        - -p 
        - Static
        - --static-allowed-domains=example.org
        - --static-allowed-subnets=2001:db8::/32,192.0.2.0/24 
        - --static-match-host=^(control-plane|worker)-[0-9]+$
  target:
    kind: Deployment
    name: kubelet-serving-approver
```

### Flux
The installation manifests are also published on ghcr.io.
Manifests can be referenced as an [OCIRepository source](https://fluxcd.io/flux/components/kustomize/kustomizations/#source-reference).

OCI repository: `oci://ghcr.io/karelvanhecke/manifests/kubelet-serving-approver`

## Policies

Available policies:
* Static
