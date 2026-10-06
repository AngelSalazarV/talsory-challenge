"use client";

import { useState } from "react";

type QRResponse = {
  q: number[][];
  r: number[][];
  stats: {
    max: number;
    min: number;
    average: number;
    total: number;
    isDiagonal: boolean;
  };
};

const GO_API_URL =
  process.env.NEXT_PUBLIC_GO_API_URL ?? "http://localhost:3000"

export default function Home() {
  const [username, setUsername] = useState("admin");
  const [password, setPassword] = useState("admin123");

  const [matrixText, setMatrixText] = useState(
    "[[1,2],[3,4],[5,6]]"
  );

  const [token, setToken] = useState("");
  const [result, setResult] = useState<QRResponse | null>(null);
  const [error, setError] = useState("");

  async function login() {
    setError("");

    const response = await fetch(`${GO_API_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username,
        password,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      setError(data.error);
      return;
    }

    setToken(data.token);
  }

  async function calculateQR() {
    setError("");

    if (!token) {
      setError("Primero debes iniciar sesión");
      return;
    }

    const matrix = JSON.parse(matrixText);

    const response = await fetch(`${GO_API_URL}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        matrix,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      setError(data.error);
      return;
    }

    setResult(data);
  }

  return (
    <main className="min-h-screen p-8">
      <div className="mx-auto max-w-4xl">
        <h1 className="mb-8 text-3xl font-bold">
          Matrix QR Calculator
        </h1>

        <section className="mb-6 rounded-lg border p-6">
          <h2 className="mb-4 text-xl font-semibold">
            Authentication
          </h2>

          <div className="flex gap-3">
            <input
              className="rounded border p-2"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Username"
            />

            <input
              className="rounded border p-2"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Password"
            />

            <button
              className="rounded bg-black px-4 py-2 text-white cursor-pointer hover:bg-gray-700"
              onClick={login}
            >
              Login
            </button>
          </div>

          {token && (
            <p className="mt-3 text-green-600">
              Authenticated successfully
            </p>
          )}
        </section>

        <section className="mb-6 rounded-lg border p-6">
          <h2 className="mb-4 text-xl font-semibold">
            Matrix
          </h2>

          <textarea
            className="h-32 w-full rounded border p-3 font-mono"
            value={matrixText}
            onChange={(e) => setMatrixText(e.target.value)}
          />

          <button
            className="mt-4 rounded bg-blue-600 px-5 py-2 text-white cursor-pointer hover:bg-blue-500"
            onClick={calculateQR}
          >
            Calculate QR
          </button>
        </section>

        {error && (
          <div className="mb-6 rounded border border-red-300 p-4 text-red-600">
            {error}
          </div>
        )}

        {result && (
          <section className="rounded-lg border p-6">
            <h2 className="mb-4 text-xl font-semibold">
              Results
            </h2>

            <h3 className="font-semibold">Q</h3>
            <pre className="mb-6 rounded bg-gray-100 p-4">
              {JSON.stringify(result.q, null, 2)}
            </pre>

            <h3 className="font-semibold">R</h3>
            <pre className="mb-6 rounded bg-gray-100 p-4">
              {JSON.stringify(result.r, null, 2)}
            </pre>

            <h3 className="font-semibold">Statistics</h3>

            <div className="mt-3 grid grid-cols-2 gap-3">
              <p>Max: {result.stats.max}</p>
              <p>Min: {result.stats.min}</p>
              <p>Average: {result.stats.average}</p>
              <p>Total: {result.stats.total}</p>
              <p>
                Diagonal: {result.stats.isDiagonal ? "Yes" : "No"}
              </p>
            </div>
          </section>
        )}
      </div>
    </main>
  );
}