// // app/_firebasePushInit.ts
// 'use client';
// // _firebasePushInit.ts

// import { useEffect } from 'react';
// import { Capacitor } from '@capacitor/core';
// import { requestNotificationPermission, setupPushNotificationHandlers } from './firebaseConfig';
// import { PushNotifications } from '@capacitor/push-notifications';

// export const useFirebasePushInit = () => {
//   useEffect(() => {
//     const initPushNotifications = async () => {
//       if (Capacitor.getPlatform() !== 'web') {
//         // Mobile Platform (Android or iOS) -> Initialize Capacitor PushNotifications Plugin
//         try {
//           await PushNotifications.requestPermissions();
//           await PushNotifications.register();

//           PushNotifications.addListener('registration', (token) => {
//             console.log('Push notification token:', token);
//             // Send this token to your server to manage notifications
//           });

//           PushNotifications.addListener('pushNotificationReceived', (notification) => {
//             console.log('Received push notification:', notification);
//             // Handle mobile notification here
//           });

//           PushNotifications.addListener('pushNotificationActionPerformed', (action) => {
//             console.log('Push notification action performed:', action);
//             // Handle notification action (click)
//           });
//         } catch (error) {
//           console.error('Error initializing PushNotifications:', error);
//         }
//       } else {
//         // Web Platform -> Initialize Firebase Web Push Notifications
//         requestNotificationPermission();
//         setupPushNotificationHandlers();
//       }
//     };

//     initPushNotifications();
//   }, []);
// };
