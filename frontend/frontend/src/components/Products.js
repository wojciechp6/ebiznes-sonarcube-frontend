import React, { useEffect, useState, useContext } from "react";
import api from "../axiosConfig";
import { CartContext } from "../context/CartContext";

function ProductsPage() {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const { refreshCart } = useContext(CartContext);

    useEffect(() => {
        api.get("/products")
            .then(res => setProducts(res.data))
            .finally(() => setLoading(false));
    }, []);

    const handleAddToCart = async (productId) => {
        try {
            await api.post("/carts/current/items", { productId });
            refreshCart();
            alert("Dodano do koszyka!");
        } catch (err) {
            alert("Błąd dodawania do koszyka.");
        }
    };

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Produkty</h1>
            <ul>
                {products.map(prod => (
                    <li key={prod.id}>
                        <b>{prod.name}</b> — {prod.price} PLN&nbsp;
                        <button onClick={() => handleAddToCart(prod.id)}>
                            Dodaj do koszyka
                        </button>
                    </li>
                ))}
            </ul>
        </main>
    );
}

export default ProductsPage;