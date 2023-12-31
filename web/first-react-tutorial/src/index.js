import React, {createContext} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import ExpensesProvider from "./ExpensesProvider";
import reportWebVitals from './reportWebVitals';

import stuff from "./file.json"

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
<ExpensesProvider>
    <App />
</ExpensesProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
