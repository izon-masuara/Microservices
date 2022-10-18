FROM golang:alpine3.16
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . /app
# RUN go run package/migration/user.go
EXPOSE 3002
CMD go run cmd/server.go