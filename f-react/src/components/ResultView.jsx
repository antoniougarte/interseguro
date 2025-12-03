export default function ResultView({ data, onBack }) {
  return (
    <div>
      <h2>Resultado</h2>

      <h3>Matriz Rotada</h3>
      <pre>{JSON.stringify(data.rotated, null, 2)}</pre>

      <h3>Matriz Q</h3>
      <pre>{JSON.stringify(data.Q, null, 2)}</pre>

      <h3>Matriz R</h3>
      <pre>{JSON.stringify(data.R, null, 2)}</pre>

      <h3>Stats</h3>
      <pre>{JSON.stringify(data.stats, null, 2)}</pre>

      <button onClick={onBack}>↩ Volver</button>
    </div>
  );
}
