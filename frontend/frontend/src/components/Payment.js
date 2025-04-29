import React, { useState } from "react";
import api from "../axiosConfig";

function PaymentsPage({ cartId = 1 }) {
    const [amount, setAmount] = useState("");
    const [status, setStatus] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setStatus(null);
        try {
            await api.post("/payments", { cart_id: cartId, amount: parseFloat(amount) });
            setStatus("success");
        } catch (err) {
            setStatus("error");
        }
    };

    return (
        <main>
            <h1>Płatność</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Kwota:
                    <input
                        type="number"
                        min="0"
                        step="0.01"
                        value={amount}
                        onChange={e => setAmount(e.target.value)}
                        required
                    />
                </label>
                <button type="submit">Zapłać</button>
            </form>
            {status === "success" && <p style={{ color: "green" }}>Płatność zakończona sukcesem!</p>}
            {status === "error" && <p style={{ color: "red" }}>Błąd płatności.</p>}
        </main>
    );
}

export default PaymentsPage;