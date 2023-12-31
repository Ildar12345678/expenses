package dblayer

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"expenses/internal/model"
	"time"
	"sync"
	"log"
)

type DB struct {
	db *sqlx.DB
}

var (
	DBInstance *DB
	once       sync.Once
)

func NewDB() (*DB, error) {
	dsn := "user=gen host=/var/run/postgresql dbname=expenses3"
	if DBInstance == nil {
		once.Do(
			func() {
				db, err := openDB(dsn)
				if err != nil {
					log.SetFlags(log.Llongfile)
					log.Fatalln(err)
				}
				DBInstance = &DB{db: db}
			})
	}
	return DBInstance, nil
}

func (db *DB) GetCity() (dest []*model.City, err error) {
	stmt := "select id, city from city"
	err = db.db.Select(&dest, stmt)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetCat() (dest []*model.Cat, err error) {
	stmt := "select id, name from cat"
	err = db.db.Select(&dest, stmt)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetSubcat() (dest []*model.Subcat, err error) {
	stmt := "select id, name, cat_id from subcat order by cat_id"
	err = db.db.Select(&dest, stmt)
	if err != nil {
		return nil, err
	}
	
	return dest, nil
}
func (db *DB) GetSupplier() (dest []*model.MOS, err error) {
	stmt := "select name, address, id from market_or_supplier"
	err = db.db.Select(&dest, stmt)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetExpensesNames() (dest []*model.Expense, err error) {
	stmt := "select id, name, subcat_id, nds from expense"
	err = db.db.Select(&dest, stmt)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetExpenses(before, after time.Time) (dest []*model.ExpenseShow, err error) {
	stmt := `select p.purchase_date as date,
                  c.count * c.price as price,
                  e.name as expense,
                  s.name as subcat,
                  cat.name as cat
						from purchase as p, purchase_check as c, expense as e, subcat as s, cat
						where e.subcat_id = s.id and
									c.expense_id = e.id and
									c.purchase_id = p.id and
									s.cat_id = cat.id and
									p.purchase_date >= $1 and
									p.purchase_date <= $2 order by date`
	fmt.Println(" GetExpenses before, after:", before, after)
	
	err = db.db.Select(&dest, stmt, before, after)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetStatistics(before, after time.Time) (dest []*model.Statistics, err error) {
	stmt := `select s.name subcat, cat.name cat, sum(c.count * c.price) sum_subcat
					 from subcat s, cat, purchase p, expense e
           left outer join purchase_check c on e.id = c.expense_id
					 where e.subcat_id = s.id and
						 cat.id = s.cat_id and
             p.id = c.purchase_id and
             p.purchase_date >= $1 and
						 p.purchase_date <= $2
					 group by s.name, cat.name order by sum_subcat desc `
	
	err = db.db.Select(&dest, stmt, before, after)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) GetStuff(name string) (dest []*model.Stuff, err error) {
	stmt := `select e.name, c.price, p.purchase_date date
from expense e, purchase_check c, purchase p
where e.id = c.expense_id and p.id = c.purchase_id and e.name like '%'||$1||'%' order by purchase_date`
	err = db.db.Select(&dest, stmt, name)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
func (db *DB) AddExpense(expToAdd *model.ExpenseAdd) error {
	tx, err := db.db.Beginx()
	if err != nil {
		fmt.Println("error while db.db.Beginx()")
		return err
	}
	var stmt string
	if !expToAdd.IsExpenseNameExisting { // новый расход
		stmt = `insert into expense (name, subcat_id, nds) VALUES ($1, $2, $3)`
		fmt.Println(stmt, expToAdd.Expense_name, expToAdd.Subcat_id, expToAdd.NDS)
		_, err = tx.Exec(stmt, expToAdd.Expense_name, expToAdd.Subcat_id, expToAdd.NDS)
		if err != nil {
			fmt.Println("error while inserting new expense")
			return err
		}
		if expToAdd.IsMOSExisting == 0 { // поставщика нет
			stmt = `insert into purchase (purchase_date, city_id, online) values ($1, $2, $3)`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			if err != nil {
				fmt.Println("error while inserting purchase (new expense, no mos)")
				return err
			}
		} else if expToAdd.IsMOSExisting == 2 { // выбрали поставщика из списка
			stmt = `insert into purchase (purchase_date, city_id, online, mos_id) values ($1, $2, $3, $4)`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online, expToAdd.Mos_id)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online, expToAdd.Mos_id)
			if err != nil {
				fmt.Println("error while inserting purchase (new expense, mos from list)")
				return err
			}
		} else if expToAdd.IsMOSExisting == 1 { // новый поставщик
			stmt = `insert into market_or_supplier (name, address) values ($1, $2)`
			fmt.Println(stmt, expToAdd.Mos_name, expToAdd.Mos_address)
			_, err = tx.Exec(stmt, expToAdd.Mos_name, expToAdd.Mos_address)
			if err != nil {
				fmt.Println("error while inserting new mos")
				return err
			}
			stmt = `insert into purchase (purchase_date, city_id, online, mos_id)
							select $1, $2, $3, last_value from market_or_supplier_id_seq`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			if err != nil {
				fmt.Println("error while inserting purchase (new expense, new mos)")
				return err
			}
		}
		stmt = `insert into purchase_check (expense_id, purchase_id, count, price)
    							select e.last_value, p.last_value, $1, $2 from expense_id_seq e, purchase_id_seq p`
		fmt.Println(stmt, expToAdd.Expense_count, expToAdd.Expense_price)
		_, err = tx.Exec(stmt, expToAdd.Expense_count, expToAdd.Expense_price)
		if err != nil {
			fmt.Println("error while inserting purchase_check (new expense)")
			return err
		}
	} else {                           // расход из списка
		if expToAdd.IsMOSExisting == 0 { // поставщика нет
			stmt = `insert into purchase (purchase_date, city_id, online) values ($1, $2, $3)`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			if err != nil {
				fmt.Println("error while inserting new purchase (expense from list, no mos)")
				return err
			}
		} else if expToAdd.IsMOSExisting == 2 { // выбрали поставщика из списка
			stmt = `insert into purchase (purchase_date, city_id, online, mos_id) values ($1, $2, $3, $4)`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online, expToAdd.Mos_id)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online, expToAdd.Mos_id)
			if err != nil {
				fmt.Println("error while inserting new purchase (expense from list, mos from list)")
				return err
			}
		} else if expToAdd.IsMOSExisting == 1 { // новый поставщик
			stmt = `insert into market_or_supplier (name, address) values ($1, $2)`
			fmt.Println(stmt, expToAdd.Mos_name, expToAdd.Mos_address)
			_, err = tx.Exec(stmt, expToAdd.Mos_name, expToAdd.Mos_address)
			if err != nil {
				fmt.Println("error while inserting new mos (expense from list)")
				return err
			}
			stmt = `insert into purchase (purchase_date, city_id, online, mos_id)
							select $1, $2, $3, last_value from market_or_supplier_id_seq`
			fmt.Println(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			_, err = tx.Exec(stmt, expToAdd.Purchase_date, expToAdd.City_id, expToAdd.Online)
			if err != nil {
				fmt.Println("error while inserting new purchase (expense from list, new mos)")
				return err
			}
		}
		stmt = `insert into purchase_check (expense_id, purchase_id, count, price)
							select $1, p.last_value, $2, $3 from purchase_id_seq p`
		fmt.Println(stmt, expToAdd.Expense_id, expToAdd.Expense_count, expToAdd.Expense_price)
		_, err = tx.Exec(stmt, expToAdd.Expense_id, expToAdd.Expense_count, expToAdd.Expense_price)
		if err != nil {
			fmt.Println("error while inserting new purchase_check (expense from list)")
			return err
		}
	}
	_ = tx.Commit()
	// fmt.Println("transaction has rollbacked")
	return nil
}

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func closeDB(db *DB) {
	if err := db.db.Close(); err != nil {
		fmt.Println("error while closing DB:", err)
	}
}
