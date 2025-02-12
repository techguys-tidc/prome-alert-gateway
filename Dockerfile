FROM golang:1.23.6-alpine3.21 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN apk add git
RUN GOCACHE=OFF
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -v -o main main.go


FROM alpine:3.21
COPY --from=builder /build/main /app/
RUN apk add -U tzdata
ENV TZ=Asia/Bangkok
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
ENV GIN_MODE release
WORKDIR /app
ENTRYPOINT ["./main"]