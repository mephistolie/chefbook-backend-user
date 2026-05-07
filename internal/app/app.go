package app

import (
	"context"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	userpb "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/mephistolie/chefbook-backend-user/internal/repository/postgres"
	"github.com/mephistolie/chefbook-backend-user/internal/transport/amqp"
	"github.com/mephistolie/chefbook-backend-user/internal/transport/dependencies/service"
	user "github.com/mephistolie/chefbook-backend-user/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

func Run(cfg *config.Config) {
	log.InitWithService("user", *cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	ctx := context.Background()

	db, err := postgres.Connect(cfg.Database)
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "app.startup.failed",
			Message:   "service startup failed",
			Component: "app",
		}, err)
		return
	}

	repository := postgres.NewRepository(db)

	userService, err := service.New(ctx, cfg, repository)
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "app.startup.failed",
			Message:   "service startup failed",
			Component: "app",
		}, err)
		return
	}

	var mqServer *amqp.Server = nil
	if len(*cfg.Amqp.Host) > 0 {
		mqServer, err = amqp.NewServer(cfg.Amqp, userService.MQ)
		if err != nil {
			return
		}
		if err := mqServer.Start(); err != nil {
			log.LogFatal(ctx, log.Event{
				Event:     "app.startup.failed",
				Message:   "service startup failed",
				Component: "app",
			}, err)
			return
		}
		log.Log(ctx, log.Event{
			Event:     "mq.server.initialized",
			Message:   "mq server initialized",
			Component: log.ComponentAMQP,
		})
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "app.startup.failed",
			Message:   "service startup failed",
			Component: "app",
		}, err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			log.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	userServer := user.NewServer(*userService)

	go monitorHealthChecking(db, healthServer)

	userpb.RegisterUserServiceServer(grpcServer, userServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.LogError(ctx, log.Event{
				Event:     "grpc.server.failed",
				Message:   "error occurred while running grpc server",
				Component: log.ComponentGRPC,
			}, err)
		} else {
			log.Log(ctx, log.Event{
				Event:     "grpc.server.started",
				Message:   "grpc server started",
				Component: log.ComponentGRPC,
			})
		}
	}()

	wait := shutdown.Graceful(ctx, 5*time.Second, map[string]shutdown.Operation{
		"grpc-server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			return db.Close()
		},
		"mq": func(ctx context.Context) error {
			if mqServer == nil {
				return nil
			}
			return mqServer.Stop()
		},
	})
	<-wait
}
