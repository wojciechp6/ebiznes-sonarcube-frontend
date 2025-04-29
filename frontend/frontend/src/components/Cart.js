import React, { useEffect, useState } from "react";
import api from "../axiosConfig";

function CartsPage() {
    const [carts, setCarts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        api.get("/carts")
            .then(res => setCarts(res.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Koszyki</h1>
            <ul>
                {carts.map(cart => (
                    <li key={cart.id}>Koszyk #{cart.id} — {cart.items.length} produktów</li>
                ))}
            </ul>
        </main>
    );
}

export default CartsPage;