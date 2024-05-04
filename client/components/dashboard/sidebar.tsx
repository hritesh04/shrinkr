"use client";
import { cn } from "@/utils/clsx";
import { sidebarItems } from "@/utils/sidebarData";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Sidebar({ children }: { children: React.ReactNode }) {
  const path = usePathname();

  return (
    <div className="h-[87.5vh] w-full grid grid-cols-12">
      <div className="h-full w-full col-span-2 p-4 flex flex-col gap-4">
        {sidebarItems.map((items, i) => {
          return (
            <Link href={items.href}>
              <div
                className={cn(
                  "flex items-center justify-start p-8 pl-12 rounded-lg gap-2",
                  path === items.href && "scale-100"
                )}
              >
                {<items.icon size={20} />}
                <h1>{items.name}</h1>
              </div>
            </Link>
          );
        })}
      </div>
      <div className="col-span-10 h-full overflow-auto">{children}</div>
    </div>
  );
}
