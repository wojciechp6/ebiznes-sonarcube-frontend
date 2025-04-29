import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import ProductsPage from "./components/Products";
import CategoriesPage from "./components/Categories";
import CartPage from "./components/Cart";
import PaymentsPage from "./components/Payment";
import { CartProvider } from "./context/CartContext";

function App() {
    return (
        <CartProvider>
            <Router>
                <Navbar />
                <Routes>
                    <Route path="/products" element={<ProductsPage />} />
                    <Route path="/categories" element={<CategoriesPage />} />
                    <Route path="/carts" element={<CartPage />} />
                    <Route path="/payments" element={<PaymentsPage />} />
                </Routes>
            </Router>
        </CartProvider>
    );
}
export default App;