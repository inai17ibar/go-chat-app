import React, { useState, useEffect, useRef } from 'react';
import './App.css';

function App() {
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  useEffect(() => {
    ws.current = new WebSocket('ws://localhost:8000/ws');
    ws.current.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, msg]);
    };
    return () => {
      ws.current.close();
    };
  }, []);

  const handleSendMessage = () => {
    const msg = {
      email: 'user@example.com',
      username: 'User',
      message: message,
    };
    ws.current.send(JSON.stringify(msg));
    setMessage('');
  };

  return (
    <div className="App">
      <ul>
        {messages.map((msg, idx) => (
          <li key={idx}>
            <strong>{msg.username}:</strong> {msg.message}
          </li>
        ))}
      </ul>
      <input
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
}

export default App;
