// firebaseConfig.ts

import { initializeApp, getApps, getApp } from 'firebase/app';
import { getMessaging, onMessage, getToken, Messaging } from 'firebase/messaging';

// Firebase configuration object (replace with your own credentials from Firebase console)
const firebaseConfig = {
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID,
  measurementId: process.env.NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID,
};



// Only initialize app once
const app = !getApps().length ? initializeApp(firebaseConfig) : getApp();

let messaging: Messaging | null = null;

// ✅ Only initialize messaging in the browser
if (typeof window !== 'undefined' && typeof navigator !== 'undefined') {
  messaging = getMessaging(app);
}

export const requestNotificationPermission = async () => {
  if (!messaging) return;

  try {
    const permission = await Notification.requestPermission();
    if (permission === 'granted') {
      const token = await getToken(messaging, {
        vapidKey: 'YOUR_WEB_PUSH_VAPID_KEY',
      });
      console.log('Web Push Notification Token:', token);
      // TODO: send token to server if needed
    } else {
      console.warn('Notification permission denied');
    }
  } catch (error) {
    console.error('Error requesting notification permission:', error);
  }
};

export const setupPushNotificationHandlers = () => {
  if (!messaging) return;

  onMessage(messaging, (payload) => {
    console.log('Message received. ', payload);

    const { title, body, icon } = payload.notification || {};

    if (typeof title === 'string' && typeof body === 'string') {
      if (Notification.permission === 'granted') {
        new Notification(title, {
          body,
          icon,
        });
      }
    } else {
      console.warn('Invalid notification payload format');
    }
  });
};

export default messaging;