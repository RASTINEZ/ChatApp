'use client';

import React, { useState, useEffect } from 'react';
import Image from 'next/image';
import UserStatusPanel from './UserStatusPanel';

const MainMenu: React.FC = () => {
    const images = [
        '/assets/banner1.jpg',
        '/assets/banner2.jpg',
        '/assets/banner3.jpg'
    ];

    const announcements = [
        "🚨 System maintenance on April 25, 10:00 PM - 12:00 AM.",
        "🎉 New feature: Booking reminders now available!",
    ];

    const tasks = [
        { id: 1, task: 'Review pending bookings', completed: false },
        { id: 2, task: 'Update room images', completed: true },
        { id: 3, task: 'Check today\'s schedule', completed: false },
    ];

    const upcomingBookings = [
        { room: 'Room A101', date: '2025-04-23', time: '10:00 AM' },
        { room: 'Room B202', date: '2025-04-23', time: '01:00 PM' },
    ];

    const [currentImage, setCurrentImage] = useState(0);

    // Auto-slide every 3 seconds
    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentImage((prev) => (prev + 1) % images.length);
        }, 3000);
        return () => clearInterval(interval);
    }, [images.length]);

    return (
        <div className="flex flex-col lg:flex-row gap-6">
            <div className="space-y-6">
                {/* Image Slideshow */}
                <div className="relative w-full max-w-4xl h-80 mx-auto rounded overflow-hidden shadow">
                    <Image
                        src={images[currentImage]}
                        alt="Banner"
                        layout="fill"
                        objectFit="cover"
                    />
                </div>


                {/* Dots Navigation */}
                <div className="flex justify-center mt-2">
                    {images.map((_, idx) => (
                        <button
                            key={idx}
                            onClick={() => setCurrentImage(idx)}
                            className={`w-3 h-3 rounded-full mx-1 transition ${currentImage === idx ? 'bg-blue-600' : 'bg-gray-400'
                                }`}
                        />
                    ))}
                </div>


                {/* Announcements */}
                <section>
                    <h2 className="text-xl font-semibold mb-3">📢 Announcements</h2>
                    <ul className="list-disc list-inside bg-yellow-100 p-4 rounded">
                        {announcements.map((note, idx) => (
                            <li key={idx}>{note}</li>
                        ))}
                    </ul>
                </section>

                {/* Grid Section */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                    {/* Upcoming Bookings */}
                    <section>
                        <h2 className="text-xl font-semibold mb-3">📅 Upcoming Bookings</h2>
                        <ul className="bg-blue-100 p-4 rounded max-h-64 overflow-y-auto">
                            {upcomingBookings.map((b, idx) => (
                                <li key={idx} className="mb-2">
                                    <strong>{b.room}</strong> — {b.date} at {b.time}
                                </li>
                            ))}
                        </ul>
                    </section>
                    {/* To-Do List */}
                    <section>
                        <h2 className="text-xl font-semibold mb-3">✅ To-Do List</h2>
                        <ul className="bg-gray-100 p-4 rounded max-h-64 overflow-y-auto">
                            {tasks.map(task => (
                                <li key={task.id} className="flex items-center mb-2">
                                    <input type="checkbox" checked={task.completed} readOnly className="mr-2" />
                                    <span className={task.completed ? 'line-through text-gray-500' : ''}>
                                        {task.task}
                                    </span>
                                </li>
                            ))}
                        </ul>
                    </section>


                </div>
            </div>


            {/* Right Side: User Status Panel */}
            <div className="w-full lg:w-1/4 bg-gray-50 p-4 rounded shadow h-fit">
                <UserStatusPanel />
            </div>
        </div >


    );
};

export default MainMenu;
