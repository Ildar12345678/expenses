import React, {createContext, useContext, useEffect, useState} from "react";
import ExpensesContext from './Context';

export default function ExpensesProvider({ children }) {
    const [expenses, setExpenses] = useState()
    const [stats, setStats] = useState();
    const [random, setRandom] = useState()


    useEffect(() => {
        // console.log("inside fetch")
        // fetch(`http://localhost:4000/expenses`)
        //     .then(response => response.json())
        //     .then(setExpenses)
        //     .catch(console.log)
        // fetch(`http://localhost:4000/stats`)
        //     .then(response => response.json())
        //     .then(setStats)
        //     .catch(console.log)

    }, [])
    return (
        <ExpensesContext.Provider value={{random, setRandom}}>
            {children}
        </ExpensesContext.Provider>
    )
}