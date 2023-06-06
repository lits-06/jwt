import React, { useEffect, useState } from 'react';
import './App.css';
import Nav from './components/Nav';
import Login from './pages/Login';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Register from './pages/Register';
import Home from './pages/Home';

function App() {
  const [name, setName] = useState("")

  useEffect(() => {
    (
      async () => {
        const response = await fetch("http://localhost:8000/api/user", {
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          }
        });

        const content = await response.json();

        setName(content.name)
      }
    )();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />

        <main className="form-signin w-100 m-auto">
          <Routes>
            <Route path="/" Component={() => <Home name={name} />} />
            <Route path="/login" Component={() => <Login name={name} setName={setName} />} />
            <Route path="/register" Component={Register} />
          </Routes>
        </main>
      </BrowserRouter>
    </div >
  );
}

export default App;
