package domain

import (
	"time"

	"github.com/google/uuid"
)

type AddExpenceReq struct {
	Value        int
	Interval     time.Duration
	Label        string
	CreationDate int64
}

type Expence struct {
	CreationDate int64
	Interval     time.Duration
	Value        int
	ID           uuid.UUID
}
