package http

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/routes"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

func getRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	httpHelper.RegisterRoute(mux, &routes.AuthByRefresh)        // Auth by refresh
	httpHelper.RegisterRoute(mux, &routes.AuthByPasswordDomain) // Auth by login
	httpHelper.RegisterRoute(mux, &routes.AuthCheck)            // Auth check

	return mux
}

func InitServer() {
	port, err := config.GetServerPort()
	if err != nil {
		utils.LogFatalf("failed to get server address error: %v", err)
		return
	}

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      getRoutes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
