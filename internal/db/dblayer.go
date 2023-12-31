package dblayer

import (
	"expenses/internal/model"
	"time"
)

type DBLayer interface {
	GetCity() ([]*model.City, error)
	GetCat() ([]*model.Cat, error)
	GetSubcat() ([]*model.Subcat, error)
	GetSupplier() ([]*model.MOS, error)
	GetExpensesNames() ([]*model.Expense, error)
	GetExpenses(before, after time.Time) ([]*model.ExpenseShow, error)
	GetStatistics(before, after time.Time) ([]*model.Statistics, error)
	GetStuff(name string) ([]*model.Stuff, error)
	AddExpense(*model.ExpenseAdd) error
}
