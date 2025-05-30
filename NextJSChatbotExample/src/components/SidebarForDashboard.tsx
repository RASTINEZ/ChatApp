// components/Sidebar.tsx
import Link from 'next/link';

const Sidebar = () => {
  return (
    <div className="bg-dark text-white p-4" style={{ minHeight: "100vh", width: "250px" }}>
      <h2 className="mb-4">Meeting Room Booking</h2>
      <ul className="list-unstyled">
        <li>
          <Link href="/dashboard"
            className="text-white d-block py-2 px-3 rounded mb-2 bg-primary">Dashboard
          </Link>
        </li>
        <li>
          <Link href="/bookings"
            className="text-white d-block py-2 px-3 rounded mb-2">Bookings
          </Link>
        </li>
        <li>
          <Link href="/settings"
            className="text-white d-block py-2 px-3 rounded mb-2">Settings
          </Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
