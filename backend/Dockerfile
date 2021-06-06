FROM golang:alpine as build-back
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN  go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN chmod +x ./golang-api
EXPOSE 8080
CMD ["./golang-api"]
