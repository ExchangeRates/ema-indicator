package internal

import (
	"ema_indicator/internal/api"
	"ema_indicator/internal/config"
	"ema_indicator/internal/controller"
	"ema_indicator/internal/service"
	"net/http"
)

func Start(config *config.Config) error {

	indicatorService := service.NewIndicator()
	indicatorController := controller.NewIndicatorController(indicatorService)

	srv := api.NewServer(indicatorController)
	bindingAddress := srv.BindingAddressFromPort(config.Port)

	return http.ListenAndServe(bindingAddress, srv)
}
