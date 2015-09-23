package i3status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type NgsWeatherWidget struct {
	BaseWidget
}

type WeatherData struct {
	Forecasts []struct {
		Temperature float32 `json:temperature`
		Wind        struct {
			Speed float32 `json:speed`
		} `json:wind`
		Links struct {
			City string `json:city`
		}
	} `json:forecasts`
}

const (
	WEATHER_URL string = "http://pogoda.ngs.ru/api/v1/forecasts/current?city="
	CITY        string = "nsk"
)

func NewNgsWeatherWidget() *NgsWeatherWidget {
	instanceCount++
	w := NgsWeatherWidget{
		BaseWidget: *NewBaseWidget(),
	}
	return &w
}

func (w *NgsWeatherWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "WEATHER"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		msg.FullText, msg.Color = w.getStatus(CITY)
		w.Output <- *msg
		time.Sleep(w.Refresh * 7200 * time.Millisecond)
	}
}

func (w *NgsWeatherWidget) getStatus(city string) (string, string) {
	data := new(WeatherData)
	respons, err := http.Get(WEATHER_URL + city)
	if err != nil {
		return fmt.Sprintf("WEATHER: CONN:ERR"), RED
	}
	defer respons.Body.Close()

	if err := json.NewDecoder(respons.Body).Decode(&data); err != nil {
		return fmt.Sprintf("WEATHER: ENCODE:ERR"), RED
	}

	return fmt.Sprintf("%gº, %g weend", data.Forecasts[0].Temperature, data.Forecasts[0].Wind.Speed), BLUE

}

func (w *NgsWeatherWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
