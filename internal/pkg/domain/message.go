package domain

import "time"

// Message is the domain object for chat messages
type Message struct {
	Xid         string
	XchatroomID string
	Xusername   string
	Xcontent    string
	XcreateTime time.Time
}

// NewMessage creates a new message object
func NewMessage(id string, chatroomID string, username string, content string, createTime time.Time) *Message {
	var m Message
	m.Xid = id
	m.XchatroomID = chatroomID
	m.Xusername = username
	m.Xcontent = content
	m.XcreateTime = createTime
	return &m
}

// ID returns the message id
func (m Message) ID() string {
	return m.Xid
}

// ChatroomID returns the id of the chatroom to which the message belongs
func (m Message) ChatroomID() string {
	return m.XchatroomID
}

// Username returns the username of the message sender
func (m Message) Username() string {
	return m.Xusername
}

// Content returns the message content
func (m Message) Content() string {
	return m.Xcontent
}

// CreateTime returns the time when the message was created
func (m Message) CreateTime() time.Time {
	return m.XcreateTime
}
