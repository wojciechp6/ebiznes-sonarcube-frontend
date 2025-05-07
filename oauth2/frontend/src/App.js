import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Login from "./Login";
import Register from "./Register";
import OAuth2Redirect from "./OAuth2Redirect";
import Profile from "./Profile";

function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link to="/login">Logowanie</Link> | <Link to="/register">Rejestracja</Link>
      </nav>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/oauth2" element={<OAuth2Redirect />} />
        <Route path="/profile" element={<Profile />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
