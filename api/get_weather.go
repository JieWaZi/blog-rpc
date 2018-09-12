package api

import (
	"golang.org/x/net/context"
	pb "github.com/blog/proto"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"github.com/blog/api/models"
)

func (*BlogService) GetWeather(ctx context.Context, req *pb.EmptyMessages) (*pb.GetWeatherResponse, error) {

	resp, err := http.Get("http://aider.meizu.com/app/weather/listWeather?cityIds=101020600")
	if err != nil {
		log.Panic(err.Error())
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	var weathers models.WeatherResponse
	json.NewDecoder(resp.Body).Decode(&weathers)
	if len(weathers.Value) == 0 {
		return &pb.GetWeatherResponse{}, nil
	}
	value := weathers.Value[0]
	var maxTemperature, minTemperature []int32
	var week, date []string
	var dayWeathers []*pb.DayWeather
	for _, weather := range value.Weathers {

		dayWeather := &pb.DayWeather{
			Date:           weather.Date,
			MaxTemperature: weather.Temp_day_c,
			MinTemperature: weather.Temp_night_c,
			Weather:        weather.Weather,
			Week:           weather.Week,
			Img:            weather.Img,
		}
		dayWeathers = append(dayWeathers, dayWeather)
		maxTemperature = append(maxTemperature, weather.Temp_day_c)
		minTemperature = append(minTemperature, weather.Temp_night_c)
		week = append(week, weather.Week)
		date = append(date, weather.Date)
	}

	return &pb.GetWeatherResponse{
		CityName:       value.City,
		MaxTemperature: append(maxTemperature[6:], maxTemperature[:6]...),
		MinTemperature: append(minTemperature[6:], minTemperature[:6]...),
		Date:           append(date[6:], date[:6]...),
		Week:           append(week[6:], week[:6]...),
		DayWeather:     dayWeathers,
	}, nil
}
