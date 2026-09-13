package amqp

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	auth "github.com/mephistolie/chefbook-backend-auth/api/mq"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
)

func (s *Server) handleProfileCreatedMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileCreated
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	events.MQMessageProcessing(ctx, logging.MQData{
		MessageID:   messageId.String(),
		MessageType: auth.MsgTypeProfileCreated,
		UserID:      userId.String(),
	})
	return s.service.CreateUser(ctx, userId, messageId)
}

func (s *Server) handleFirebaseImportMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileFirebaseImport
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	events.MQMessageProcessing(ctx, logging.MQData{
		MessageID:   messageId.String(),
		MessageType: auth.MsgTypeProfileFirebaseImport,
		UserID:      userId.String(),
	})
	return s.service.ImportFirebaseProfile(ctx, userId, body.FirebaseId, messageId)
}

func (s *Server) handleProfileDeletedMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileDeleted
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	events.MQMessageProcessing(ctx, logging.MQData{
		MessageID:   messageId.String(),
		MessageType: auth.MsgTypeProfileDeleted,
		UserID:      userId.String(),
	})
	return s.service.DeleteUser(ctx, userId, messageId)
}
