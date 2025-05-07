import React, { useState } from "react";
import axios from "axios";

export default function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");

  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:4000/auth/register", { username, password, email, name });
      setSuccess(true);
      setError("");
    } catch (err) {
      setError("Błąd rejestracji");
    }
  };

  return (
    <div>
      <h2>Rejestracja</h2>
      <form onSubmit={handleRegister}>
        <input value={username} onChange={e => setUsername(e.target.value)} placeholder="Login" />
        <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Hasło" />
        <input value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" />
        <input value={name} onChange={e => setName(e.target.value)} placeholder="Imię i nazwisko" />
        <button type="submit">Zarejestruj</button>
      </form>
      {success && <div style={{color: "green"}}>Rejestracja udana! Możesz się zalogować.</div>}
      {error && <div style={{color: "red"}}>{error}</div>}
    </div>
  );
}
