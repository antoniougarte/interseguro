import { useState } from "react";
import LoginForm from "./components/LoginForm";
import MatrixForm from "./components/MatrixForm";
import ResultView from "./components/ResultView";

function App() {
  const [token, setToken] = useState(null);
  const [result, setResult] = useState(null);

  return (
    <div style={{ padding: 30, fontFamily: "Arial" }}>
      <h1>Matrix Processor</h1>

      {!token && (
        <LoginForm onLogin={(jwt) => setToken(jwt)} />
      )}

      {token && !result && (
        <MatrixForm token={token} onResult={(res) => setResult(res)} />
      )}

      {token && result && (
        <ResultView data={result} onBack={() => setResult(null)} />
      )}
    </div>
  );
}

export default App;
