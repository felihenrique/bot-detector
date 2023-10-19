FROM golang:1.21-alpine as builder
WORKDIR /home/app
COPY . .
RUN go build -ldflags "-s -w"

FROM golang:1.21-alpine
WORKDIR /home/app
COPY --from=builder /home/app/botdetector /home/app/botdetector
COPY res /home/app/res
EXPOSE 10658
ENV GIN_MODE=release
CMD ["./botdetector"]