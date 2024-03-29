## dsapi

dsapi is a golang client API for accessing the [Dark Sky API](https://darksky.net/dev/docs). It includes an example client suitable for use on the command line. 

```
dsapi 34.09,-118.41
Forecast for lat: 34.090000, long: -118.410000:

-------------------------------------------------------
|   62|   61|   61|   62|   66|   71|   75|   78|   80|
|  🌕 |   🌕|  ⛅|   ⛅|   ⛅|   ⛅|  ☀️ |  ☀️|   ☀️|
|  Now|  9am| 10am| 11am| 12pm|  1pm|  2pm|  3pm|  4pm|
-------------------------------------------------------
 Partly Cloudy

 Feels like:        62°  |   Sunrise:         9:45am
 Precipitation:      0%  |   Sunset:          9:46pm
 Humidity:          77%  |   Wind:            1mph
 Cloud Cover:       38%  |   UV Index:        0
 
-------------------------------------------------------
Partly cloudy for the hour.

Powered by Dark Sky: https://darksky.net/poweredby/
```

## Usage

[Sign up](https://darksky.net/dev/register) and get a developer key. 
They offer a free key good for 1000 calls/day.


### Build/Installation
```
go install ./cmd/dsapi
```

### Command Line

```
DARK_SKY_APIKEY=<YOUR_API_KEY> dsapi 41.47,-81.67
```

### Lat/Long

How can you get a Lat/Long?

If you install [z2ll](https://github.com/armhold/z2ll), my zipcode-to-lat/long tool, 
you can easily look up coordinates based on zip code:

```bash
dsapi `z2ll 44120`
```

Another option is Google Maps. Right click on the map and select "What's here?"
At the bottom of your screen will be something like the following:

![Google Maps Image](https://github.com/armhold/dsapi/blob/master/map.png)
