import { IoMdLock } from "react-icons/io";

export default function ShortForm() {
  return (
    <div className=" bg-white w-full text-black p-4 flex flex-col gap-4 rounded-b-md rounded-r-md">
      <label className="">Paste the long url</label>
      <input
        type="text"
        placeholder="https://super-long-link.com/shorten-it"
        className="w-full p-1 text-lg border-1 border-[#b6b6b6] rounded-md"
      />
      <div className=" flex flex-col">
        <label>Domain</label>
        <div className="flex gap-2 mt-2 items-start justify-around">
          <input
            type="text"
            placeholder={"shrinkr.com"}
            className="p-1 border-1 border-[#b6b6b6] rounded-md w-[45%] text-lg"
            disabled
          />
          <IoMdLock
            style={{ position: "relative", right: 40, top: 8 }}
            size={20}
          />
          <span className="relative right-4 text-xl top-1">/</span>
          <div className=" relative bottom-8">
            <label>Enter back-half</label>
            <input
              type="text"
              placeholder="shrinkr.com"
              className="p-1 mt-2 border-1 border-[#b6b6b6] rounded-md w-full text-lg"
            />
          </div>
        </div>
      </div>
      <button className=" w-full bg-[#2336f9] p-3 rounded-md text-white">
        Shrink It
      </button>
    </div>
  );
}
