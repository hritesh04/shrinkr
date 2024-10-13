import { VscGraphLine } from "react-icons/vsc";
import { IoIosLink } from "react-icons/io";
import { IoQrCodeOutline } from "react-icons/io5";
import { RxDashboard } from "react-icons/rx";
import { IoSettingsOutline } from "react-icons/io5";

export const sidebarItems = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: RxDashboard,
  },
  {
    name: "Short Link",
    href: "/shrink",
    icon: IoIosLink,
  },
  {
    name: "Generate QR",
    href: "/qr",
    icon: IoQrCodeOutline,
  },
  {
    name: "Analytics",
    href: "/analytics",
    icon: VscGraphLine,
  },
  {
    name: "Settings",
    href: "/settings",
    icon: IoSettingsOutline,
  },
];
