package model

import (
	"time"
	"database/sql"
)

type Cat struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Subcat struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Cat_id int    `db:"cat_id"`
}

type City struct {
	ID   int    `db:"id"`
	City string `db:"city"`
}

type MOS struct {
	Name    string `db:"name"`
	Address string `db:"address"`
	ID      int    `db:"id"`
}

type Expense struct {
	ID        int           `db:"id"`
	Name      string        `db:"name"`
	Subcat_id int           `db:"subcat_id"`
	NDS       sql.NullInt32 `db:"nds"`
}

type Reply struct {
	Id          int            `db:"id"`
	Rate        int            `db:"rate"`
	Description sql.NullString `db:"description"`
	Mos_id      sql.NullInt32  `db:"mos_id"`
	Expense_id  sql.NullInt32  `db:"expense_id"`
}

type Purchase struct {
	ID            int            `db:"id"`
	Purchase_date time.Time      `db:"purchase_date"`
	City_id       int            `db:"city_id"`
	Online        bool           `db:"online"`
	Description   sql.NullString `db:"description"`
	Mos_id        sql.NullInt32  `db:"mos_id"`
}

type PurchaseCheck struct {
	ID          int     `db:"id"`
	Purchase_id int     `db:"purchase_id"`
	Expense_id  int     `db:"expense_id"`
	Count       float32 `db:"count"`
	Price       int     `db:"price"`
}

type ExpenseShow struct {
	Date    time.Time `db:"date"`
	Price   float32   `db:"price"`
	Expense string    `db:"expense"`
	Subcat  string    `db:"subcat"`
	Cat     string    `db:"cat"`
}

type Statistics struct {
	Cat       string  `db:"cat"`
	Subcat    string  `db:"subcat"`
	SumSubcat float32 `db:"sum_subcat"`
}

type Stuff struct {
	Name  string    `db:"name"`
	Price string    `json:"price"`
	Date  time.Time `json:"date"`
}

type ExpenseAdd struct {
	Purchase_date         string `json:"purchase_date"`
	Expense_name          string `json:"expense_name"`
	Expense_id            int    `json:"expense_id"`
	Subcat_id             int    `json:"subcat_id"`
	Mos_name              string `json:"mos_name"`
	Mos_address           string `json:"mos_address"`
	Mos_id                int    `json:"mos_id"`
	City_id               int8   `json:"city_id"`
	Online                bool   `json:"online"`
	Expense_count         string `json:"expense_count"`
	Expense_price         string `json:"expense_price"`
	NDS                   int    `json:"nds"`
	IsExpenseNameExisting bool   `json:"is_expense_name_existing"`
	IsMOSExisting         int8   `json:"is_mos_existing"`
}
