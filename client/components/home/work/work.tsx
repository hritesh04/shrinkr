import {
  TbCircleNumber1,
  TbCircleNumber2,
  TbCircleNumber3,
  TbCircleNumber4,
  TbCircleNumber5,
} from "react-icons/tb";

export default function HowWeWork() {
  return (
    <div className="px-20 h-60 flex flex-col gap-16">
      <div className="flex justify-between">
        <h1 className=" text-2xl">HOW WE WORK</h1>
        <p>
          All product you need to build brand connections, manage links and
          <br />
          QR Codes, and connect with audiences everywhere, in a single unified
          platform.
        </p>
      </div>
      <div className="bg-[#444d55] flex justify-between rounded-md p-4 px-8">
        <p className=" flex items-center gap-2">
          <TbCircleNumber1 style={{ color: "white" }} size={20} /> Put Link
        </p>
        <p className=" flex items-center gap-2">
          <TbCircleNumber2 style={{ color: "white" }} size={20} /> Click Shrink
          it
        </p>
        <p className=" flex items-center gap-2">
          <TbCircleNumber3 style={{ color: "white" }} size={20} />
          Create Custom Url
        </p>
        <p className=" flex items-center gap-2">
          <TbCircleNumber4 style={{ color: "white" }} size={20} />
          Generate QR Code
        </p>
        <p className=" flex items-center gap-2">
          <TbCircleNumber5 style={{ color: "white" }} size={20} />
          Monitor Your Link
        </p>
      </div>
    </div>
  );
}
