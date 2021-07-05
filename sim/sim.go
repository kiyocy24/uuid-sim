package sim

import (
	"github.com/google/uuid"
)

type Sim struct {
	digit int
}

func NewSim(d int) *Sim {
	return &Sim{
		digit: d,
	}
}

func (s *Sim) Run() (count uint64) {
	count = 1
	u1 := uuid.New()
	u2 := uuid.New()
	for !s.isSameUuid(u1, u2) {
		u2 = uuid.New()
		count++
	}

	return count
}

func (s *Sim) isSameUuid(u1, u2 uuid.UUID) bool {
	s1 := u1.String()[:s.digit]
	s2 := u2.String()[:s.digit]

	return s1 == s2
}
