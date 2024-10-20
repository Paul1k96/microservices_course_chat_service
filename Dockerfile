FROM golang:1.23.2-alpine as builder

COPY .  /github.com/Paul1k96/microservices_course_chat/
WORKDIR /github.com/Paul1k96/microservices_course_chat/

RUN go mod download
RUN go build -o ./bin/server cmd/main.go

FROM alpine:3.13

WORKDIR /root/
COPY --from=builder /github.com/Paul1k96/microservices_course_chat/bin/server .

CMD ["./server"]
