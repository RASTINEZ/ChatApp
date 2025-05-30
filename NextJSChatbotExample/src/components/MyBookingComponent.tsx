'use client';

import React, { useEffect, useState } from 'react';

import { fetchBookings, Booking } from '../api/booking';



const MyBookings: React.FC = () => {
    const [bookings, setBookings] = useState<Booking[]>([]);

    useEffect(() => {
      const loadBookings = async () => {
        const data = await fetchBookings();
        setBookings(data);
      };
  
      loadBookings();
    }, []);

    return (
        <div>
          <h4 className="mb-3">📋 Your Bookings</h4>
          {bookings.length === 0 ? (
            <p>No bookings found.</p>
          ) : (
            <ul className="list-group">
              {bookings.map((booking) => (
                <li key={booking.id} className="list-group-item">
                  <strong>{booking.room_name}</strong> — {booking.date} at {booking.time_label}
                  <br />
                  <small className="text-muted">Booked on: {new Date(booking.created_at).toLocaleString()}</small>
                </li>
              ))}
            </ul>
          )}
        </div>
      );
    };
    
    export default MyBookings;
