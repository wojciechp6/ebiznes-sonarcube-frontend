import React, { createContext, useState, useCallback } from "react";
import api from "../axiosConfig";

export const CartContext = createContext();

export function CartProvider({ children }) {
    const [cart, setCart] = useState(null);
    const [loading, setLoading] = useState(true);

    const fetchCart = useCallback(() => {
        setLoading(true);
        api.get("/carts/current")
            .then(res => setCart(res.data))
            .catch((err) => {
                setCart(null);
                if (process.env.NODE_ENV !== 'production') {
                    // eslint-disable-next-line no-console
                    console.error('Fetch cart error:', err);
                }
            })
            .finally(() => setLoading(false));
    }, []);

    React.useEffect(() => { fetchCart(); }, [fetchCart]);

    const refreshCart = fetchCart;

    return (
        <CartContext.Provider value={{ cart, loading, refreshCart }}>
            {children}
        </CartContext.Provider>
    );
}