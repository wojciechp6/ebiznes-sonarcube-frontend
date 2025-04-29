// src/components/Platnosci.js
import React, { useState } from 'react';
import axiosInstance from '../api/axiosInstance';

function Platnosci() {
    const [danePlatnosci, setDanePlatnosci] = useState({
        numerKarty: '',
        termin: '',
        cvv: ''
    });
    const [status, setStatus] = useState(null);

    const handleChange = (e) => {
        setDanePlatnosci({
            ...danePlatnosci,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        axiosInstance.post('/platnosci', danePlatnosci)
            .then(response => {
                setStatus("Płatność zrealizowana pomyślnie");
            })
            .catch(err => {
                setStatus(`Błąd: ${err.message}`);
            });
    };

    return (
        <div>
            <h2>Płatności</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Numer karty:</label>
                    <input
                        type="text"
                        name="numerKarty"
                        value={danePlatnosci.numerKarty}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>Termin (MM/YY):</label>
                    <input
                        type="text"
                        name="termin"
                        value={danePlatnosci.termin}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>CVV:</label>
                    <input
                        type="text"
                        name="cvv"
                        value={danePlatnosci.cvv}
                        onChange={handleChange}
                    />
                </div>
                <button type="submit">Dokonaj płatności</button>
            </form>
            {status && <p>{status}</p>}
        </div>
    );
}

export default Platnosci;
