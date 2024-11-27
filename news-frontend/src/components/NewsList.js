import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './NewsList.css'; // Импортируем CSS файл

const NewsList = () => {
  // Хранение новостей
  const [news, setNews] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/news')
      .then(response => {
        setNews(response.data); // Обновление новостей
      })
      .catch(error => {
        console.error('There was an error fetching the news!', error); // Логи ошибок
      });
  }, []);

  return (
    <div>
      <h1>News List</h1>
      <ul>
        {news.map(item => (
          <li key={item.id}>
            <span className="news-id">ID: {item.id}</span>
            <h2>{item.title}</h2>
            <p>Author: {item.author}</p>
            <p>{item.content}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default NewsList;
