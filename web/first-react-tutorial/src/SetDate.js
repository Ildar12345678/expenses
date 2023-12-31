import React, {useContext, useState} from "react";
import ExpensesContext from "./Context"

export default function Date() {
    const [formData, setFormData] = useState({
        before: '',
        after: '',
    });
    const { setRandom } = useContext(ExpensesContext);
    let count = 0
    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        // Make an HTTP request to send the data to the server
        fetch('http://localhost:4000/date', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
           // .then(data => data.json())
            .then(() => setRandom(count++))
            .catch(error => {
                // Handle error, show error message or perform any necessary actions
                console.error('Error adding data:', error);
            });
        // count++
    };

    return (//style={{display: 'flex'}}
            <form onSubmit={handleSubmit} className="add-data-form">
                <label>
                    Before:
                    <input className="input-date" style={{fontSize: '20px'}} type="text" onChange={handleChange} name="before" value={formData.before}/>
                </label>
                <label>
                    After:
                    <input className="input-date" style={{fontSize: '20px'}} type="text" onChange={handleChange} name="after" value={formData.after}/>
                </label>
                <button type="submit">Send</button>
            </form>
    )
}