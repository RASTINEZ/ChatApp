"use client";

import React, { useState, KeyboardEvent } from "react";
import BookingPage from "../components/BookingPage"; // import your booking component
import { sendMessage as sendChatMessage } from "../api/chat";
import Image from "next/image";

interface Message {
  sender: "user" | "bot";
  text: string;
  type?: "text" | "component";
  componentName?: string;
}

interface ChatWindowProps {
  room: string;
  setActiveView: React.Dispatch<
    React.SetStateAction<
      | "dashboard"
      | "chatHome"
      | "IT Group"
      | "Marketing Group"
      | "HR Chatbot"
      | "General Chatbot"
    >
  >;
}

const ChatWindow: React.FC<ChatWindowProps> = ({ room, setActiveView }) => {
  const [messages, setMessages] = useState<Message[]>([
    {
      sender: "bot",
      text: "",
      type: "text",
      componentName: "DefaultBanner",
    },
  ]);

  const [userInput, setUserInput] = useState<string>("");

  const sendMessage = async () => {
    if (userInput.trim() === "") return;

    const newMessages: Message[] = [
      ...messages,
      { sender: "user", text: userInput },
    ];
    setMessages(newMessages);

    try {
      const data = await sendChatMessage(userInput);

      if (data.response === "__show_booking__") {
        setMessages([
          ...newMessages,
          {
            sender: "bot",
            text: "I can help you with booking! Click the button below to proceed.",
            type: "text",
            componentName: "BookingButton",
          },
        ]);
      } else {
        setMessages([...newMessages, { sender: "bot", text: data.response }]);
      }
    } catch (error) {
      console.error("Error:", error);
    }

    setUserInput("");
  };

  const handleKeyPress = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      sendMessage();
    }
  };

  const renderMessage = (message: Message, index: number) => {
    // Show the BookingPage component
    if (
      message.type === "component" &&
      message.componentName === "BookingPage"
    ) {
      return <BookingPage key={index} embedded />;
    }

    // Show message with a booking button
    if (message.sender === "bot" && message.componentName === "BookingButton") {
      return (
        <div key={index} className="pb-4">
          <p>
            <strong>Bot:</strong> {message.text}
          </p>
          <button
            className="mt-2 px-4 py-2 bg-blue-500 text-white rounded"
            onClick={handleShowBookingComponent}
          >
            Book Now
          </button>
        </div>
      );
    }
    // if (message.sender === 'bot' && message.componentName === 'DefaultBanner') {
    //     return (
    //         <div key={index} className="sticky bottom-0 flex justify-center bg-white pt-4">
    //             <div
    //                 className="cursor-pointer transition transform hover:scale-105"
    //                 onClick={handleBannerClick}
    //             >
    //                 <Image
    //                     src="/assets/booking_Banner.jpg"
    //                     alt="Book a Room"
    //                     width={200}
    //                     height={100}
    //                     className="rounded shadow-lg"
    //                 />
    //             </div>
    //         </div>
    //     );
    // }

    // Default rendering
    return (
      <p key={index}>
        {message.sender === "user" && (
          <>
            <strong>User: </strong>
            {message.text}
          </>
        )}
        {message.sender === "bot" && (
          <>
            <strong>AI: </strong>{" "}
            {/* Or "Bot:" if you prefer consistency with BookingButton */}
            <span dangerouslySetInnerHTML={{ __html: message.text }} />
          </>
        )}
      </p>
    );
  };

  const handleShowBookingComponent = () => {
    setMessages((prevMessages) => [
      ...prevMessages,
      {
        sender: "bot",
        text: "",
        type: "component",
        componentName: "BookingPage",
      },
    ]);
  };

  const handleBannerClick = async () => {
    const bookMessage = "book";

    setMessages((prev) => [...prev, { sender: "user", text: bookMessage }]);
    setUserInput(""); // clear input

    try {
      const data = await sendChatMessage(bookMessage);

      if (data.response === "__show_booking__") {
        setMessages((prev) => [
          ...prev,
          {
            sender: "bot",
            text: "I can help you with booking! Click the button below to proceed.",
            type: "text",
            componentName: "BookingButton",
          },
        ]);
      } else {
        setMessages((prev) => [
          ...prev,
          { sender: "bot", text: data.response },
        ]);
      }
    } catch (err) {
      console.error("❌ Error from banner click:", err);
    }
  };

  return (
    <div className="p-4 w-full max-w-full md:max-w-5xl mx-auto relative">
      {/* Back Button */}
      <button
        className="mb-4 px-4 py-2 bg-gray-300 hover:bg-gray-400 text-black rounded"
        onClick={() => setActiveView("chatHome")}
      >
        ⬅ Back to Chats
      </button>

      <h4>{room} Chat</h4>

      {/* กล่องข้อความ */}
      <div className="chat-messages mt-3 h-[500px] overflow-y-auto border border-gray-300 p-4 rounded bg-white">
        {messages.map((message, index) => renderMessage(message, index))}
      </div>

      {/* Input */}
      <div className="mt-3 d-flex">
        <input
          type="text"
          value={userInput}
          onChange={(e) => setUserInput(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="Type your message..."
          className="form-control me-2"
        />
        <button onClick={sendMessage} className="btn btn-primary">
          Send
        </button>
      </div>

      {/* 🚨 Floating Banner */}
      {messages.length <= 1 && ( // แสดงเฉพาะตอนเปิดใหม่
        <div className="fixed bottom-55 left-1/2 transform -translate-x-1/2 z-50">
          <div
            className="cursor-pointer transition transform hover:scale-105"
            onClick={handleBannerClick}
          >
            <Image
              src="/assets/booking_Banner.jpg"
              alt="Book a Room"
              width={250}
              height={120}
              className="rounded shadow-lg"
            />
          </div>
        </div>
      )}
    </div>
  );
};

export default ChatWindow;
