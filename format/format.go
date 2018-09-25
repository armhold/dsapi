package format

import (
	"fmt"
	"github.com/armhold/dsapi"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	maxHours         = 9
	hourFormat       = "3pm"
	sunriseSetFormat = "3:04pm"
)

func Print(forecast dsapi.Forecast) string {
	currHour := time.Unix(int64(forecast.Currently.Time), 0).Hour()

	// TODO: revisit if we go beyond 24 hours for maxHours
	moonPhaseForDay := forecast.Daily.Data[0].MoonPhase

	tempLine := "|"
	hourLine := "|"
	spacerLine := "|"
	border := "-------------------------------------------------------"

	for i := 0; i < len(forecast.Hourly.Data) && i < maxHours; i++ {
		hour := time.Unix(int64(forecast.Hourly.Data[i].Time), 0).Hour()

		if hour == currHour {
			tempLine += formatTemp(forecast.Currently.Temperature)
			hourLine += "  Now|"
			spacerLine += emojiForIcon(forecast.Currently.Icon, moonPhaseForDay)
		} else {
			tempLine += formatTemp(forecast.Hourly.Data[i].Temperature)
			hourLine += formatHour(forecast.Hourly.Data[i])
			spacerLine += emojiForIcon(forecast.Hourly.Data[i].Icon, moonPhaseForDay)
		}
	}

	cond := "\n"

	cond += fmt.Sprintf(" Feels like: %9d°  |   Sunrise: %14s\n",
		roundedTemp(forecast.Currently.ApparentTemperature),
		formatSunriseSet(forecast.Daily.Data[0].SunriseTime))

	cond += fmt.Sprintf(" Precipitation: %6d%%  |   Sunset: %15s\n",
		roundedProb(forecast.Currently.PrecipProbability*100),
		formatSunriseSet(forecast.Daily.Data[0].SunsetTime))

	cond += fmt.Sprintf(" Humidity: %11d%%  |   Wind: %12dmph\n",
		roundedProb(forecast.Currently.Humidity*100),
		roundedProb(forecast.Currently.WindSpeed))

	cond += fmt.Sprintf(" Dew Point: %10d°  |   UV Index: %8d\n",
		roundedProb(forecast.Currently.DewPoint),
		roundedProb(forecast.Currently.UVIndex))

	cond += fmt.Sprintf(" Cloud Cover: %8d%%  |\n",
		roundedProb(forecast.Currently.CloudCover*100))

	currSummary := fmt.Sprintf(" %s", forecast.Currently.Summary)

	return strings.Join([]string{
		border,
		tempLine,
		spacerLine,
		hourLine,
		border,
		currSummary,
		cond,
		border,
		forecast.Minutely.Summary,
		""},
		"\n")
}

func formatSunriseSet(unixTime int) string {
	t := time.Unix(int64(unixTime), 0)
	return t.Format(sunriseSetFormat)
}

func formatHour(data dsapi.Data) string {
	t := time.Unix(int64(data.Time), 0)
	s := fmt.Sprintf("%5s|", t.Format(hourFormat))

	return s
}

func formatTemp(temp float64) string {
	return fmt.Sprintf("%5s|", strconv.Itoa(roundedTemp(temp)))
}

func roundedTemp(temp float64) int {
	return int(math.Round(temp))
}

func roundedProb(prob float64) int {
	return int(math.Round(prob))
}

func emojiForIcon(icon string, moonPhase float64) string {

	switch icon {
	case "clear-day":
		return clearDayMsg

	case "clear-night":
		return msgForMoonPhase(moonPhase)

	case "rain":
		return rainMsg

	case "snow":
		return snowMsg

	case "sleet":
		return sleetMsg

	case "wind":
		return windMsg

	case "fog":
		return fogMsg

	case "cloudy":
		return cloudyMsg

	case "partly-cloudy-day":
		return partlyCloudyDayMsg

	case "partly-cloudy-night":
		// TODO: find way to show clouds
		return msgForMoonPhase(moonPhase)

	default:
		return defaultMsg
	}
}

func msgForMoonPhase(phase float64) string {

	// 🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘

	if phase < 0.125 {
		return "   🌑|"
	} else if phase < 0.25 {
		return "   🌒|"
	} else if phase < 0.375 {
		return "   🌓|"
	} else if phase < 0.5 {
		return "   🌔|"
	} else if phase < 0.625 {
		return "   🌕|"
	} else if phase < 0.75 {
		return "   🌖|"
	} else if phase < 0.875 {
		return "   🌗|"
	} else {
		return "   🌘|"
	}
}

const (
	// we define these with custom leading/trailing spaces because weather emoji do not seem
	// to render correctly in macOS Terminal/iTerm2. I believe this is due to the problem
	// described here: https://denisbider.blogspot.com/2015/09/when-monospace-fonts-arent-unicode.html

	clearDayMsg        = "   ☀️ |"
	rainMsg            = "   🌧 |️"
	snowMsg            = "   ❄️ |"
	sleetMsg           = defaultMsg
	windMsg            = "   🌬 |️"
	fogMsg             = "   🌫 |️"
	cloudyMsg          = "   ☁️ |"
	partlyCloudyDayMsg = "   ⛅|"
	defaultMsg         = "     |"
)
