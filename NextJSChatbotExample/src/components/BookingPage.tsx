'use client';

import React, { useState } from 'react';
import BookingLayout from '../components/BookingLayout';
import BookingDetails from './BookingDetails';



interface BookingPageProps {
    embedded?: boolean;
  }
  
  const BookingPage: React.FC<BookingPageProps> = ({ embedded = false }) => {
  const [selectedDate, setSelectedDate] = useState<string | null>(null);

  const handleDateClick = (date: string) => {
    setSelectedDate(date);
  };

  return (
    <div className={`${embedded ? '' : 'p-4'}`}>
      {!embedded && <h1 className="text-center my-4">Booking Page</h1>}
  
      <div className={`${embedded ? '' : 'flex gap-4'}`}>
        {/* Calendar */}
        <div className={`${embedded ? '' : 'w-1/2'}`}>
          <BookingLayout onDateClick={handleDateClick} />
        </div>
  
        {/* Booking Details */}
        <div className={`${embedded ? 'mt-4' : 'w-1/2 border rounded-lg p-4'}`}>
          {selectedDate ? (
            <>
              
              <BookingDetails selectedDate={selectedDate} />

            </>
          ) : (
            <p>Select a date to view or make a booking.</p>
          )}
        </div>
      </div>
    </div>
  );
  
};

export default BookingPage;
