package api

import (
	"context"
	"ema_indicator/internal/controller"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Server interface {
	GracefullListenAndServe(port int) error
}

type server struct {
	router     *mux.Router
	logger     *logrus.Logger
	controller *controller.IndicatorController
}

func NewServer(controller *controller.IndicatorController) *server {
	s := &server{
		router:     mux.NewRouter(),
		logger:     logrus.New(),
		controller: controller,
	}

	s.configureRouter()

	logrus.Info("starting api server")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Path("/calculate").
		Handler(s.controller.HandleCalculate()).
		Methods(http.MethodPost)
}

func (s *server) bindingAddressFromPort(port int) string {
	logrus.Info("listening on port ", port)
	return fmt.Sprintf(":%d", port)
}

func (s *server) GracefullListenAndServe(port int) error {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer := &http.Server{
		Addr:    s.bindingAddressFromPort(port),
		Handler: s,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		},
	}

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	return g.Wait()
}
