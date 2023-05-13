package chapter3

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	id        uuid.UUID
	name      string
	legalForm LegalForm
	date      time.Time
}

func (c Customer) ToPerson() Person {
	return Person{
		fullName: c.name,
	}
}

func (c Customer) ToCompany() Company {
	return Company{
		name:         c.name,
		creationDate: c.date,
	}
}

type Person struct {
	fullName string
	birthday Birthday
}

type Company struct {
	name         string
	creationDate time.Time
}
