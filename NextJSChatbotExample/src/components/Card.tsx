// components/Card.tsx
interface CardProps {
    title: string;
    value: string;
    description: string;
    icon: string;
  }
  
  const Card: React.FC<CardProps> = ({ title, value, description, icon }) => {
    return (
      <div className="card shadow-sm">
        <div className="card-body">
          <div className="d-flex align-items-center">
            <div className="display-4 text-primary">{icon}</div>
            <div className="ml-3">
              <h5 className="card-title">{title}</h5>
              <p className="card-text">{description}</p>
            </div>
          </div>
          <h2 className="mt-3 text-center">{value}</h2>
        </div>
      </div>
    );
  };
  
  export default Card;
  