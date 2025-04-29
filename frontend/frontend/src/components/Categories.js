import React, { useEffect, useState } from "react";
import api from "../axiosConfig";

function CategoriesPage() {
    const [categories, setCategories] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        api.get("/categories")
            .then(res => setCategories(res.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <div>Loading...</div>;

    return (
        <main>
            <h1>Kategorie</h1>
            <ul>
                {categories.map(cat => (
                    <li key={cat.id}>{cat.name}</li>
                ))}
            </ul>
        </main>
    );
}

export default CategoriesPage;