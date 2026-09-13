package logging

import (
	"context"

	"github.com/mephistolie/chefbook-backend-common/log"
)

type MQData struct {
	MessageID   string
	MessageType string
	UserID      string
}

func (Events) ConfigLoaded(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "config.loaded",
		Message:   "service configuration loaded",
		Component: "config",
	})
}

func (Events) StartupFailed(ctx context.Context, operation string, err error) {
	log.LogFatal(ctx, log.Event{
		Event:     "app.startup.failed",
		Message:   "service startup failed",
		Component: "app",
		Operation: operation,
	}, err)
}

func (Events) GRPCServerFailed(ctx context.Context, err error) {
	log.LogError(ctx, log.Event{
		Event:     "grpc.server.failed",
		Message:   "gRPC server stopped with an error",
		Component: log.ComponentGRPC,
		Operation: "serve",
	}, err)
}

func (Events) GRPCServerStarted(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "grpc.server.started",
		Message:   "gRPC server started",
		Component: log.ComponentGRPC,
		Operation: "serve",
	})
}

func (Events) PostgresHealthCheckFailed(ctx context.Context, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "postgres.health_check.failed",
		Message:   "database is unavailable",
		Component: log.ComponentPostgres,
		Operation: "ping",
	}, err)
}

func (Events) MQServerInitialized(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "mq.server.initialized",
		Message:   "MQ server initialized",
		Component: log.ComponentAMQP,
	})
}

func (Events) MQConsumerStopped(ctx context.Context, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "mq.consumer.stopped",
		Message:   "RabbitMQ consumer stopped with an error",
		Component: log.ComponentAMQP,
		Operation: "consume",
	}, err)
}

func (Events) MQMessageInvalidID(ctx context.Context, messageType string) {
	log.LogWarn(ctx, log.Event{
		Event:     "mq.message.invalid_id",
		Message:   "message has an invalid identifier",
		Component: log.ComponentAMQP,
		Payload: map[string]any{
			"message_type": messageType,
		},
	})
}

func (Events) MQMessageUnsupported(ctx context.Context, data MQData) {
	log.LogWarn(ctx, log.Event{
		Event:     "mq.message.unsupported_type",
		Message:   "unsupported message type",
		Component: log.ComponentAMQP,
		MessageID: data.MessageID,
		Payload: map[string]any{
			"message_type": data.MessageType,
		},
	})
}

func (Events) MQMessageRequeued(ctx context.Context, data MQData, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "mq.message.requeued",
		Message:   "message requeued",
		Component: log.ComponentAMQP,
		MessageID: data.MessageID,
		Payload: map[string]any{
			"message_type": data.MessageType,
		},
	}, err)
}

func (Events) MQMessageProcessing(ctx context.Context, data MQData) {
	log.Log(ctx, log.Event{
		Event:     "mq.message.processing",
		Message:   "processing message",
		Component: log.ComponentAMQP,
		MessageID: data.MessageID,
		UserID:    data.UserID,
		Payload: map[string]any{
			"message_type": data.MessageType,
		},
	})
}

func (Events) FirebaseImportDisabled(ctx context.Context, userID string) {
	log.LogWarn(ctx, log.Event{
		Event:     "user.firebase.import_disabled",
		Message:   "Firebase profile import is disabled",
		Component: log.ComponentFirebase,
		UserID:    userID,
		Operation: "import_profile",
	})
}

func (Events) FirebaseProfileLoadFailed(ctx context.Context, userID string, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "user.firebase.profile_load_failed",
		Message:   "Firebase profile could not be loaded",
		Component: log.ComponentFirebase,
		UserID:    userID,
		Operation: "load_profile",
	}, err)
}
