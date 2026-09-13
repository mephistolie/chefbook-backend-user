package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
)

func (r *Repository) CreateUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (user_id)
		VALUES ($1)
	`, usersTable)

	if _, err = tx.ExecContext(ctx, query, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "create",
		}, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return r.commitTransaction(ctx, tx)
}

func (r *Repository) ImportFirebaseName(ctx context.Context, userId uuid.UUID, displayName *string, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	query := fmt.Sprintf(`
		UPDATE %s
		SET display_name=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err = tx.ExecContext(ctx, query, displayName, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "import_firebase_name",
		}, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return r.commitTransaction(ctx, tx)
}

func (r *Repository) DeleteUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1
	`, usersTable)

	if _, err := tx.ExecContext(ctx, query, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "delete",
		}, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return r.commitTransaction(ctx, tx)
}

func (r *Repository) handleMessageIdempotently(ctx context.Context, messageId uuid.UUID) (*sql.Tx, error) {
	tx, err := r.startTransaction(ctx)
	if err != nil {
		return nil, err
	}

	addMessageQuery := fmt.Sprintf(`
		INSERT INTO %s (message_id)
		VALUES ($1)
	`, inboxTable)

	if _, err = tx.ExecContext(ctx, addMessageQuery, messageId); err != nil {
		if !isUniqueViolationError(err) {
			r.events.InboxOperationFailed(ctx, "insert", err)
		}
		return nil, errorWithTransactionRollback(tx, err)
	}

	deleteOutdatedMessagesQuery := fmt.Sprintf(`
		DELETE FROM %[1]v
		WHERE ctid IN
		(
			SELECT ctid
			FROM %[1]v
			ORDER BY timestamp DESC
			OFFSET 1000
		)
	`, inboxTable)

	if _, err = tx.ExecContext(ctx, deleteOutdatedMessagesQuery); err != nil {
		r.events.InboxOperationFailed(ctx, "delete_outdated", err)
		return nil, errorWithTransactionRollback(tx, err)
	}

	return tx, nil
}
