package http

type AddExpense struct {
	Label        string
	Value        int
	Interval     string
	CreationDate int64
}
