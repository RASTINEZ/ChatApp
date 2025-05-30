// components/Header.tsx
const Header = () => {
    return (
      <div className="bg-white shadow-sm p-3 d-flex justify-content-between align-items-center">
        <h1 className="h4">Dashboard</h1>
        <div className="d-flex align-items-center">
          <span className="mr-3">Welcome, TINTIN</span>
          <button className="btn btn-outline-primary">Logout</button>
        </div>
      </div>
    );
  };
  
  export default Header;
  