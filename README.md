# Telegram Bot Readme
[![Is compiling?](https://github.com/faculerena/telegrambot/actions/workflows/go.yml/badge.svg)](https://github.com/faculerena/telegrambot/actions/workflows/go.yml)
[![Fly Deploy](https://github.com/faculerena/telegrambot/actions/workflows/deploy.yml/badge.svg)](https://github.com/faculerena/telegrambot/actions/workflows/deploy.yml)

This is a Telegram bot that can provide current weather information and memes.

## Weather Functionality

_(The default [city] is Buenos Aires)_

To get current weather information for a specific city, use the following command:

```
/current [city]
```
![Current weather example](./assets/current_paris.png "current weather example")


To get the forecast for the next 3 days, use the following command:

```
/nextdays [city]
```
![Forecast weather example](./assets/forecast_paris.png "3 day forecast example")

The bot will respond with the forecasted weather conditions for the next 3 days for the specified city.


## Meme Functionality

To generate a handshake meme, use the following command:

```
/handshake a - b - c
```

You can replace "a", "b", and "c" with any words or phrases of your choosing to create a custom meme. The bot will respond with the handshake meme featuring your chosen words.

Example:

```
/handshake foo - bar - foobar
```

returns:

![Two hands, left hand "foo" and right hand "bar" handshaking in "foobar" ](./assets/foobar_handshake.jpg "foo bar example")


# More Info

This bot can be accessed at @hourlyweathercaba_bot in telegram. It's hosted in Fly.io using permanent storage to open and send the photos.

# To Do

- Adding verification to view logs in telegram.
- Adding more memes to a weather bot? probably.
- Caching the "forecast" to avoid multiple identical calls to the weather API



