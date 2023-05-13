package chapter3

type LegalForm int

const (
	Freelancer = iota
	Partnership
	LLC
	Corporation
)

func (l LegalForm) IsIndividual() bool {
	return l == Freelancer
}

func (l LegalForm) HasLimitedResponsibility() bool {
	return l == LLC || l == Corporation
}
