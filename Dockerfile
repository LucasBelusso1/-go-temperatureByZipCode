FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-temperatureByZipCode

FROM scratch
WORKDIR /app
COPY --from=build /app/go-temperatureByZipCode .
ENTRYPOINT ["./go-temperatureByZipCode"]