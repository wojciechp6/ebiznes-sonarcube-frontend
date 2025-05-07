import React, { useState } from "react";
import axios from "axios";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:4000/auth/login", { username, password });
      localStorage.setItem("token", res.data.token);
      window.location = "/profile";
    } catch (err) {
      setError("Błędny login lub hasło");
    }
  };

  const handleOAuth = (provider) => {
    window.location = `http://localhost:4000/auth/${provider}`;
  };

  return (
    <div>
      <h2>Logowanie</h2>
      <form onSubmit={handleLogin}>
        <input value={username} onChange={e => setUsername(e.target.value)} placeholder="Login" />
        <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Hasło" />
        <button type="submit">Zaloguj</button>
      </form>
      {error && <div style={{color: "red"}}>{error}</div>}
      <button onClick={() => handleOAuth("google")}>Zaloguj przez Google</button>
      <button onClick={() => handleOAuth("github")}>Zaloguj przez Github</button>
    </div>
  );
}
