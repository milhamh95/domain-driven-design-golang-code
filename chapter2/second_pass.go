package chapter2

import "context"

/*
in the future, use can become
- lead
- lost lead
- customer
- churned

in preliminary_design, the domain doesn't have the concept of adding users.
using this phrase with domain expert is likely to confuse them.
*/

type LeadRequest struct {
	email string
}

type Lead struct {
	id string
}

type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	leadID string
	userID string
}

func (c *Customer) UserID() string {
	return c.userID
}

func (c *Customer) SetUserID(userID string) {
	c.userID = userID
}

type LeadConvertor interface {
	ConvertLead(ctx context.Context, subSelection SubscriptionType) (Customer, error)
}

func (l Lead) Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error) {
	panic("implement me")
}
