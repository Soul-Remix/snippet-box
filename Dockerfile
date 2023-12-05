FROM golang:1.20 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/web/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk update && apk add bash
RUN apk add --update --no-cache openssh sshpass gcc musl-dev
WORKDIR /app

COPY --from=build /app .

CMD ["./app"]
