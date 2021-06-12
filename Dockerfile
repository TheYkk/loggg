FROM golang:1.16.5-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o logg main.go

FROM scratch
COPY --from=builder /app/logg /logg

CMD ["/logg"]