FROM golang:1.25-alpine AS build

ARG SERVICE

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app ./cmd/${SERVICE}

FROM alpine:3.20
COPY --from=build /app /app
ENTRYPOINT ["/app"]



