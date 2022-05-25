package http

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Yujiman/e_commerce/goods/gatway/internal/config"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/routes"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/utils"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

func getRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	/* V1 */
	httpHelper.RegisterGroupRoutes(mux, &routes.MeRouteGroup) // me/...

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
