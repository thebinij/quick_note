import { Outlet } from "react-router-dom";
import Header from "./Header";
import Footer from "./Footer";

export default function Layout() {
  const mainStyle = {
    flex: 1,
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        minHeight: "90vh",
      }}
    >
      <Header />
      <main style={mainStyle}>
        <Outlet />
      </main>
      <Footer />
    </div>
  );
}
