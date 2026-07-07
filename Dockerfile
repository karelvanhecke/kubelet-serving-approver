FROM docker.io/golang:1.26.5-bookworm@sha256:eb37f58646a901dc7727cf448cae36daaefaba79de33b5058dab79aa4c04aefb AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath github.com/karelvanhecke/kubelet-serving-approver

FROM gcr.io/distroless/static-debian12:nonroot@sha256:d093aa3e30dbadd3efe1310db061a14da60299baff8450a17fe0ccc514a16639
COPY --from=build /src/kubelet-serving-approver /bin/kubelet-serving-approver
ENTRYPOINT [ "kubelet-serving-approver" ]
