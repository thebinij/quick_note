import "./App.css";
import {
  createBrowserRouter,
  createRoutesFromElements,
  redirect,
  Route,
  RouterProvider,
} from "react-router-dom";
import Home, { loader as HomeLoader } from "./pages/Home";
import Login, { loader as loginLoader, action as loginAction } from "./pages/Login";

import Error from "./components/Error";
import Layout from "./components/Layout";
import Profile from "./pages/Profile";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<Layout />} errorElement={<Error />}>
      <Route index element={<Home />} loader={HomeLoader} />
      <Route path="login" element={<Login />} loader={loginLoader}  action={loginAction}/>
      <Route path="profile" element={<Profile/>} loader={async()=>{
        const isLoggedin = localStorage.getItem('USER_CRED')
        if(!isLoggedin){
          throw redirect("/login?message=You must login first")
        }
        return null
      }} />
    </Route>
  )
);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
