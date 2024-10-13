import { IoIosLink } from "react-icons/io";
import Accordian from "../../../components/dashboard/analytics/accordian";
import { urls } from "../../../utils/dummyUrl";

export default function Analytics() {
  return (
    <div className="h-full w-full flex">
      <div className="h-full w-2/3 text-center overflow-auto scrollbar-hide pb-4">
        <h1 className=" text-2xl mb-6">Performance</h1>
        <div className="flex w-full flex-col gap-2">
          {urls.map((url, i) => {
            return <Accordian key={i} url={url} />;
          })}
        </div>
      </div>
      <div className=" p-4 w-1/3 flex flex-col mt-10 gap-7">
        <div className=" flex flex-col gap-3 p-4 bg-[#0c1562] rounded-md">
          <div className=" flex flex-col gap-1">
            <h1 className=" flex items-center gap-2 text-2xl">
              <div>
                <span>CREATE NEW</span>
                <br />
                <p className=" flex gap-2 items-center">
                  LINK <IoIosLink />
                </p>
              </div>
            </h1>
            <p className=" text-xs">Create, short and manage your links</p>
          </div>
          <input
            type="text"
            placeholder="https://super-long-link.com/shorten-it"
            className="w-full p-1 border-1 border-[#b6b6b6] rounded-md"
          />
          <button className=" w-full bg-[#2336f9] p-2 rounded-md">
            Create Link
          </button>
        </div>
        <div className=" flex flex-col bg-[#0c1562] rounded-md  p-4">
          <div className="w-full">
            <h1 className=" text-xl mb-5">
              LINK SHORTNER & <br /> QR GENERATOR <br />
              FOR ANY OF YOUR NEED
            </h1>
            <p className="text-xs">
              Create short links, QR Codes, share them anywhere. Track what's
              working, and what's not.
            </p>
            {/* <button className=" w-full mt-4 p-4 rounded-md bg-[#2336f9]">
              Get Started
            </button> */}
          </div>
        </div>
      </div>
    </div>
  );
}

//bg-[#101105]
