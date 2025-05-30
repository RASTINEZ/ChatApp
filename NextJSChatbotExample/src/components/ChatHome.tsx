interface ChatHomeProps {
    setActiveView: React.Dispatch<React.SetStateAction<
        'dashboard' | 'chatHome' | 'IT Group' | 'Marketing Group' | 'HR Chatbot' | 'General Chatbot'
    >>;
}

const ChatHome: React.FC<ChatHomeProps> = ({ setActiveView }) => {
    const recentChats = [
        { id: 'IT Group', name: '🖥️ IT Group' },
        { id: 'Marketing Group', name: '📊 Marketing Group' },
        { id: 'HR Chatbot', name: '🤖 HR Chatbot' },
        { id: 'General Chatbot', name: '🤖 General Chatbot' },
    ];
    



    return (
        <div>
            <h2 className="text-2xl font-semibold mb-4">💬 Chats</h2>
            <ul className="list-group">
                {recentChats.map(chat => (
                    <li
                        key={chat.id}
                        className="list-group-item hover:bg-gray-100 cursor-pointer"
                        onClick={() => setActiveView(chat.id as 
                            'IT Group' | 'Marketing Group' | 'HR Chatbot' | 'General Chatbot')}
                          
                       
                    >
                        {chat.name}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ChatHome;
