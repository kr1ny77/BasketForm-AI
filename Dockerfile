FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/server/

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /bin/server .
COPY web/ ./web/
RUN mkdir -p uploads results

EXPOSE 8080
ENV PORT=8080
ENV UPLOAD_DIR=uploads
ENV RESULTS_DIR=results

CMD ["./server"]
