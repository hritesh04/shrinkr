import { Outlet } from "react-router-dom";
import Sidebar from "../../components/dashboard/sidebar";
import Navbar from "../../components/home/navbar/navbar";

export default function DashboardLayout() {
  return (
    <div className="">
      <Navbar />
      <Sidebar>
        <Outlet />
      </Sidebar>
    </div>
  );
}
