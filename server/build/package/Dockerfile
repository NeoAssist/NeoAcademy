FROM golang:1.14-alpine as base


FROM base as builder

WORKDIR /root

COPY . .

RUN go mod download

FROM base as runner

WORKDIR /root

COPY --from=builder /root /root

EXPOSE 80

CMD ["go", "run", "cmd/neoacademy/neoacademy.go"]
