FROM docker.io/golang:1.27.1-trixie@sha256:137ca8442e368f5f5bb4d4f84d8cc1f6c2d898edf8a380459eb407f89d149b0d AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath github.com/karelvanhecke/kubelet-serving-approver

FROM gcr.io/distroless/static-debian13:nonroot@sha256:e2e927ec666bae08560abb3c55d0659eceabb657f56b6782ab500a9fc7f555e3
COPY --from=build /src/kubelet-serving-approver /bin/kubelet-serving-approver
ENTRYPOINT [ "kubelet-serving-approver" ]
