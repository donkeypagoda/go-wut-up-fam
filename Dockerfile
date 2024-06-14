# syntax=docker/dockerfile:1

FROM golang:1.22.4
RUN mkdir /build
WORKDIR /build
ADD go.mod cmd/web/main.go cmd/web/handlers.go /build/
COPY ui/html/home.html /build/ui/html/home.html
COPY internal/api.json /build/internal/api.json
ENV GOPATH /build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/web
# RUN go install /build/web
# COPY web /build/web
EXPOSE 3000
ENTRYPOINT ["./web"]