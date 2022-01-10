import { FormEventHandler, useState } from "react";

function App() {
  const [password, setPassword] = useState("");
  const [passwordLength, setPasswordLength] = useState(6);
  const [withNumbers, setWithNumbers] = useState(false);

  const generatePassword: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();

    fetch("http://localhost:8080/new", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        length: passwordLength,
        with_numbers: withNumbers,
      }),
    })
      .then((res) => res.json())
      .then(({ password }) => setPassword(password))
      .catch((error) => console.error(error));
  };

  return (
    <div className="container">
      <main>
        <h1>Password Generator</h1>
        <article>
          New password: <strong>{password || "..."}</strong>
        </article>
        <form onSubmit={generatePassword}>
          <fieldset>
            <label htmlFor="length">Length</label>
            <input
              type="number"
              name="length"
              value={passwordLength}
              onChange={(e) => setPasswordLength(parseInt(e.target.value))}
            />
          </fieldset>
          <fieldset>
            <label htmlFor="numbers">With numbers</label>
            <input
              type="checkbox"
              name="numbers"
              checked={withNumbers}
              onChange={(e) => setWithNumbers(e.target.checked)}
            />
          </fieldset>
          <button>Generate new password</button>
        </form>
      </main>
    </div>
  );
}

export default App;
