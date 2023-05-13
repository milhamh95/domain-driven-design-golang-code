package chapter3

import (
	"github.com/Rhymond/go-money"
	"time"
)

// Auction is an entity to represent our action construct
type Auction struct {
	ID int

	// We use a specific money library as floats are not good ways to represent money
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
