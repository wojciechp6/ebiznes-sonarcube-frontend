import React, { useEffect, useState } from "react";
import axios from "axios";

export default function Profile() {
  const [profile, setProfile] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return;
    axios.get("http://localhost:4000/profile", {
      headers: { Authorization: "Bearer " + token }
    }).then(res => setProfile(res.data));
  }, []);

  if (!profile) return <div>Ładowanie profilu...</div>;
  return (
    <div>
      <h2>Profil</h2>
      <pre>{JSON.stringify(profile, null, 2)}</pre>
      <button onClick={() => { localStorage.removeItem("token"); window.location = "/login"; }}>Wyloguj</button>
    </div>
  );
}
