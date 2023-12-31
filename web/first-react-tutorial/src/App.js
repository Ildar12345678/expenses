import './App.css';
import React, {useState} from "react";
import ExpenseContainer from "./ExpenseContainer.js"
import StatisticContainer from "./StatisticContainer.js"
import StuffContainer from "./StuffContainer.js"
import AddExpenseContainer from "./AddExpenseContainer.js"
import Date from "./SetDate"
import ExpensesProvider from "./ExpensesProvider";



export default function App() {
  const [currentPage, setCurrentPage] = useState('Expenses')
  const [show, setShow] = useState(<ExpenseContainer />)
  return (
    <>
      <div>
        <h1>{currentPage}</h1>
        <nav>
          <ul>
            <li>
              <button onClick={() => {
                setCurrentPage('Expenses')
                setShow(<ExpenseContainer />)
              }}>Expenses</button>
            </li>
            <li>
              <button onClick={() => {
                setCurrentPage('Statistics')
                setShow(<StatisticContainer />)
              }}>Statistics</button>
            </li>
            <li>
              <button onClick={() => {
                setCurrentPage('Stuff')
                setShow(<StuffContainer />)
              }}>Stuff</button>
            </li>
            <li>
              <button onClick={() => {
                setCurrentPage('Add Expense')
                setShow(<AddExpenseContainer />)
              }}>Add Expense</button>
            </li>
            <li>
                <Date />
            </li>
          </ul>
        </nav>
        <div>
          {show}
        </div>
      </div>
    </>
  );
}

