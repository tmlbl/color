FROM golang:1.11.4

ENV CGO_ENABLED=0

COPY main.go /go/color/main.go
WORKDIR /go/color
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /color

FROM scratch

COPY --from=0 /color /color

ENTRYPOINT [ "/color" ]
