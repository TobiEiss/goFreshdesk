package gofreshdesk

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

type TicketResponse struct {
	ID int `json:"id,omitempty"`
}

// CreateTicket posts a new Ticket to freshdesk
func (freshdesk *Freshdesk) CreateTicket(ticket Ticket) (ticketResponse TicketResponse, err error) {
	err = freshdesk.query(&ticketResponse, "/api/v2/tickets", ticket, nil)
	return ticketResponse, err
}

// func (freshdesk *Freshdesk) GetAnswer() (reponse interface{}, err error) {
// 	freshdesk.query("/api/v2/tickets")
// 	return
// }
