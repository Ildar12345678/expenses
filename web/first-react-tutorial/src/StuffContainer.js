import React, { useState, useEffect } from 'react';
import CreatableSelect from 'react-select/creatable'

const StuffContainer = () => {
    const [name, setName] = useState("");
    const [expensesName, setExpenseNames] = useState([]);
    useEffect(() => {
        // Fetch existing expense names from the server
        fetch("http://localhost:4000/expname")
            .then(response => response.json())
            .then(data => setExpenseNames(data))
            .catch(error => console.error("Error:", error));
    }, []);

    const expensesOptions = [
        ...expensesName.map(expense => ({
            value: expense.Name,
            label: expense.Name,
            name: 'expense_name',
            expense_id: expense.ID
        }))
    ];
    // const handleExpenseChange = (selectedOption) => {
    //     setName({name:selectedOption});
    //     console.log("name", name.name.value)
    // };
    // ответ от сервера с названием товара, цено и датой
    const [answer, setAnswer] = useState([])
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("stringify", JSON.stringify({'name' : name}))
        // Make an HTTP request to send the data to the server
        fetch('http://localhost:4000/stuff', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({'name' : name})
        })
            .then(response => response.json())
            .then(setAnswer)
            .catch(error => {
                // Handle error, show error message or perform any necessary actions
                console.error('Error sending data:', error);
            });
    };

    const handleChange = (e) => {
        setName(e.target.value);
    };
    console.log(answer)

    return (
        <div>

            <form onSubmit={handleSubmit} className="add-data-form">
                <label>
                    Expanse Name:
                    <input type="text" name="name" onChange={handleChange} className="input-field" />
                </label>
                <button type="submit" className="submit-button">Send expense</button>
            </form>
            <div>
                <ul>
                    {
                        answer != null ?
                        answer.map((ans, i) => (
                        <div key={i} className="container">
                            <div className="box">{ans.Name}</div>
                            <div className="box">{ans.price}</div>
                            <div className="box" style={{width:'200px'}}>{ans.date.slice(0,10)}</div>
                        </div>
                    )) :
                            <div>no info</div>
                    }
                </ul>
            </div>
        </div>

    )
}


export default StuffContainer