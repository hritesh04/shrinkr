import Logo from "/icons/mainIcon.png";
import UserStatus from "./Status";
import { useLocation } from "react-router-dom";
import { navItems } from "../../../utils/navbarData";
export default function Navbar() {
  const {pathname} = useLocation();
  console.log(pathname);
  return (
    <div className="flex w-full h-full p-2">
      <div>
        <img src={Logo} alt="Shrinkr" height={80} />
      </div>
      <div className="w-full items-center justify-between gap-8 flex">
        <div className="gap-10 flex m-auto">
          {navItems.map((item, index) => {
            return (
              <a href={item.href} key={index}>
                <h1 className="text-xl">{item.name}</h1>
              </a>
            );
          })}
        </div>
        <div>
          <UserStatus />
        </div>
      </div>
    </div>
  );
}
