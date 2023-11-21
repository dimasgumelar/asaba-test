import "./App.css";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Barang from "./Pages/Barang/Barang";
import BarangAdd from "./Pages/Barang/BarangAdd";
import History from "./Pages/History/History";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Barang />} />
          <Route path="/barang" element={<Barang />} />
          <Route path="/barang/add" element={<BarangAdd />} />
          <Route path="/history" element={<History />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
