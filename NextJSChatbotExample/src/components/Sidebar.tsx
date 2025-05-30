'use client';

import React, { useState } from 'react';

interface SidebarProps {
  activeView: string;
  setActiveView: React.Dispatch<React.SetStateAction<
    'dashboard' | 'chatHome' | 'IT Group' | 'Marketing Group' | 'HR Chatbot' | 'General Chatbot'
  >>;
}

const Sidebar: React.FC<SidebarProps> = ({ activeView, setActiveView }) => {
  const [showGroupChats, setShowGroupChats] = useState<boolean>(true);
  const [showBotChats, setShowBotChats] = useState<boolean>(true);

  return (
    <div className="bg-light border-end p-3 h-screen overflow-y-auto">
      <h5 className="mb-4">Menu</h5>
      <ul className="list-group">
        {/* Main Menu */}
        <li
          className={`list-group-item ${activeView === 'dashboard' ? 'active' : ''}`}
          onClick={() => setActiveView('dashboard')}
          style={{ cursor: 'pointer' }}
        >
          🏠 Main Menu
        </li>

        {/* Chat Home */}
        <li
          className={`list-group-item ${activeView === 'chatHome' ? 'active' : ''}`}
          onClick={() => setActiveView('chatHome')}
          style={{ cursor: 'pointer' }}
        >
          💬 Chat
        </li>

        {/* Group Chats Dropdown */}
        <li
          className="list-group-item bg-secondary text-white d-flex justify-content-between align-items-center"
          onClick={() => setShowGroupChats(!showGroupChats)}
          style={{ cursor: 'pointer' }}
        >
          📂 Group Chats
          <span>{showGroupChats ? '▼' : '►'}</span>
        </li>
        {showGroupChats && (
          <>
            <li
              className="list-group-item ms-3"
              onClick={() => setActiveView('IT Group')}
              style={{ cursor: 'pointer' }}
            >
              # IT Group
            </li>
            <li
              className="list-group-item ms-3"
              onClick={() => setActiveView('Marketing Group')}
              style={{ cursor: 'pointer' }}
            >
              # Marketing Group
            </li>
          </>
        )}

        {/* Bot Chats Dropdown */}
        <li
          className="list-group-item bg-secondary text-white d-flex justify-content-between align-items-center mt-2"
          onClick={() => setShowBotChats(!showBotChats)}
          style={{ cursor: 'pointer' }}
        >
          🤖 Bot Chats
          <span>{showBotChats ? '▼' : '►'}</span>
        </li>
        {showBotChats && (
          <>
            <li
              className="list-group-item ms-3"
              onClick={() => setActiveView('HR Chatbot')}
              style={{ cursor: 'pointer' }}
            >
              # HR Chatbot
            </li>
            <li
              className="list-group-item ms-3"
              onClick={() => setActiveView('General Chatbot')}
              style={{ cursor: 'pointer' }}
            >
              # General Chatbot
            </li>
          </>
        )}
      </ul>
    </div>
  );
};

export default Sidebar;
