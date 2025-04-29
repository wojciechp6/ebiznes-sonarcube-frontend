import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import ProductsPage from "./components/Products";
import CategoriesPage from "./components/Categories";
import CartsPage from "./components/Cart";
import Navbar from "./components/Navbar";
import PaymentsPage from "./components/Payment";

function App() {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route path="/" element={<Navigate to="/products" />} />
                <Route path="/products" element={<ProductsPage />} />
                <Route path="/categories" element={<CategoriesPage />} />
                <Route path="/carts" element={<CartsPage />} />
                <Route path="/payments" element={<PaymentsPage />} />
            </Routes>
        </Router>
    );
}

export default App;