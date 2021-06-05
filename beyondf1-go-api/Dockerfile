FROM golang:alpine
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN  go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN chmod +x ./golang-api
CMD ["./golang-api"]
