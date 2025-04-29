import React, { useEffect, useState } from "react";
import api from "../axiosConfig";

function ProductsPage({ cartId = 1 }) {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [adding, setAdding] = useState(null); // productId, jeśli właśnie dodajemy

    useEffect(() => {
        api.get("/products")
            .then(res => setProducts(res.data))
            .finally(() => setLoading(false));
    }, []);

    const handleAddToCart = async (productId) => {
        setAdding(productId);
        try {
            await api.post(`/carts/${cartId}/items`, {
                productId,
                quantity: 1
            });

        } catch (e) {
            alert("Błąd przy dodawaniu do koszyka");
        }
        setAdding(null);
    };

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Produkty</h1>
            <ul>
                {products.map(prod => (
                    <li key={prod.ID}>
                        <b>{prod.name}</b> — {prod.price} PLN{" "}
                        <button
                            onClick={() => handleAddToCart(prod.ID)}
                            disabled={adding === prod.ID}
                        >
                            {adding === prod.ID ? "Dodaję..." : "Dodaj do koszyka"}
                        </button>
                    </li>
                ))}
            </ul>
        </main>
    );
}

export default ProductsPage;