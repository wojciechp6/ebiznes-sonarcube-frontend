// src/components/Koszyk.js
import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext';

function Koszyk() {
    const { koszyk } = useContext(AppContext);

    return (
        <div>
            <h2>Koszyk</h2>
            {koszyk.length === 0 ? (
                <p>Koszyk jest pusty</p>
            ) : (
                <ul>
                    {koszyk.map(item => (
                        <li key={item.id}>
                            {item.nazwa} - ilość: {item.ilosc}
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}

export default Koszyk;
