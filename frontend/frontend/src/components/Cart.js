import React, { useEffect, useContext } from "react";
import api from "../axiosConfig";
import { CartContext } from "../context/CartContext";

function CartPage() {
    const { cart, loading, refreshCart } = useContext(CartContext);

    const handleRemoveFromCart = async (itemId) => {
        try {
            await api.delete(`/carts/current/items/${itemId}`); // lub inny endpoint zgodny z backendem
            refreshCart();
        } catch (err) {
            alert("Błąd usuwania z koszyka.");
        }
    };

    if (loading) return <div>Loading...</div>;
    if (!cart) return <div>Koszyk jest pusty.</div>;

    return (
        <main>
            <h1>Koszyk</h1>
            <ul>
                {cart.items && cart.items.length > 0 ? (
                    cart.items.map(item => (
                        <li key={item.id}>
                            {item.name} — {item.price} PLN&nbsp;
                            <button onClick={() => handleRemoveFromCart(item.id)}>
                                Usuń
                            </button>
                        </li>
                    ))
                ) : (
                    <li>Koszyk jest pusty.</li>
                )}
            </ul>
        </main>
    );
}

export default CartPage;