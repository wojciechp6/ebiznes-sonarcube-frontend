// src/components/Produkty.js
import React, { useState, useEffect } from 'react';
import axiosInstance from '../api/axiosInstance';

function Produkty() {
    const [produkty, setProdukty] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        axiosInstance.get('/produkty')
            .then(response => {
                setProdukty(response.data);
                setLoading(false);
            })
            .catch(err => {
                setError(err.message);
                setLoading(false);
            });
    }, []);

    if (loading) return <div>Ładowanie produktów...</div>;
    if (error) return <div>Błąd: {error}</div>;

    return (
        <div>
            <h2>Produkty</h2>
            <ul>
                {produkty.map((produkt) => (
                    <li key={produkt.id}>
                        {produkt.nazwa} - {produkt.cena} zł
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Produkty;
