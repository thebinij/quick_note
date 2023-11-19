import { NavLink } from "react-router-dom";

export default function Header() {
  const isLoggedIn = localStorage.getItem("USER_CRED");
  return (
    <nav style={{ display: "flex", justifyContent: "center", gap: "1rem" }}>
      <NavLink to=".">Home</NavLink>
      <NavLink to="profile">Profile</NavLink>

      {!isLoggedIn ? (
        <NavLink to="login">Login</NavLink>
      ) : (
        <button
          onClick={() => {
            localStorage.removeItem("USER_CRED");
          }}
        >
          Logout
        </button>
      )}
    </nav>
  );
}
