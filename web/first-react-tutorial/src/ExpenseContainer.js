import React, {useState, useEffect, useContext} from "react";
import ExpensesContext from './Context';

function GetExpenses() {
    const { random } = useContext(ExpensesContext);
    const [expenses, setExpenses] = useState()
    useEffect(() => {
        fetch(`http://localhost:4000/expenses`)
            .then(response => response.json())
            .then(setExpenses)
            .catch(console.log)
    }, [random])

    // console.log("expenses", expenses)
    if (expenses)
        return (
            <>
                <div className="main-container">
                    <ul>
                        {expenses.map((expense, i) => (
                            <div key={i} className="container">
                                <div className="box" style={{width:'200px'}}>{expense.Date.slice(0,10)}</div>
                                <div className="box">{expense.Price}</div>
                                <div className="box">{expense.Expense}</div>
                                <div className="box">{expense.Cat}</div>
                                <div className="box">{expense.Subcat}</div>
                            </div>
                        ))}
                    </ul>
                </div>
            </>
        )
    return null
}

export default function ExpenseContainer() {
    return (
            <GetExpenses />
    )
}