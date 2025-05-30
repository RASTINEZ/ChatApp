export type NotificationData = {
    title: string;
    body: string;
  };
  
  export type ForegroundEvent = {
    getNotification: () => NotificationData;
    complete: (notification: NotificationData) => void;
  };
  
  export type OneSignalCordovaPlugin = {
    setAppId: (appId: string) => void;
    promptForPushNotificationsWithUserResponse: (cb: (accepted: boolean) => void) => void;
    setNotificationWillShowInForegroundHandler: (cb: (event: ForegroundEvent) => void) => void;
    setNotificationOpenedHandler: (cb: (event: unknown) => void) => void;
  };
  