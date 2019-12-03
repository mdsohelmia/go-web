package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"ikdev/smartcherry/config"
	"net/http"
)

func GetHttpServer(router *mux.Router, cfg config.Conf) *http.Server {
	serverString := fmt.Sprintf(":%d", cfg.Server.Port)

	return &http.Server{
		Addr:    serverString,
		Handler: router,
	}
}
