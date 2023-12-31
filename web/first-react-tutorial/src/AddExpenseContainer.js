import React, { useState, useEffect } from 'react';
import CreatableSelect from 'react-select/creatable'

const AddExpenseContainer = () => {
    //контейнер для формы со всеми значениями покупки
    const [formData, setFormData] = useState({
        purchase_date: '',
        expense_name: '',
        expense_id: 0,
        subcat_id: 0,
        mos_name: '',
        mos_address: '',
        mos_id: 0,
        city_id: 0,
        online: '',
        expense_count: '',
        expense_price: '',
        nds: 0,
        is_expense_name_existing: false, //false - что-то новое; true - выбрали из списка
        is_mos_existing: 0, //0 - товар без поставщика, 1 - что-то новое, 2 - выбрали из списка
    });

    const [expensesName, setExpenseNames] = useState([]);
    const [mosNames, setMosNames] = useState([]);
    const [subcat, setSubcat] = useState([]);
    const [city, setCity] = useState([]);
    const [answer, setAnswer] = useState({
        error: '',
        response: '',
    });


    useEffect(() => {
        // Fetch existing expense names from the server
        fetch("http://localhost:4000/expname")
            .then(response => response.json())
            .then(data => setExpenseNames(data))
            .catch(error => console.error("Error:", error));

        // Fetch existing market/supplier names from the server
        fetch("http://localhost:4000/supplier")
            .then(response => response.json())
            .then(data => setMosNames(data))
            .catch(error => console.error("Error:", error));

        fetch("http://localhost:4000/subcat")
            .then(response => response.json())
            .then(data => setSubcat(data))
            .catch(error => console.error("Error:", error));

        fetch("http://localhost:4000/city")
            .then(response => response.json())
            .then(data => setCity(data))
            .catch(error => console.error("Error:", error));
    }, []);

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        // Make an HTTP request to send the data to the server
        fetch('http://localhost:4000/addexpense', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => response.json())
            .then(setAnswer)
            .then(() => {
            //     Handle success, show success message or perform any necessary actions
            //     console.log('Data added successfully:', data);
                // Reset the form fields
                setFormData({
                    purchase_date: '',
                    expense_name: '',
                    expense_id: 0,
                    subcat_id: 0,
                    mos_name: '',
                    mos_address: '',
                    mos_id: 0,
                    city_id: 0,
                    online: '',
                    expense_count: '',
                    expense_price: '',
                    nds: 0,
                    is_expense_name_existing: false,
                    is_mos_existing: 0, //0 - товар без поставщика, 1 - что-то новое, 2 - выбрали из списка
                });
            })
            .catch(error => {
                // Handle error, show error message or perform any necessary actions
                console.error('Error adding data:', error);
            });
    };

    const [selectedName, setSelectedName] = useState('');
    const [selectedMos, setSelectedMos] = useState('');
    // const [address, setAddress] = useState('');

    //обработчики селектов, от выбора которых зависит значение в других селектах (название траты и название поставщика)
    const handleExpenseChange = (selectedOption) => {
        setSelectedName(selectedOption);
        selectedOption.name
            ?
            setFormData({
                ...formData,
                [selectedOption.name]: selectedOption.value,
                ["expense_id"]: selectedOption.expense_id,
                ["is_expense_name_existing"]: true
    })
            :
            setFormData({
                ...formData,
                ["expense_name"]: selectedOption.value,
                ["is_expense_name_existing"]: false
            })

    };
    const handleMOSChange = (selectedOption) => {
        setSelectedMos(selectedOption);
        console.log(selectedOption.name)
        selectedOption.name
            ?
            setFormData({
                ...formData,
                [selectedOption.name]: selectedOption.value,
                ["mos_id"]: selectedOption.id,
                ["is_mos_existing"]: 2,
            })
            :
            setFormData({
                ...formData,
                ["mos_name"]: selectedOption.value,
                ["is_mos_existing"]: 1,
            })
    };
    const handleSelectChange = (selectedOption) => {
        setFormData({
            ...formData,
            [selectedOption.name]: selectedOption.value
        });
    };

    const cityOption = city.map(city => ({
        value: city.ID,
        label: city.City,
        name: 'city_id'
    }))

    const expensesOptions = [
        ...expensesName.map(expense => ({
            value: expense.Name,
            label: expense.Name,
            subcat_id: expense.Subcat_id,
            name: 'expense_name',
            expense_id: expense.ID
        }))
    ];

    const subcatIdOptions = !selectedName.__isNew__
        ? expensesName
            .filter((expense) => expense.Name === selectedName.value)
            .map((expense) => ({
                value: expense.Subcat_id,
                label: expense.Subcat_id,
                name: 'subcat_id'
            }))
        : subcat.map(item => ({label: item.Name, value: item.ID, name: 'subcat_id'}));

    const mosOptions = [
        ...mosNames.map(mos => ({
            value: mos.Name,
            label: mos.Name,
            address: mos.Address,
            name: 'mos_name',
            id: mos.ID
        }))
    ];

    const mosAddrOptions = mosNames
            .filter(mos => mos.Name === selectedMos.value)
            .map(mos => ({
                value: mos.Address,
                label: mos.Address,
                name: 'mos_address'
            }))

    const mosAddr = !selectedMos.__isNew__
        ? <CreatableSelect options={mosAddrOptions} onChange={handleSelectChange}/>
        : <input type="text" name="mos_address" value={formData.mos_address} onChange={handleChange} className="input-field" />

    console.log(formData)

    return (
        <form onSubmit={handleSubmit} className="add-data-form">
            <label>
                Purchase Date:
                <input type="text" name="purchase_date" value={formData.purchase_date} onChange={handleChange} className="input-field" />
            </label>
            <label>
                Expanse Name:
                <CreatableSelect options={expensesOptions} onChange={handleExpenseChange}/>
            </label>
            <label>
                Subcat ID:
                <CreatableSelect options={subcatIdOptions} onChange={handleSelectChange}/>
            </label>
            <label>
                Market or Supplier:
                <CreatableSelect options={mosOptions} onChange={handleMOSChange} />
            </label>
            <label>
                Address
                {mosAddr}
            </label>
            <label>
                City:
                <CreatableSelect options={cityOption} onChange={handleSelectChange}/>
            </label>
            <label>
                Online:
                <CreatableSelect
                    options={[{value:true,label:'true',name:'online'},
                        {value:false,label:'false',name:'online'}]}
                    onChange={handleSelectChange}
                />
            </label>
            <label>
                Expense Count:
                <input type="text" name="expense_count" value={formData.expense_count} onChange={handleChange} className="input-field" />
            </label>
            <label>
                Expense Price:
                <input type="text" name="expense_price" value={formData.expense_price} onChange={handleChange} className="input-field" />
            </label>
            <label>
                Expense NDS:
                <CreatableSelect
                    options={[{value:10,label:'10',name:'nds'},
                        {value:20,label:'20',name:'nds'},
                        {value:0,label:'0',name:'nds'}]}
                    onChange={handleSelectChange}
                />
            </label>
            <button type="submit" className="submit-button">Add Data</button>
            <label>
                answer:
                {answer.error ? answer.error : answer.response}
            </label>
        </form>
    );
};


export default AddExpenseContainer;