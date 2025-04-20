package rpcs

import (
	"fmt"
	"net/http"

	"github.com/whynotavailable/svc"
)

type GetWeatherRequest struct {
	Location string `json:"location"`
}

type GetWeatherResponse struct {
	Weather string `json:"weather"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	body, err := svc.ReadJson[GetWeatherRequest](r)
	if err != nil {
		svc.WriteErrorBadRequest(w)
		return
	}

	svc.WriteJson(w, GetWeatherResponse{
		Weather: fmt.Sprintf("The weather in %s is rain.", body.Location),
	})
}
