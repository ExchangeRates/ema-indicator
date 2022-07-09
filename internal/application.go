package internal

import (
	"ema_indicator/internal/api"
	"ema_indicator/internal/config"
	"ema_indicator/internal/controller"
	"ema_indicator/internal/service"
)

func Start(config *config.Config) error {

	indicatorService := service.NewIndicator()
	indicatorController := controller.NewIndicatorController(indicatorService)

	srv := api.NewServer(indicatorController)

	return srv.GracefullListenAndServe(config.Port)
}
