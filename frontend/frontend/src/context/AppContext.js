import React, { createContext, useContext, useReducer } from "react";

const CartContext = createContext();

const initialState = {
    items: [],
};

function cartReducer(state, action) {
    switch (action.type) {
        case "ADD_ITEM":
            return { ...state, items: [...state.items, action.payload] };
        case "REMOVE_ITEM":
            return { ...state, items: state.items.filter(i => i.id !== action.payload) };
        case "CLEAR_CART":
            return { ...state, items: [] };
        default:
            return state;
    }
}

export function CartProvider({ children }) {
    const [state, dispatch] = useReducer(cartReducer, initialState);

    const addItem = (item) => dispatch({ type: "ADD_ITEM", payload: item });
    const removeItem = (id) => dispatch({ type: "REMOVE_ITEM", payload: id });
    const clearCart = () => dispatch({ type: "CLEAR_CART" });

    return (
        <CartContext.Provider value={{ ...state, addItem, removeItem, clearCart }}>
            {children}
        </CartContext.Provider>
    );
}

export const useCart = () => useContext(CartContext);