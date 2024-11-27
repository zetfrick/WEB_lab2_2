import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import NewsList from './components/NewsList';
import AddNews from './components/AddNews';
import DeleteNews from './components/DeleteNews';
import './App.css'; // Импортируем CSS файл

function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li><Link to="/">News List</Link></li> {/* Ссылка на новости */}
            <li><Link to="/add">Add News</Link></li> {/* Ссылка на добавления новости */}
            <li><Link to="/delete">Delete News</Link></li> {/* Ссылка на удаления новости */}
          </ul>
        </nav>
        <Routes>
          <Route path="/" element={<NewsList />} /> {/* Маршрут для отображения новостей */}
          <Route path="/add" element={<AddNews />} /> {/* Маршрут для отображения добавления новости */}
          <Route path="/delete" element={<DeleteNews />} /> {/* Маршрут для отображения удаления новости */}
        </Routes>
      </div>
    </Router>
  );
}

export default App;
