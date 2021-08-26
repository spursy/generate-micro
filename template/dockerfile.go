package template

var (
	DockerFileSRV = `FROM {{lower .Alias}} as builder

RUN mkdir /tmp/building

COPY . /tmp/building/

RUN cd /tmp/building && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s"  -o /app ./cmd/search/
ENV GOPROXY=https://goproxy.dev.longbridge-inc.com

FROM docker.longbridge-inc.com/lb/alpine-golang

COPY --from=builder /app /app
COPY --from=builder /tmp/building/config /config

WORKDIR /

CMD [ "./app", "--server_address", "0.0.0.0:8080"]

`
)



