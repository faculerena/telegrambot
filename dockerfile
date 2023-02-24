FROM golang:1.17-alpine

WORKDIR /app

COPY .. .

RUN go mod download

RUN go build -o ./build/hourlyWeatherBot

EXPOSE 8080

ADD /memes /assets
CMD ["./build/hourlyWeatherBot"]