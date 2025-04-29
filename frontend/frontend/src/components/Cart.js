import React, { useEffect, useState } from "react";
import api from "../axiosConfig";

function CartsPage() {
    const [carts, setCarts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [creating, setCreating] = useState(false);
    const [error, setError] = useState(null);

    // Pobieranie koszyków
    const fetchCarts = () => {
        setLoading(true);
        api.get("/carts")
            .then(res => setCarts(res.data))
            .catch(() => setCarts([]))
            .finally(() => setLoading(false));
    };

    useEffect(() => {
        fetchCarts();
    }, []);

    // Dodawanie nowego koszyka
    const handleCreateCart = async (e) => {
        e.preventDefault();
        setCreating(true);
        setError(null);
        try {
            await api.post("/carts", {}); // Jeśli potrzebujesz dodatkowych pól, dodaj je tutaj
            fetchCarts();
        } catch (err) {
            setError("Nie udało się utworzyć koszyka.");
        }
        setCreating(false);
    };

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Koszyki</h1>
            <ul>
                {carts.map(cart => (
                    <li key={cart.ID}>
                        Koszyk #{cart.ID} — {cart.items?.length ?? 0} produktów
                    </li>
                ))}
            </ul>
            <form onSubmit={handleCreateCart}>
                <button type="submit" disabled={creating}>
                    {creating ? "Tworzenie..." : "Dodaj nowy koszyk"}
                </button>
                {error && <div style={{color: "red"}}>{error}</div>}
            </form>
        </main>
    );
}

export default CartsPage;