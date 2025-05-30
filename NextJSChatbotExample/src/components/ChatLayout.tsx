'use client';

import React, { useState } from 'react';
// import Sidebar from './Sidebar';
// import ChatWindow from './ChatWindow';

const ChatLayout: React.FC = () => {
  // const [activeRoom, setActiveRoom] = useState<string>('Marketing');
  const [showSidebar, setShowSidebar] = useState<boolean>(false);

  // const rooms: string[] = ['Marketing', 'Sales', 'IT'];

  return (
    <div className="flex flex-col md:flex-row h-screen">
      {/* Toggle button for mobile */}
      <div className="md:hidden p-2 bg-gray-100 border-b">
        <button
          className="px-4 py-2 bg-blue-600 text-white rounded"
          onClick={() => setShowSidebar(!showSidebar)}
        >
          {showSidebar ? 'Close Menu' : 'Open Menu'}
        </button>
      </div>

      {/* Sidebar — always visible on md+, toggle on mobile */}
      <div
        className={`w-full md:w-1/4 border-r bg-white ${showSidebar ? 'block' : 'hidden'
          } md:block`}
      >
        {/* <Sidebar
          rooms={rooms}
          activeRoom={activeRoom}
          setActiveRoom={(room) => {
            setActiveRoom(room);
            setShowSidebar(false); // auto-close sidebar on mobile when a room is selected
          }}
        /> */}
      </div>



      {/* ChatWindow */}
      <div className="flex-1 overflow-y-auto">
        {/* <ChatWindow room={activeRoom} /> */}
      </div>
    </div>
  );
};

export default ChatLayout;
