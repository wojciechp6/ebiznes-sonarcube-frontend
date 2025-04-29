// src/context/AppContext.js
import React, { createContext, useState } from 'react';

export const AppContext = createContext();

export const AppProvider = ({ children }) => {
    const [produkty, setProdukty] = useState([]); // Lista produktów pobrana z serwera
    const [koszyk, setKoszyk] = useState([]);       // Lista produktów wybranych do zamówienia

    // Funkcja do dodawania produktu do koszyka
    const dodajDoKoszyka = (produkt) => {
        setKoszyk(prev => {
            const istnieje = prev.find(item => item.id === produkt.id);
            if (istnieje) {
                return prev.map(item =>
                    item.id === produkt.id
                        ? { ...item, ilosc: item.ilosc + 1 }
                        : item
                );
            } else {
                return [...prev, { ...produkt, ilosc: 1 }];
            }
        });
    };

    return (
        <AppContext.Provider value={{ produkty, setProdukty, koszyk, setKoszyk, dodajDoKoszyka }}>
            {children}
        </AppContext.Provider>
    );
};
