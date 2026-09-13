package app

import (
	"context"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	userpb "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
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

var events logging.Events

func Run(cfg *config.Config) {
	log.InitWithService("user", *cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	ctx := context.Background()

	db, err := postgres.Connect(cfg.Database)
	if err != nil {
		events.StartupFailed(ctx, "connect_postgres", err)
		return
	}

	repository := postgres.NewRepository(db)

	userService, err := service.New(ctx, cfg, repository)
	if err != nil {
		events.StartupFailed(ctx, "initialize_service", err)
		return
	}

	var mqServer *amqp.Server = nil
	if len(*cfg.Amqp.Host) > 0 {
		mqServer, err = amqp.NewServer(cfg.Amqp, userService.MQ)
		if err != nil {
			events.StartupFailed(ctx, "initialize_mq_server", err)
			return
		}
		if err := mqServer.Start(); err != nil {
			events.StartupFailed(ctx, "start_mq_server", err)
			return
		}
		events.MQServerInitialized(ctx)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		events.StartupFailed(ctx, "listen_grpc", err)
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

	events.GRPCServerStarted(ctx)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			events.GRPCServerFailed(ctx, err)
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
