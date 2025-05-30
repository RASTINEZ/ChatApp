'use client';
import React, { useState } from 'react';
import FullCalendar from '@fullcalendar/react';
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import 'bootstrap/dist/css/bootstrap.min.css';
import MyBookings from './MyBookingComponent';

interface BookingLayoutProps {
  onDateClick: (date: string) => void;
}

const BookingLayout: React.FC<BookingLayoutProps> = ({ onDateClick }) => {
  const [activeTab, setActiveTab] = useState<'calendar' | 'myBookings'>('calendar');

  return (
    <div className="container-fluid">
      <div className="row flex-column flex-md-row">
        {/* Sidebar */}
        <div className="col-12 col-md-3 p-3 bg-light border-bottom border-md-end">
          <h5>My Bookings</h5>
          <ul className="list-group">
            <li
              className={`list-group-item ${activeTab === 'calendar' ? 'active' : ''}`}
              onClick={() => setActiveTab('calendar')}
              style={{ cursor: 'pointer' }}
            >
              Booking Calendar
            </li>
            <li
              className={`list-group-item ${activeTab === 'myBookings' ? 'active' : ''}`}
              onClick={() => setActiveTab('myBookings')}
              style={{ cursor: 'pointer' }}
            >
              My Bookings
            </li>
          </ul>
        </div>

        {/* Main Content */}
        <div className="col-12 col-md-9 p-3">
          {activeTab === 'calendar' ? (
            <FullCalendar
              plugins={[dayGridPlugin, interactionPlugin]}
              initialView="dayGridMonth"
              dateClick={(info) => onDateClick(info.dateStr)}
            />
          ) : (
            <MyBookings />
          )}
        </div>
      </div>
    </div>
  );
};

export default BookingLayout;


