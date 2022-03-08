package controller

import (
	"ema_indicator/internal/service"
	"encoding/json"
	"net/http"
)

type IndicatorController struct {
	service service.IndicatorService
}

func NewIndicatorController(service service.IndicatorService) *IndicatorController {
	return &IndicatorController{
		service: service,
	}
}

func (c *IndicatorController) HandleCalculate() http.HandlerFunc {
	type request struct {
		Prev    float64 `json:"prev"`
		Current float64 `json:"current"`
		Period  int     `json:"period"`
	}
	type response struct {
		Value float64 `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body := &request{}
		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			// TODO send error
			return
		}

		value := c.service.Calculate(body.Prev, body.Current, body.Period)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response{
			Value: value,
		}); err != nil {
			// TODO send error
			return
		}
	}
}
