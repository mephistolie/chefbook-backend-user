package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"strings"
)

func (r *Repository) CreateUser(userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
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

	if _, err = tx.Exec(query, userId); err != nil {
		log.Errorf("unable to create user %s: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) ImportFirebaseName(userId uuid.UUID, username *string, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	var firstName, secondName *string = nil, nil
	if username != nil {
		parts := strings.Split(*username, " ")
		firstName = &parts[0]
		if len(parts) > 1 {
			secondName = &parts[1]
		}
	}

	query := fmt.Sprintf(`
		UPDATE %s
		SET first_name=$1, last_name=$2
		WHERE user_id=$3
	`, usersTable)

	if _, err = tx.Exec(query, firstName, secondName, userId); err != nil {
		log.Errorf("unable to create user %s: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) DeleteUser(userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
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

	if _, err := tx.Exec(query, userId); err != nil {
		log.Errorf("unable to delete user %s: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) handleMessageIdempotently(messageId uuid.UUID) (*sql.Tx, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	addMessageQuery := fmt.Sprintf(`
		INSERT INTO %s (message_id)
		VALUES ($1)
	`, inboxTable)

	if _, err = tx.Exec(addMessageQuery, messageId); err != nil {
		if !isUniqueViolationError(err) {
			log.Error("unable to add message to inbox: ", err)
		}
		return nil, errorWithTransactionRollback(tx, err)
	}

	deleteOutdatedMessagesQuery := fmt.Sprintf(`
		DELETE FROM %[1]v
		WHERE ctid IN
		(
			SELECT ctid IN
			FROM %[1]v
			ORDER BY timestamp DESC
			OFFSET 1000
		)
	`, inboxTable)

	if _, err = tx.Exec(deleteOutdatedMessagesQuery); err != nil {
		return nil, errorWithTransactionRollback(tx, err)
	}

	return tx, nil
}
