import Sidebar from "@/components/dashboard/sidebar";
import Navbar from "@/components/home/navbar/navbar";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="">
      <Navbar />
      <Sidebar>{children}</Sidebar>
    </div>
  );
}
