FROM golang:1.22.0-alpine AS build
WORKDIR /usermanagement/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./cmd/microservice_2 ./cmd/main.go

FROM alpine:latest
WORKDIR /project/friendzy/

COPY --from=build /usermanagement/cmd/microservice_2 ./cmd/microservice_2
COPY --from=build /usermanagement/dev.env ./dev.env

RUN apk update && apk add --no-cache ca-certificates tzdata

EXPOSE 50001
ENTRYPOINT ["./cmd/microservice_2"]
