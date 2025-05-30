'use client';

import React, { useState, useEffect, useRef } from 'react';
import Image from 'next/image';

const users = [
  { name: 'Alice', status: 'online', avatar: '/assets/images/alice.png', role: 'Developer' },
  { name: 'Bob', status: 'online', avatar: '/assets/images/bob.png', role: 'Designer' },
  { name: 'Charlie', status: 'online', avatar: '/assets/images/charlie.png', role: 'PM' },
  { name: 'David', status: 'offline', avatar: '/assets/images/charlie.png', role: 'Tester' },
  { name: 'Eve', status: 'offline', avatar: '/assets/images/charlie.png', role: 'Analyst' },
];

const UserStatusPanel: React.FC = () => {
  const [showProfile, setShowProfile] = useState<{
    name: string;
    avatar: string;
    role: string;
    x: number;
    y: number;
  } | null>(null);

  const popoverRef = useRef<HTMLDivElement>(null);

  // Detect click outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (showProfile && popoverRef.current && !popoverRef.current.contains(event.target as Node)) {
        setShowProfile(null);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [showProfile]);

  const handleUserClick = (
    user: { name: string; avatar: string; role: string },
    e: React.MouseEvent
  ) => {
    setShowProfile({
      ...user,
      x: e.clientX,
      y: e.clientY,
    });
  };

  return (
    <div className="h-full flex flex-col">
      <h5 className="font-semibold mb-2">🟢 Online</h5>
      <ul className="mb-4 overflow-y-auto flex-1">
        {users.filter(u => u.status === 'online').map(user => (
          <li
            key={user.name}
            className="flex items-center mb-3 cursor-pointer hover:bg-gray-200 p-2 rounded"
            onClick={(e) => handleUserClick(user, e)}
          >
            <Image src={user.avatar} alt={user.name} width={32} height={32} className="rounded-full mr-2" />
            <span>{user.name}</span>
          </li>
        ))}
      </ul>

      <h5 className="font-semibold mb-2 mt-4">⚫ Offline</h5>
      <ul className="overflow-y-auto flex-1 text-gray-500">
        {users.filter(u => u.status === 'offline').map(user => (
          <li
            key={user.name}
            className="flex items-center mb-3 cursor-pointer hover:bg-gray-100 p-2 rounded"
            onClick={(e) => handleUserClick(user, e)}
          >
            <Image src={user.avatar} alt={user.name} width={32} height={32} className="rounded-full mr-2" />
            <span>{user.name}</span>
          </li>
        ))}
      </ul>

      {/* Profile Popover */}
      {showProfile && (
        <div
          ref={popoverRef}
          className="fixed bg-white border rounded shadow-lg p-4"
          style={{
            top: Math.min(showProfile.y, window.innerHeight - 150),
            left: Math.min(showProfile.x, window.innerWidth - 220),
            width: '200px',
            zIndex: 1000,
          }}
        >
          <Image
            src={showProfile.avatar}
            alt={showProfile.name}
            width={64}
            height={64}
            className="rounded-full mx-auto mb-2"
          />
          <h4 className="text-center font-semibold">{showProfile.name}</h4>
          <p className="text-center text-sm text-gray-500">{showProfile.role}</p>
        </div>
      )}
    </div>
  );
};

export default UserStatusPanel;
