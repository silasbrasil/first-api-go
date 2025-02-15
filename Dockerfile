# Build stage
FROM golang:1.23.5 AS builder
WORKDIR /api
COPY . .
RUN go build -o app .

# Final Stage
FROM golang:1.23.5-alpine
WORKDIR /api
COPY --from=builder /api/app .
CMD [ "./app" ]
