// app/layout.tsx
'use client';

import './globals.css'; // Make sure to import global styles here
import 'bootstrap/dist/css/bootstrap.min.css';
// import { useFirebasePushInit } from './_firebasePushInit';

// import Sidebar from '../components/SidebarForDashboard';
import Header from '../components/Header';

export default function RootLayout({
  
  children,
  
}: {
  children: React.ReactNode;
  
}) {
  // useFirebasePushInit();
  return (
    <html lang="en">
      <body>
        <div className="d-flex">
          {/* Sidebar */}
          {/* <Sidebar /> */}

          <div className="flex-grow-1">
            {/* Header */}
            <Header />

            {/* Main content */}
            <div className="container mt-4">{children}</div>
          </div>
        </div>
      </body>
    </html>
  );
}
