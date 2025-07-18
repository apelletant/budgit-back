package domain

import (
	"time"

	"github.com/google/uuid"
)

type AddExpenceReq struct {
	Value        int           `json:"value"`
	Interval     time.Duration `json:"interval"`
	Label        string        `json:"label"`
	CreationDate int64         `json:"creation_date"`
}

type Expence struct {
	CreationDate int64         `json:"creation_date"`
	Interval     time.Duration `json:"interval"`
	Value        int           `json:"value"`
	Label        string        `json:"label"`
	ID           uuid.UUID     `json:"id"`
}
