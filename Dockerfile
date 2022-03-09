FROM golang AS builder


RUN mkdir /app
WORKDIR /app

ADD go.mod  go.sum  ./
RUN go mod download

ADD main.go main.go ./
RUN CGO_ENABLED=0 go build -o /server

FROM alpine 
RUN apk add ca-certificates curl
ADD server.crt server.key ./
RUN chmod 755 server.crt server.key
COPY --from=builder /server .
CMD ["/server"]
