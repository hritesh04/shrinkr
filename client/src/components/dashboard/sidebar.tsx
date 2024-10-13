import { useLocation } from "react-router-dom";
import { cn } from "../../utils/clsx";
import { sidebarItems } from "../../utils/sidebarData";

export default function Sidebar({ children }: { children: React.ReactNode }) {
  const {pathname} = useLocation();
  return (
    <div className="h-[87.5vh] w-full grid grid-cols-12">
      <div className="h-full w-full col-span-2 p-4 flex flex-col gap-4">
        {sidebarItems.map((items) => {
          return (
            <a href={items.href}>
              <div
                className={cn(
                  "flex items-center justify-start p-8 pl-12 rounded-lg gap-2",
                  pathname === items.href && "scale-100"
                )}
              >
                {<items.icon size={20} />}
                <h1>{items.name}</h1>
              </div>
            </a>
          );
        })}
      </div>
      <div className="col-span-10 h-full overflow-auto">{children}</div>
    </div>
  );
}
