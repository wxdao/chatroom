package domain

// MessageRepository defines the behaviors of the repository of messages
type MessageRepository interface {
	MessageByID(id string) (*Message, error)
	MessageInRange(startID string, endID string) ([]*Message, error)
	SaveMessage(message *Message) error
}
