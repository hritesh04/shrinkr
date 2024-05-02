"use client";

import QrPreview from "./QrPreview";

export default function QrForm() {
  return (
    <div className=" bg-white w-full text-black p-4 flex flex-col gap-4 rounded-md">
      <div className="flex flex-col gap-2">
        <label>Paste Your Url</label>
        <input
          type="text"
          className="p-1 border-1 border-[#b6b6b6] rounded-md text-medium"
          placeholder="https://super-long-link.com/shorten-it"
        />
        <button className=" w-full bg-[#2336f9] p-3 rounded-md text-white mt-2">
          Generate QR Code
        </button>
      </div>
      <div className=" w-full h-52">
        <QrPreview />
      </div>
    </div>
  );
}
