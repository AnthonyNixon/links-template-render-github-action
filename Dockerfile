FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

COPY ./ ./
RUN GOOS=linux GOARCH=amd64 go build -o /bin/links-template-render-github-action main.go

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /bin/links-template-render-github-action ./

# Set the environment variable if it's passed as an argument
ENV GIT_SHA=$4

ENTRYPOINT ["/root/links-template-render-github-action"]