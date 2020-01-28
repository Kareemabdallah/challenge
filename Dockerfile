FROM golang:1.8 as builder

RUN mkdir -p /build
ADD ./app_1 /build/
WORKDIR /build

RUN go get -d github.com/gorilla/mux
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app1 .

# Stage 2
FROM alpine:latest
RUN apk add --no-cache curl
COPY --from=builder /build/ /app/
WORKDIR /app
EXPOSE 9000
CMD [ "./app1" ]