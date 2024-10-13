import { createBrowserRouter } from "react-router-dom";
import Analytics from "../pages/dashboard/analytics/analytics";
import DashboardLayout from "../pages/dashboard/Dashboard";
import Home from "../pages/home/Home";

export const Routes = createBrowserRouter([
    {
      path:"/",
      element:<Home />,
      children: []
    },
    {
      path:"/dashboard",
      element:<DashboardLayout />,
      children:[
        // {
        //   path: "/",
        //   element:<></>
        // },
        {
          path:"analytics",
          element:<Analytics />
        }
      ]
    }
  ])