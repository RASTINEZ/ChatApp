// app/dashboard/page.tsx
import Card from '../../components/Card';

const Dashboard = () => {
  return (
    <div>
      <div className="row row-cols-1 row-cols-md-3 row-cols-lg-4 g-4">
        {/* Example Cards */}
        <div className="col">
          <Card
            title="Total Bookings"
            value="120"
            description="Total number of bookings"
            icon="📅"
          />
        </div>
        <div className="col">
          <Card
            title="Upcoming Bookings"
            value="5"
            description="Bookings scheduled for today"
            icon="🔔"
          />
        </div>
        <div className="col">
          <Card
            title="Pending Approvals"
            value="3"
            description="Bookings waiting for approval"
            icon="⏳"
          />
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
