# Changelog

## [0.1.7](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.6...v0.1.7) (2026-07-23)


### Bug Fixes

* **deps:** update docker.io/golang docker tag to v1.26.5 ([#96](https://github.com/karelvanhecke/kubelet-serving-approver/issues/96)) ([e863972](https://github.com/karelvanhecke/kubelet-serving-approver/commit/e863972388ed0489bf76c54dff9ce57579525bcd))
* **deps:** update kubernetes monorepo to v0.36.3 ([#111](https://github.com/karelvanhecke/kubelet-serving-approver/issues/111)) ([9bc5f7a](https://github.com/karelvanhecke/kubelet-serving-approver/commit/9bc5f7a4a61fdf4c43be1b0b19b4c13f13c4095f))
* **Dockerfile:** Golang container image digest ([#112](https://github.com/karelvanhecke/kubelet-serving-approver/issues/112)) ([f9dc4f4](https://github.com/karelvanhecke/kubelet-serving-approver/commit/f9dc4f44045f731871c41a615c1f66bb3e5835ea))
* **Dockerfile:** switch to Debian 13 based image ([#115](https://github.com/karelvanhecke/kubelet-serving-approver/issues/115)) ([4c21c11](https://github.com/karelvanhecke/kubelet-serving-approver/commit/4c21c1109054de92b303d77322e7eac31b5a67ec))

## [0.1.6](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.5...v0.1.6) (2026-06-14)


### Bug Fixes

* **deps:** update go transitive dependencies ([#79](https://github.com/karelvanhecke/kubelet-serving-approver/issues/79)) ([8a81a80](https://github.com/karelvanhecke/kubelet-serving-approver/commit/8a81a806ef0a27379973319c14adc783ebc87210))
* **deps:** update kubernetes monorepo to v0.36.2 ([#58](https://github.com/karelvanhecke/kubelet-serving-approver/issues/58)) ([5326e26](https://github.com/karelvanhecke/kubelet-serving-approver/commit/5326e26c0390f73270230ba0a488e077d94e8b10))

## [0.1.5](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.4...v0.1.5) (2026-06-13)


### Bug Fixes

* **deps:** update gcr.io/distroless/static-debian12:nonroot docker digest to d093aa3 ([#40](https://github.com/karelvanhecke/kubelet-serving-approver/issues/40)) ([cf55392](https://github.com/karelvanhecke/kubelet-serving-approver/commit/cf55392edff47a4692ece0a6e9ef3263746d7d92))
* **deps:** update go to v1.26.4 ([#73](https://github.com/karelvanhecke/kubelet-serving-approver/issues/73)) ([1019505](https://github.com/karelvanhecke/kubelet-serving-approver/commit/101950510bad2a704f9a98a0af1cabf15f720762))
* **deps:** update module github.com/miekg/dns to v1.1.72 ([#43](https://github.com/karelvanhecke/kubelet-serving-approver/issues/43)) ([2b455dd](https://github.com/karelvanhecke/kubelet-serving-approver/commit/2b455dda4950e4d952aa60577798c43b44b8131d))
* **deps:** update module k8s.io/klog/v2 to v2.140.0 ([#59](https://github.com/karelvanhecke/kubelet-serving-approver/issues/59)) ([aba9a53](https://github.com/karelvanhecke/kubelet-serving-approver/commit/aba9a534f5eed96951d13612dd84d0c05272828f))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.24.1 ([#60](https://github.com/karelvanhecke/kubelet-serving-approver/issues/60)) ([bd8b60d](https://github.com/karelvanhecke/kubelet-serving-approver/commit/bd8b60d8aeae5263b66a3a329c7b1a5ae09fb1f5))

## [0.1.4](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.3...v0.1.4) (2025-10-18)


### Bug Fixes

* **deps:** update Go depedencies ([#36](https://github.com/karelvanhecke/kubelet-serving-approver/issues/36)) ([bb053dd](https://github.com/karelvanhecke/kubelet-serving-approver/commit/bb053ddf371abd2fdd170ed75bcf1156a2b9b9d5))
* **deps:** update go to v1.25.3 ([#33](https://github.com/karelvanhecke/kubelet-serving-approver/issues/33)) ([8282129](https://github.com/karelvanhecke/kubelet-serving-approver/commit/828212993b56483d25bc861b81a12b83d6280ced))

## [0.1.3](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.2...v0.1.3) (2025-05-27)


### Features

* **csr:** log info when reconcile is successful ([#25](https://github.com/karelvanhecke/kubelet-serving-approver/issues/25)) ([cbae0a2](https://github.com/karelvanhecke/kubelet-serving-approver/commit/cbae0a29c4b6c1c395d0025f7947784838781564))

## [0.1.2](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.1...v0.1.2) (2025-05-19)


### Features

* **manifests:** release manifests with images pinned by digest ([#16](https://github.com/karelvanhecke/kubelet-serving-approver/issues/16)) ([f77e491](https://github.com/karelvanhecke/kubelet-serving-approver/commit/f77e491a897e49600387d9a0bfdbaccfd617ab83))

## [0.1.1](https://github.com/karelvanhecke/kubelet-serving-approver/compare/v0.1.0...v0.1.1) (2025-05-19)


### Bug Fixes

* **manifests:** trigger new release to fix missing manifests ([#12](https://github.com/karelvanhecke/kubelet-serving-approver/issues/12)) ([95187c5](https://github.com/karelvanhecke/kubelet-serving-approver/commit/95187c5c93770d6e2477d45811f99f421bc97671))

## 0.1.0 (2025-05-18)


### ⚠ BREAKING CHANGES

* **static:** drop RequireFQDN option ([#6](https://github.com/karelvanhecke/kubelet-serving-approver/issues/6))

### Features

* implement approver ([96c65b2](https://github.com/karelvanhecke/kubelet-serving-approver/commit/96c65b280a3b5012661a5e662ff40366a5519ba0))
* **static:** drop RequireFQDN option ([#6](https://github.com/karelvanhecke/kubelet-serving-approver/issues/6)) ([513e9e4](https://github.com/karelvanhecke/kubelet-serving-approver/commit/513e9e471acfb8fa7b5c1376fb5e752a9e9b7bcc))


### Continuous Integration

* add release workflow ([#3](https://github.com/karelvanhecke/kubelet-serving-approver/issues/3)) ([e27d0e3](https://github.com/karelvanhecke/kubelet-serving-approver/commit/e27d0e3f91fa6dd013b6868902c38b64be437a57))
