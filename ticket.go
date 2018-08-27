package gofreshdesk

import (
	"fmt"

	"github.com/TobiEiss/goFreshdesk/models"
)

// CreateTicket posts a new Ticket to freshdesk
func (freshdesk *Freshdesk) CreateTicket(ticket models.Ticket) (ticketResponse models.TicketResponse, err error) {
	err = freshdesk.query(&ticketResponse, "/api/v2/tickets", ticket, nil)
	return ticketResponse, err
}

// GetConversations fetch all converations
func (freshdesk *Freshdesk) GetConversations(ticketID string) (response models.Conversation, err error) {
	freshdesk.query(&response, fmt.Sprintf("/api/v2/tickets/%s/conversations", ticketID), nil, nil)
	return
}
