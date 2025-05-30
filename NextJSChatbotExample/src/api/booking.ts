// src/api/booking.ts
import axios from 'axios';

const BASE_URL = `${process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '')}/api`;

export interface Booking {
  id: number;
  room_name: string;
  date: string;
  time_label: string;
  created_at: string;
}

export const fetchRooms = async () => {
  const res = await axios.get(`${BASE_URL}/rooms`);
  return res.data;
};

export const fetchTimeSlots = async () => {
  const res = await axios.get(`${BASE_URL}/timeslots`);
  return res.data;
};

export const submitBooking = async (booking: {
  room_id: number;
  timeslot_id: number;
  date: string;
}) => {
  const res = await axios.post(`${BASE_URL}/bookings`, booking);
  return res.data;
};

export const fetchBookings = async (): Promise<Booking[]> => {
  try {
    const res = await axios.get<Booking[]>(`${BASE_URL}/bookings/all`);
    return res.data;
  } catch (err) {
    console.error('Failed to fetch bookings:', err);
    return [];  // Return empty array on error to avoid undefined issues
  }
};