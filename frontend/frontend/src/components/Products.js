import React, { useEffect, useState } from "react";
import api from "../axiosConfig";

function ProductsPage() {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        api.get("/products")
            .then(res => setProducts(res.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Produkty</h1>
            <ul>
                {products.map(prod => (
                    <li key={prod.id}>
                        <b>{prod.name}</b> — {prod.price} PLN
                    </li>
                ))}
            </ul>
        </main>
    );
}

export default ProductsPage;