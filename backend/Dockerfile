FROM golang:latest as Builder
WORKDIR /build

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o car-rent .

FROM alpine:latest
WORKDIR /app

COPY --from=Builder /build/car-rent .

EXPOSE 8080
CMD [ "./car-rent" ]