// src/api/chat.ts
import axios from 'axios';

const API_URL = `${process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '')}/chat`;

// Send message API
export const sendMessage = async (message: string) => {
  try {
    const response = await axios.post(API_URL, { message });
    return response.data;
  } catch (error) {
    console.error('Error sending message:', error);
    throw error;  // You can handle error more specifically if needed
  }
};
