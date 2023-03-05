// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    meteosourceWeatherForecastResponse, err := UnmarshalMeteosourceWeatherForecastResponse(bytes)
//    bytes, err = meteosourceWeatherForecastResponse.Marshal()

package meteosource

import "encoding/json"

func UnmarshalMeteosourceWeatherForecastResponse(data []byte) (*MeteosourceWeatherForecastResponse, error) {
	var r *MeteosourceWeatherForecastResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MeteosourceWeatherForecastResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MeteosourceWeatherForecastResponse struct {
	Lat       string      `json:"lat"`
	Lon       string      `json:"lon"`
	Elevation int64       `json:"elevation"`
	Timezone  string      `json:"timezone"`
	Units     string      `json:"units"`
	Current   Current     `json:"current"`
	Hourly    Hourly      `json:"hourly"`
	Daily     interface{} `json:"daily"`
}

type Current struct {
	Icon        string  `json:"icon"`
	IconNum     int64   `json:"icon_num"`
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
	Wind        Wind    `json:"wind"`
	CloudCover  int64   `json:"cloud_cover"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Angle int64   `json:"angle"`
	Dir   string  `json:"dir"`
}

type Hourly struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	Date        string     `json:"date"`
	Weather     string     `json:"weather"`
	Icon        int64      `json:"icon"`
	Summary     string     `json:"summary"`
	Temperature float64    `json:"temperature"`
	Wind        Wind       `json:"wind"`
	CloudCover  CloudCover `json:"cloud_cover"`
}

type CloudCover struct {
	Total int64 `json:"total"`
}

type Type string

const (
	None Type = "none"
)
