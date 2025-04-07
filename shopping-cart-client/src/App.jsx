import './App.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Homepage from './pages/Homepage';
import CartPage from './pages/CartPage';
import MenuBar from './components/MenuBar';

export default function App() {
  return (
    <>
      <MenuBar />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Homepage />}/>
          <Route path="/cart" element={<CartPage />}/>
        </Routes>
      </BrowserRouter>
    </>
  );
}
