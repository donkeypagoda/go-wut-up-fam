# syntax=docker/dockerfile:1

FROM golang:1.22.4
RUN mkdir /build
ADD go.mod cmd/web/main.go cmd/web/handlers.go /build/
COPY ui/html/home.html /build/ui/html/home.html
COPY internal/api.json /build/internal/api.json
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./...
EXPOSE 3000
CMD ["./..."]