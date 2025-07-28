// TODO REMOVE
//
//nolint:unused
package pgsql

import "time"

type Expense struct {
	id           string
	interval     time.Duration
	creationDate time.Time
}
