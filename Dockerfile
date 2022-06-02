FROM golang:1.17-alpine AS builder

ENV GOOS linux
RUN apk update --no-cache

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o dist


FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/dist /app/dist

CMD ["./dist"]