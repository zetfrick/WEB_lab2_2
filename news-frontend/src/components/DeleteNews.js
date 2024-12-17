import React, { useState } from 'react';
import axios from 'axios';
import './DeleteNews.css'; // Импортируем CSS файл

const DeleteNews = () => {
  // ID новости, которую нужно удалить
  const [id, setId] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.delete(`http://localhost:8080/news/${id}`)
      .then(response => {
        console.log('News deleted:', response.data); // Логи успехов
      })
      .catch(error => {
        console.error('There was an error deleting the news!', error); // Логи ошибок
      });
  };

  return (
    <div>
      <h1>Delete News</h1>
      <form className="delete-news-form" onSubmit={handleSubmit}>
        <div>
          <label>ID:</label>
          <input
            type="text"
            value={id}
            onChange={(e) => setId(e.target.value)}
          />
        </div>
        <button type="submit">Delete News</button>
      </form>
    </div>
  );
};

export default DeleteNews;
