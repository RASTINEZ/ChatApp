'use client';

import React, { useState } from 'react';
import Sidebar from './Sidebar';
import ChatWindow from './ChatWindow';
import MainMenu from './MainMenu';
import ChatHome from './ChatHome';
import GroupChatLayout from './GroupChatLayout';

const MainLayout: React.FC = () => {
    const [showSidebar, setShowSidebar] = useState<boolean>(false);

    // Expand activeView to handle all cases
    const [activeView, setActiveView] = useState<
        'dashboard' | 'chatHome' | 'IT Group' | 'Marketing Group' | 'HR Chatbot' | 'General Chatbot'
    >('dashboard');

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

            {/* Sidebar */}
            <div className={`w-full md:w-1/4 border-r bg-white ${showSidebar ? 'block' : 'hidden'} md:block`}>
                <Sidebar activeView={activeView} setActiveView={setActiveView} />
            </div>

            {/* Main Content */}
            <div className="flex-1 overflow-y-auto p-4">
                {activeView === 'dashboard' && <MainMenu />}
                {activeView === 'chatHome' && <ChatHome setActiveView={setActiveView} />}

                {['IT Group', 'Marketing Group'].includes(activeView) && (
                    <GroupChatLayout onBack={() => setActiveView('chatHome')} />
                )}

                {['HR Chatbot', 'General Chatbot'].includes(activeView) && <ChatWindow room={activeView} setActiveView={setActiveView} />}


            </div>
        </div>
    );
};

export default MainLayout;
