FROM golang as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build  -o ip-log ./cmd/ip-log/

FROM scratch

WORKDIR /app
COPY --from=builder /app/ip-log ./

EXPOSE 8080
ENTRYPOINT [ "/app/ip-log" ]
