package domain

import (
	"fmt"
	"time"

	core_errors "github.com/BladeRunner322/ChatGO/internal/core/errors"
)

type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Content    string
	SentAt     time.Time
	IsRead     bool
}

func NewMessage(
	id int,
	senderID int,
	receiverID int,
	content string,
	sentAt time.Time,
	isRead bool,
) Message {
	return Message{
		ID:         id,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		SentAt:     sentAt,
		IsRead:     isRead,
	}
}

func NewMessageUninitialised(
	senderID int,
	receiverID int,
	content string,
) Message {
	return Message{
		ID:         UninitializedID,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		SentAt:     time.Now(),
		IsRead:     false,
	}
}

func (m *Message) Validate() error {
	contentLen := len([]rune(m.Content))
	if contentLen < 1 || contentLen > 1000 {
		return fmt.Errorf(
			"invalid 'Content' len: %d: %w",
			contentLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
