package chapter3

import (
	"errors"
	"github.com/Rhymond/go-money"
	"time"
)

type AuctionRefactored struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AuctionRefactored) GetAuctionElapsedDuration() time.Duration {
	return a.auctionStart.Sub(a.auctionEnd)
}

func (a *AuctionRefactored) GetAuctionEndTimeInUTC() time.Time {
	return a.auctionEnd
}

func (a *AuctionRefactored) SetAuctionEnd(auctionEnd time.Time) error {
	err := a.validateTimeZone(auctionEnd)
	if err != nil {
		return err
	}

	a.auctionEnd = auctionEnd
	return nil
}

func (a *AuctionRefactored) GetAuctionStartTimeInUTC() time.Time {
	return a.auctionStart
}

func (a *AuctionRefactored) SetAuctionStartTimeinUTC(auctionStart time.Time) error {
	err := a.validateTimeZone(auctionStart)
	if err != nil {
		return err
	}

	// in reality, we would like persist this to a database
	a.auctionStart = auctionStart
	return nil
}

func (a *AuctionRefactored) validateTimeZone(t time.Time) error {
	tz, _ := t.Zone()
	if tz != time.UTC.String() {
		return errors.New("time zone must be UTC")
	}
	return nil
}
