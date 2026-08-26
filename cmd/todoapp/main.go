package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_postgres_pool "github.com/mishalyamets1/golang_todoapp/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/server"
	users_repository_postgres "github.com/mishalyamets1/golang_todoapp/internal/features/users/repository/postgres"
	users_service "github.com/mishalyamets1/golang_todoapp/internal/features/users/service"
	user_transport_http "github.com/mishalyamets1/golang_todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	fmt.Println("Hello, todoapp")
	

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing new connection pool ")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust(),)

	if err != nil {
		logger.Fatal("failed to init postgres connection pool: %w", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))

	usersRepository := users_repository_postgres.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing http server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),

	)

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterApiRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server error", zap.Error(err))
	}
}