package logging

import (
	"context"

	"github.com/mephistolie/chefbook-backend-common/log"
)

type Events struct{}

type UserOperationData struct {
	UserID    string
	Operation string
}

type BatchOperationData struct {
	Operation      string
	RequestedCount int
}

func (Events) DatabaseTransactionFailed(ctx context.Context, operation string, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.postgres.transaction_failed",
		Message:   "database transaction failed",
		Component: log.ComponentPostgres,
		Operation: operation,
	}, err)
}

func (Events) UserMutationFailed(ctx context.Context, data UserOperationData, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.postgres.user_mutation_failed",
		Message:   "user data mutation failed",
		Component: log.ComponentPostgres,
		UserID:    data.UserID,
		Operation: data.Operation,
	}, err)
}

func (Events) UserReadFailed(ctx context.Context, data UserOperationData, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.postgres.user_read_failed",
		Message:   "user data read failed",
		Component: log.ComponentPostgres,
		UserID:    data.UserID,
		Operation: data.Operation,
	}, err)
}

func (Events) InboxOperationFailed(ctx context.Context, operation string, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.postgres.inbox_operation_failed",
		Message:   "inbox operation failed",
		Component: log.ComponentPostgres,
		Operation: operation,
	}, err)
}

func (Events) MinimalInfoQueryFailed(ctx context.Context, data BatchOperationData, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.postgres.minimal_info_query_failed",
		Message:   "minimal user information query failed",
		Component: log.ComponentPostgres,
		Operation: data.Operation,
		Payload: map[string]any{
			"requested_count": data.RequestedCount,
		},
	}, err)
}

func (Events) MinimalInfoQueryDegraded(ctx context.Context, data BatchOperationData, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "user.postgres.minimal_info_query_degraded",
		Message:   "a minimal user information row could not be decoded",
		Component: log.ComponentPostgres,
		Operation: data.Operation,
		Payload: map[string]any{
			"requested_count": data.RequestedCount,
		},
	}, err)
}

func (Events) AvatarObjectDeleteFailed(ctx context.Context, userID string, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.s3.avatar_delete_failed",
		Message:   "avatar object could not be deleted",
		Component: log.ComponentS3,
		UserID:    userID,
		Operation: "delete",
	}, err)
}

func (Events) AvatarUploadPolicyFailed(ctx context.Context, operation string, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.s3.avatar_upload_policy_failed",
		Message:   "avatar upload policy could not be built",
		Component: log.ComponentS3,
		Operation: operation,
		Payload: map[string]any{
			"object_type": "avatar",
		},
	}, err)
}

func (Events) AvatarUploadLinkGenerationFailed(ctx context.Context, err error) {
	log.LogError(ctx, log.Event{
		Event:     "user.s3.avatar_upload_link_generation_failed",
		Message:   "avatar upload link could not be generated",
		Component: log.ComponentS3,
		Operation: "generate",
		Payload: map[string]any{
			"object_type": "avatar",
		},
	}, err)
}

func (Events) AvatarLinkRejected(ctx context.Context, reason string) {
	log.LogDebug(ctx, log.Event{
		Event:     "user.s3.avatar_link_rejected",
		Message:   "avatar link was rejected",
		Component: log.ComponentS3,
		Operation: "parse",
		Payload: map[string]any{
			"object_type": "avatar",
			"reason":      reason,
		},
	})
}

func (Events) FirebaseClientInitialized(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "user.firebase.client_initialized",
		Message:   "Firebase client initialized",
		Component: log.ComponentFirebase,
	})
}
