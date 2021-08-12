FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 go build -o /hello-argo

EXPOSE 8080
ENTRYPOINT ["/hello-argo"]