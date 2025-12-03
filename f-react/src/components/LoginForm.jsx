import { useState } from "react";

export default function LoginForm({ onLogin }) {
  const [username, setUsername] = useState("admin");
  const [pass, setPass] = useState("admin123");
  const [error, setError] = useState(null);

  const login = async () => {
    try {
      const res = await fetch(`http://18.188.99.23:4000/api/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password: pass })
      });

      const data = await res.json();
      console.log(data);
      if (data.success && data.data.token) {
        onLogin(data.data.token);
      } else {
        setError("Credenciales incorrectas");
      }
    } catch (e) {
      setError("Error conectando con el servidor");
    }
  };

  return (
    <div>
      <h2>Login</h2>

      <input value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" />
      <br />

      <input value={pass} onChange={e => setPass(e.target.value)} placeholder="Password" type="password" />
      <br />

      <button onClick={login}>Ingresar</button>

      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}
