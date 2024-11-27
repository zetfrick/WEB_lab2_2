import React, { useState } from 'react';
import axios from 'axios';
import './AddNews.css'; // Импортируем CSS файл

const AddNews = () => {
  // Храненим значения полей
  const [title, setTitle] = useState('');
  const [author, setAuthor] = useState('');
  const [content, setContent] = useState('');

  // Обработчик отправки формы
  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/news', { title, author, content })
      .then(response => {
        console.log('News added:', response.data); // Логи успехи
      })
      .catch(error => {
        console.error('There was an error adding the news!', error); // Логи ошибок
      });
  };

  return (
    <div>
      <h1>Add News</h1>
      <form className="add-news-form" onSubmit={handleSubmit}>
        <div>
          <label>Title:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)} // Обновление состояние title при изменении поля
          />
        </div>
        <div>
          <label>Author:</label>
          <input
            type="text"
            value={author}
            onChange={(e) => setAuthor(e.target.value)} // Обновление состояние author при изменении поля
          />
        </div>
        <div>
          <label>Content:</label>
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)} // Обновление состояние content при изменении поля
          />
        </div>
        <button type="submit">Add News</button>
      </form>
    </div>
  );
};

export default AddNews;
