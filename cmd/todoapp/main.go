package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/mishalyamets1/golang_todoapp/internal/core/config"
	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_pgx_pool "github.com/mishalyamets1/golang_todoapp/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/server"
	statistics_postgres_repository "github.com/mishalyamets1/golang_todoapp/internal/features/statistics/repository/postgres"
	statistics_service "github.com/mishalyamets1/golang_todoapp/internal/features/statistics/service"
	statistics_transport_http "github.com/mishalyamets1/golang_todoapp/internal/features/statistics/transport/http"
	tasks_postgres_repository "github.com/mishalyamets1/golang_todoapp/internal/features/tasks/repository/postgres"
	task_service "github.com/mishalyamets1/golang_todoapp/internal/features/tasks/service"
	tasks_transport "github.com/mishalyamets1/golang_todoapp/internal/features/tasks/transport/http"
	users_repository_postgres "github.com/mishalyamets1/golang_todoapp/internal/features/users/repository/postgres"
	users_service "github.com/mishalyamets1/golang_todoapp/internal/features/users/service"
	user_transport_http "github.com/mishalyamets1/golang_todoapp/internal/features/users/transport/http"
	web_fs_repository "github.com/mishalyamets1/golang_todoapp/internal/features/web/repository/file_system"
	web_service "github.com/mishalyamets1/golang_todoapp/internal/features/web/service"
	web_transport_http "github.com/mishalyamets1/golang_todoapp/internal/features/web/transport/http"
	"go.uber.org/zap"

	_ "github.com/mishalyamets1/golang_todoapp/docs"
)

// @title Golang TODO API
// @version 1.0
// @description Todo application REST-API schema
// @host localhost:8080
// @BasePath /api/v1


func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()
	
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("time_zone", time.Local))

	logger.Debug("initializing new connection pool ")

	pool, err := core_pgx_pool.NewConnectionPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool: %w", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))

	usersRepository := users_repository_postgres.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRespository(pool)
	tasksService := task_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport.NewTasksHTTPHandler(tasksService)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPhandler(statisticsService)

	logger.Debug("initializing feature", zap.String("feature", "web"))
	webRespository := web_fs_repository.NewWebRepository()
	webService := web_service.NewWebService(webRespository)
	webTransportHTTP := web_transport_http.NewWebHTTPHandler(webService)

	logger.Debug("initializing http server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouter1.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterApiRouters(apiVersionRouter1)
	httpServer.RegisterRoutes(webTransportHTTP.Routes()...)
	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server error", zap.Error(err))
	}
}