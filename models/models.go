package models

import "time"

type source int
type status int
type priority int

const (
	Email source = 1 + iota
	Portal
	Phone
	Forum
	Twitter
	Facebook
	Chat
)

const (
	Open status = 2 + iota
	Pending
	Resolved
	Closed
)

const (
	Low priority = 1 + iota
	Medium
	High
	Urgent
)

// Ticket represents a single helpdesk ticket
type Ticket struct {
	ProductID   int      `json:"product_id"`
	GroupID     int      `json:"group_id"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Subject     string   `json:"subject"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Status      status   `json:"status"`
	Priority    priority `json:"priority"`
	Source      source   `json:"source"`
}

// TicketResponse response for "createTicket"
type TicketResponse struct {
	ID int `json:"id,omitempty"`
}

// Conversations holds a bunch of conversations
type Conversations []Conversation

// Conversation is one piece of a conversation
type Conversation struct {
	BodyText     string        `json:"body_text"`
	Body         string        `json:"body"`
	ID           int           `json:"id"`
	Incoming     bool          `json:"incoming"`
	Private      bool          `json:"private"`
	UserID       int           `json:"user_id"`
	SupportEmail interface{}   `json:"support_email"`
	Source       int           `json:"source"`
	TicketID     int           `json:"ticket_id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	FromEmail    string        `json:"from_email"`
	ToEmails     []string      `json:"to_emails"`
	CcEmails     []string      `json:"cc_emails"`
	BccEmails    []string      `json:"bcc_emails"`
	Attachments  []interface{} `json:"attachments"`
}
