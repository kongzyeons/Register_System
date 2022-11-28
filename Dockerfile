FROM golang:1.18 AS build
WORKDIR /go/src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /app
RUN mkdir images
COPY --from=build /go/src/app .
EXPOSE 3000
CMD ["./app"]