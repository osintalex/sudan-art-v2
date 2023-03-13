import logo from "./logo.svg";
import "./App.css";
import { useState } from "react";

function App() {
  const [response, setResponse] = useState("");

  async function callLambdaFunc() {
    const apiResponse = await fetch("/api/v1");
    const responseText = await apiResponse.text();
    setResponse(responseText);
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={callLambdaFunc}>click me friend</button>
        <p>
          Api response is: <i>{response}</i>
        </p>
      </header>
    </div>
  );
}

export default App;
