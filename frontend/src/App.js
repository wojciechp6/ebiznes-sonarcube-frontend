// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Produkty from './components/Produkty';
import Koszyk from './components/Koszyk';
import Platnosci from './components/Platnosci';

function App() {
  // Na początek przykładowy stan koszyka (później zastąpimy globalnym stanem)
  const [produktyWKoszyku, setProduktyWKoszyku] = React.useState([]);

  return (
      <Router>
        <nav>
          <ul>
            <li><Link to="/">Produkty</Link></li>
            <li><Link to="/koszyk">Koszyk</Link></li>
            <li><Link to="/platnosci">Płatności</Link></li>
          </ul>
        </nav>
        <Routes>
          <Route path="/" element={<Produkty />} />
          <Route path="/koszyk" element={<Koszyk produktyWKoszyku={produktyWKoszyku} />} />
          <Route path="/platnosci" element={<Platnosci />} />
        </Routes>
      </Router>
  );
}

export default App;
