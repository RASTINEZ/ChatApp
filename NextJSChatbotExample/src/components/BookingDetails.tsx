'use client';

import { Capacitor } from '@capacitor/core';
import { sendNotification, isPermissionGranted, requestPermission } from '@tauri-apps/plugin-notification';
import { useEffect, useState } from 'react';
import Image from 'next/image';
import { fetchRooms, fetchTimeSlots, submitBooking } from '../api/booking';

interface BookingDetailsProps {
    selectedDate: string;
}

type NotificationData = {
    title: string;
    body: string;
};

type ForegroundEvent = {
    getNotification: () => NotificationData;
    complete: (notification: NotificationData) => void;
};

type OneSignalCordovaPlugin = {
    setAppId: (appId: string) => void;
    promptForPushNotificationsWithUserResponse: (cb: (accepted: boolean) => void) => void;
    setNotificationWillShowInForegroundHandler: (cb: (event: ForegroundEvent) => void) => void;
    setNotificationOpenedHandler: (cb: (event: unknown) => void) => void;
};

interface Room {
    id: number;
    name: string;
    image_url: string;
}

interface TimeSlot {
    id: number;
    time_label: string;
}

const BookingDetails: React.FC<BookingDetailsProps> = ({ selectedDate }) => {
    const [rooms, setRooms] = useState<Room[]>([]);
    const [timeSlots, setTimeSlots] = useState<TimeSlot[]>([]);
    const [selectedRoom, setSelectedRoom] = useState<number | null>(null);
    const [selectedTime, setSelectedTime] = useState<number | null>(null);
    const ENABLE_ONESIGNAL = false;

    useEffect(() => {
        const fetchData = async () => {
            const roomsRes = await fetchRooms();        // ← already the data, no .json() needed
            const timeslotsRes = await fetchTimeSlots();

            setRooms(roomsRes);         // ✅ this is fine now
            setTimeSlots(timeslotsRes); // ✅
        };
        fetchData();
    }, []);


    useEffect(() => {
        const init = async () => {
            try {
                const platform = Capacitor.getPlatform();

                if (platform === 'web') {
                    const granted = await isPermissionGranted();
                    if (!granted) {
                        const permission = await requestPermission();
                        console.log('Tauri notification permission:', permission);
                    }
                } else {
                    if (!ENABLE_ONESIGNAL) {
                        console.log('OneSignal temporarily disabled for testing.');
                        return;
                    }

                    const OneSignal = (window as unknown as { plugins: { OneSignal: OneSignalCordovaPlugin } })?.plugins?.OneSignal;
                    if (!OneSignal) {
                        console.warn('OneSignal not available');
                        return;
                    }

                    OneSignal.setAppId('YOUR_ONESIGNAL_APP_ID');

                    OneSignal.promptForPushNotificationsWithUserResponse((accepted) => {
                        console.log('Push permission accepted:', accepted);
                    });

                    OneSignal.setNotificationWillShowInForegroundHandler(async (event) => {
                        const notification = event.getNotification();
                        alert(`📅 ${notification.title}\n${notification.body}`);
                        const audio = new Audio('/assets/sounds/siren-alert-96052.mp3');
                        audio.play();
                        event.complete(notification);
                    });

                    OneSignal.setNotificationOpenedHandler((event) => {
                        console.log('Notification opened:', event);
                    });
                }
            } catch (error) {
                console.error('Initialization error:', error);
            }
        };

        init();
    }, [ENABLE_ONESIGNAL]);

    useEffect(() => {
        if (!selectedRoom || !selectedTime) return;

        const notify = async () => {
            const roomName = rooms.find((r) => r.id === selectedRoom)?.name || 'a room';
            const timeLabel = timeSlots.find((t) => t.id === selectedTime)?.time_label || 'a time';
            const platform = Capacitor.getPlatform();

            const title = '📅 Booking Confirmed';
            const body = `You've selected ${roomName} at ${timeLabel}`;

            if (platform === 'web') {
                await sendNotification({ title, body });
                new Audio('/assets/sounds/siren-alert-96052.mp3').play();
                
            } else {
                alert(`${title}\n${body}`);
                new Audio('/assets/sounds/siren-alert-96052.mp3').play();

                console.log('📡 Submitting booking...');
                
            }
            await submitBooking({
                room_id: selectedRoom,
                timeslot_id: selectedTime,
                date: selectedDate,
            });
        };

        notify();
    }, [rooms, selectedDate, selectedRoom, selectedTime, timeSlots]);


    return (
        <div>
            <h2 className="text-xl font-semibold mb-2">Booking Details for {selectedDate}</h2>

            <div className="mb-4">
                <p className="font-medium mb-2">Select a Room:</p>
                <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
                    {rooms.map((room) => (
                        <div
                            key={room.id}
                            onClick={() => {
                                setSelectedRoom(room.id);
                                setSelectedTime(null);
                            }}
                            className={`cursor-pointer border rounded-lg overflow-hidden shadow-md hover:shadow-lg transition ${selectedRoom === room.id ? 'ring-2 ring-blue-500' : ''}`}
                        >
                            <Image
                                src={room.image_url}
                                alt={room.name}
                                width={500}
                                height={128}
                                className="w-full h-32 object-cover"
                            />
                            <p className="text-center py-2 font-medium">{room.name}</p>
                        </div>
                    ))}
                </div>
            </div>

            {selectedRoom && (
                <div className="mb-4">
                    <p className="font-medium mb-2">Select a Time:</p>
                    <div className="flex flex-wrap gap-2">
                        {timeSlots.map((slot) => (
                            <button
                                key={slot.id}
                                className={`px-4 py-2 border rounded ${selectedTime === slot.id ? 'bg-green-500 text-white' : 'bg-gray-100'}`}
                                onClick={() => setSelectedTime(slot.id)}
                            >
                                {slot.time_label}
                            </button>
                        ))}
                    </div>
                </div>
            )}

            {selectedRoom && selectedTime && (
                <div className="mt-4 p-4 bg-green-100 rounded">
                    <p>
                        ✅ You selected <strong>{rooms.find((r) => r.id === selectedRoom)?.name}</strong> at <strong>{timeSlots.find((t) => t.id === selectedTime)?.time_label}</strong>
                    </p>
                </div>
            )}
        </div>
    );
};

export default BookingDetails;
