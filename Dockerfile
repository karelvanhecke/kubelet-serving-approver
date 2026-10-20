FROM docker.io/golang:1.27.3-trixie@sha256:1896654cfc5c2b5ad74c07f29cb36f5d2d1cbed42e3054070fbfe444e521e83c AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath github.com/karelvanhecke/kubelet-serving-approver

FROM gcr.io/distroless/static-debian13:nonroot@sha256:516757106de35205efb70b91ba03f6f1298424a0678c603255cf2852a7000087
COPY --from=build /src/kubelet-serving-approver /bin/kubelet-serving-approver
ENTRYPOINT [ "kubelet-serving-approver" ]
