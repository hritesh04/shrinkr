import { IoLinkOutline } from "react-icons/io5";

export default function QrPreview() {
  return (
    <div className=" h-full w-full gap-2 flex">
      <img
        src={"https://www.w3schools.com/howto/img_5terre.jpg"}
        className="h-full w-1/3 object-cover"
      />
      <div className=" w-full h-full">
        <div className=" flex items-center mb-16 gap-2">
          <IoLinkOutline style={{ color: "#2336f9" }} size={20} />
          <span>hhfoeofoeahfoae</span>
        </div>
        <button className=" w-full bg-[#2336f9] p-3 rounded-md text-white mt-2">
          Download
        </button>
        <button className=" w-full bg-[#2336f9] p-3 rounded-md text-white mt-2">
          Share
        </button>
      </div>
    </div>
  );
}
