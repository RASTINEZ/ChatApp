'use client';

import React, { useState, useEffect, useRef } from 'react';

import Image from 'next/image';

interface GroupChatLayoutProps {
    onBack: () => void;
}

const GroupChatLayout: React.FC<GroupChatLayoutProps> = ({ onBack }) => {

    const [showProfile, setShowProfile] = useState<{
        name: string;
        x: number;
        y: number;
    } | null>(null);

    const [previewImage, setPreviewImage] = useState<string | null>(null);

    const people = [
        { name: 'Alice', avatar: '/assets/images/alice.png', role: 'Developer' },
        { name: 'Bob', avatar: '/assets/images/bob.png', role: 'Designer' },
        { name: 'Charlie', avatar: '/assets/images/charlie.png', role: 'Project Manager' },
    ];

    const media = ['/assets/banner1.jpg', '/assets/banner2.jpg'];
    const documents = ['Project_Specs.pdf', 'Design_Doc.docx'];

    const handlePersonClick = (personName: string, e: React.MouseEvent) => {
        setShowProfile({ name: personName, x: e.clientX, y: e.clientY });
    };
    const popoverRef = useRef<HTMLDivElement>(null);



    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (popoverRef.current && !popoverRef.current.contains(event.target as Node)) {
                setShowProfile(null);
            }
        };

        if (showProfile) {
            document.addEventListener('mousedown', handleClickOutside);
        }

        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, [showProfile]);



    return (
        <div className="flex h-screen relative">
            {/* Main Chat Area */}
            <div className="flex-1 flex flex-col p-4 overflow-hidden">
                <button
                    className="mb-4 px-4 py-2 bg-gray-300 hover:bg-gray-400 text-black rounded w-fit"
                    onClick={onBack}
                >
                    ⬅ Back to Chats
                </button>

                <h2 className="text-xl font-semibold mb-4">🖥️ IT Group Chat</h2>
                <div className="flex-1 overflow-y-auto border p-4 rounded bg-white">
                    <p><strong>Alice:</strong> Let&apos;s finalize the design today.</p>
                    <p><strong>Bob:</strong> Sure! Uploading assets now.</p>
                </div>
                <div className="mt-3 flex">
                    <input type="text" placeholder="Type your message..." className="flex-1 border rounded px-3 py-2 mr-2" />
                    <button className="bg-blue-500 text-white px-4 py-2 rounded">Send</button>
                </div>
            </div>

            {/* Sidebar */}
            <div className="w-72 border-l p-4 bg-gray-50 overflow-y-auto">
                {/* People */}
                <section className="mb-6">
                    <h4 className="font-semibold mb-2">👥 People</h4>
                    <ul>
                        {people.map((person, idx) => (
                            <li
                                key={idx}
                                className="flex items-center mb-3 cursor-pointer hover:bg-gray-200 p-2 rounded"
                                onClick={(e) => handlePersonClick(person.name, e)}
                            >
                                <Image src={person.avatar} alt={person.name} width={32} height={32} className="rounded-full mr-2" />
                                <span>{person.name}</span>
                            </li>
                        ))}
                    </ul>
                </section>

                {/* Media */}
                <section className="mb-6">
                    <h4 className="font-semibold mb-2">🖼️ Media</h4>
                    <div className="flex flex-wrap gap-2">
                        {media.map((src, idx) => (
                            <img
                                key={idx}
                                src={src}
                                alt={`media-${idx}`}
                                className="w-20 h-20 object-cover rounded cursor-pointer"
                                onClick={() => setPreviewImage(src)}
                            />
                        ))}
                    </div>
                </section>

                {/* Documents */}
                <section>
                    <h4 className="font-semibold mb-2">📄 Documents</h4>
                    <ul className="list-disc list-inside">
                        {documents.map((doc, idx) => (
                            <li key={idx} className="text-blue-600 cursor-pointer hover:underline">{doc}</li>
                        ))}
                    </ul>
                </section>
            </div>

            {/* Profile Popover */}
            {showProfile && (
                <div
                    className="fixed bg-white border rounded shadow-lg p-4"
                    ref={popoverRef}
                    style={{
                        top: Math.min(showProfile.y, window.innerHeight - 150),
                        left: Math.min(showProfile.x, window.innerWidth - 220),
                        width: '200px',
                        zIndex: 1000
                    }}
                >
                    {(() => {
                        const person = people.find(p => p.name === showProfile.name);
                        return person ? (
                            <>
                                <Image src={person.avatar} alt={person.name} width={64} height={64} className="rounded-full mx-auto mb-2" />
                                <h4 className="text-center font-semibold">{person.name}</h4>
                                <p className="text-center text-sm text-gray-500">{person.role}</p>
                            </>
                        ) : null;
                    })()}
                </div>
            )}




            {/* Image Preview Lightbox */}
            {previewImage && (
                <div className="fixed inset-0 bg-black bg-opacity-70 flex justify-center items-center">
                    <img src={previewImage} alt="Preview" className="max-w-3xl max-h-[80vh] rounded shadow" />
                    <button className="absolute top-5 right-5 text-white text-2xl" onClick={() => setPreviewImage(null)}>✖</button>
                </div>
            )}
        </div>
    );
};

export default GroupChatLayout;
