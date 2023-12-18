package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	policyRoutes "github.com/anaclaraddias/brick/adapter/http/routes/policy"
	userRoutes "github.com/anaclaraddias/brick/adapter/http/routes/user"
	vehicleRoutes "github.com/anaclaraddias/brick/adapter/http/routes/vehicle"
	"github.com/anaclaraddias/brick/core/domain/helper"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

type Server struct {
	app        *gin.Engine
	httpServer *http.Server
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	serverConfiguration := &http.Server{
		Addr: routesConsts.DefaultPortConst,
	}

	return &Server{
		app,
		serverConfiguration,
	}
}

func (server *Server) Start() error {
	server.corsConfig()

	server.register()

	fmt.Println("server started")

	server.initialize()

	if err := server.shutdown(); err != nil {
		return err
	}

	return nil
}

func (server *Server) corsConfig() {
	headers := handlers.AllowedHeaders([]string{
		"Origin",
		"Content-Type",
		"Accept",
		"Content-Length",
		"Accept-Language",
		"Accept-Encoding",
		"Connection",
		"Access-Control-Allow-Origin",
	})

	origins := handlers.AllowedOrigins([]string{
		routesConsts.OriginLocalhostConst,
	})

	methods := handlers.AllowedMethods([]string{helper.GET, helper.POST, helper.PUT, helper.PATCH, helper.DELETE, helper.OPTIONS})
	credentials := handlers.AllowCredentials()

	corsHandler := handlers.CORS(headers, origins, methods, credentials)(server.app)
	server.httpServer.Handler = corsHandler
}

func (server *Server) register() {
	vehicleRoutes.NewVehicleRoutes(
		server.app,
	).Register()

	userRoutes.NewUserRoutes(
		server.app,
	).Register()

	policyRoutes.NewPolicyRoutes(
		server.app,
	).Register()
}

func (server *Server) initialize() {
	go func() {
		err := server.httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (server *Server) shutdown() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Println("kill signal received, shutting down server...")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(context); err != nil {
		return fmt.Errorf(helper.ErrorShuttingDownServerConst, err)
	}

	fmt.Println("server has been successfully shut down")

	return nil
}
