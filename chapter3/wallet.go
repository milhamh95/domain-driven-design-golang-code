package chapter3

import (
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type WalletItem interface {
	GetBalance() (money.Money, error)
}

type Wallet struct {
	// This is our aggregate root, and is our wallet's identity
	id uuid.UUID

	// ownerID is the identity of the entity who owns the wallet.
	// We do not need to know all the details of an owner at all times
	// but this gives us the ability go get them when necessary
	ownerID uuid.UUID

	walletItems []WalletItem
}

type Card struct {
	balance money.Money
}

func NewCard(balance money.Money) Card {
	return Card{
		balance: balance,
	}
}

func (c Card) GetBalance() (money.Money, error) {
	return c.balance, nil
}

func (w Wallet) GetWalletBalance() (*money.Money, error) {
	var bal *money.Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}

		bal, err = bal.Add(&itemBal)
		if err != nil {
			return nil, errors.New("failed to increment balance")
		}
	}

	return bal, nil
}
