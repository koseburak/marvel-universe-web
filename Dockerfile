## STAGE 1 - build env.
FROM golang:1.17-buster AS build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/marvel-universe-web

WORKDIR /app
COPY *.html .
COPY *.env .
COPY assets/ assets/

## STAGE 2 - deploy env.
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app/marvel-universe-web"]