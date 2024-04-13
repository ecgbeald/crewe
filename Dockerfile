FROM golang:1.22.2 AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  COPY *.go ./

  RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# FROM build-stage AS test-stage
#   RUN go test -v ./...

FROM scratch AS release-stage
  WORKDIR /app

  COPY --from=build-stage /api /api

  EXPOSE 8080

  CMD ["/api"]