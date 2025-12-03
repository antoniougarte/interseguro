import { useState } from "react";

export default function MatrixForm({ token, onResult }) {
  const [inputText, setInputText] = useState('[[1,2,3],[4,5,6],[7,8,9]]');
  const [error, setError] = useState(null);

  const submitMatrix = async () => {
    try {
      const matrix = JSON.parse(inputText);

      const res = await fetch(`http://18.188.99.23:4000/api/matrix/rotate`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({ matrix })
      });

      const data = await res.json();
      if (data.success) {
        onResult(data.data);
      } else {
        setError("Error en la operación");
      }
    } catch (e) {
      console.log(e);
      setError("Formato de matriz inválido");
    }
  };

  return (
    <div>
      <h2>Enviar Matriz</h2>
      <p>La matriz debe enviarse correctamente (mismo número de columnas en cada fila) para mostrar los resultados</p>

      <textarea
        rows="5"
        cols="40"
        value={inputText}
        onChange={(e) => setInputText(e.target.value)}
      />

      <br />
      <button onClick={submitMatrix}>Procesar</button>

      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}
