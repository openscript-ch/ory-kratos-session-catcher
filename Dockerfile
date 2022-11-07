
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src
COPY . .
RUN go mod download
RUN go mod verify 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/ory-kratos-session-catcher

FROM scratch
COPY --from=builder /go/bin/ory-kratos-session-catcher /go/bin/ory-kratos-session-catcher
EXPOSE 3000
ENTRYPOINT ["/go/bin/ory-kratos-session-catcher"]