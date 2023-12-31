

export default function handleSubmit(e, link, formData) {
    e.preventDefault();

    // Make an HTTP request to send the data to the server
    fetch('http://localhost:4000/' + link, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {
            // Handle success, show success message or perform any necessary actions
            console.log('Data added successfully:', data);
            // Reset the form fields
            // setFormData({
            //     name: '',
            //     email: ''
            // });
        })
        .catch(error => {
            // Handle error, show error message or perform any necessary actions
            console.error('Error adding data:', error);
        });
};