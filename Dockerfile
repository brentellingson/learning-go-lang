FROM golang:1.23.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /learning-golang-api cmd/server/main.go

FROM gcr.io/distroless/static-debian12

EXPOSE 8080
ENV GIN_MODE=release

COPY --from=build /learning-golang-api /
CMD ["/learning-golang-api"]