package models

type WeatherResponse struct {
	Code  string     `json:"code,omitempy"`
	Value []Weathers `json:"value,omitempy"`
}

type Weathers struct {
	City               string             `json:"city,omitempy"`
	CityId             int32              `json:"cityid,omitempy"`
	ProvinceName       string             `json:"provinceName,omitempy"`
	WeatherDetailsInfo WeatherDetailsInfo `json:"weatherDetailsInfo,omitempy"`
	Weathers           []Weather          `json:"weathers,omitempy"`
}

type WeatherDetailsInfo struct {
	Weather3HoursDetailsInfos []Weather3HoursDetailsInfo `json:"weather3HoursDetailsInfos,omitempy"`
}

type Weather3HoursDetailsInfo struct {
	StartTime           string `json:"startTime,omitempy"`
	EndTime             string `json:"endTime,omitempy"`
	HighestTemperature  int32  `json:"highestTemperature,string,omitempy"`
	LowerestTemperature int32  `json:"lowerestTemperature,string,omitempy"`
	Weather             string `json:"weather,omitempy"`
}

type Weather struct {
	Date         string `json:"date,omitempy" `
	Temp_day_c   int32  `json:"temp_day_c,string,omitempy"`
	Temp_night_c int32  `json:"temp_night_c,string,omitempy"`
	Weather      string `json:"weather,omitempy"`
	Week         string `json:"week,omitempy"`
	Img          string `json:"img,omitempy"`
}
