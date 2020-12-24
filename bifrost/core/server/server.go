package server

import (
	log "github.com/go-kratos/kratos/pkg/log"
	"learn_code_go/bifrost/core/config"
	"learn_code_go/bifrost/core/route"
	"net/http"
	"time"
)

func listen(port string) {

	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
func Init(port string, routeRepo *route.RouteRepo, routeConfigs config.RouteConfigs) {
	http.Handle("/", &GateWayRequestHandler{RouteRepo: routeRepo, RoutesConfig: routeConfigs})
	log.Info("service starting.listen on " + port)
	//必须放到最后执行
	listen(port)
}
