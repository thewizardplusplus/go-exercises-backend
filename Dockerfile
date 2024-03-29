FROM golang:1.15-alpine AS builder

RUN apk update && \
  apk add --no-cache curl git && \
  curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 && \
  chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/thewizardplusplus/go-exercises-backend
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only -v

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -ldflags='-w -s -extldflags "-static"' ./...

FROM alpine:3.13.2

RUN apk update && apk add --no-cache bash

COPY --from=builder /go/bin/go-exercises-backend /usr/local/bin/go-exercises-backend
COPY --from=builder /go/bin/go-exercises-manager /usr/local/bin/go-exercises-manager
COPY tools/wait-for-it.sh /usr/local/bin/wait-for-it.sh

CMD ["/usr/local/bin/go-exercises-backend"]
