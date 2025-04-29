import React from "react";
import { Link } from "react-router-dom";

function Navbar() {
    return (
        <nav>
            <Link to="/products">Produkty</Link>
            <Link to="/categories">Kategorie</Link>
            <Link to="/carts">Koszyki</Link>
            <Link to="/payments">Płatności</Link>
        </nav>
    );
}

export default Navbar;