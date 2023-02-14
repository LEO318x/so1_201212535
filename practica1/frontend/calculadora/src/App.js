import './App.css';

import {BrowserRouter, Routes, Route} from "react-router-dom"
import Calculator from './components/Calculator';
import Layout from './pages/Layout';
import Logs from './pages/Logs';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
        <Route path="calc" element={<Calculator />} />
        <Route path="logs" element={<Logs />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
