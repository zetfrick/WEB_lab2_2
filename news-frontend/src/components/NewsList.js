import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './NewsList.css'; // Импортируем CSS файл

const NewsList = () => {
  // Хранение новостей
  const [news, setNews] = useState([]);
  const [expandedIds, setExpandedIds] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/news')
      .then(response => {
        setNews(response.data); // Обновление новостей
      })
      .catch(error => {
        console.error('There was an error fetching the news!', error); // Логи ошибок
      });
  }, []);

  const toggleExpand = (id) => {
    setExpandedIds(prevIds =>
      prevIds.includes(id) ? prevIds.filter(itemId => itemId !== id) : [...prevIds, id]
    );
  };

  return (
    <div className="news-list-container">
      <h1>News List</h1>
      <ul>
        {news.map(item => (
          <li
            key={item.id}
            className={`news-list-item ${expandedIds.includes(item.id) ? 'expanded' : ''}`}
            onClick={() => toggleExpand(item.id)}
          >
            <span className="news-id">ID: {item.id}</span>
            <h2>{item.title}</h2>
            <p>Author: {item.author}</p>
            <p className="content">{item.content}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default NewsList;
