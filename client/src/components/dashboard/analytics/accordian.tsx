"use client";
import { IoLinkOutline } from "react-icons/io5";
import { IoIosLink } from "react-icons/io";
import { useState } from "react";
import Analytics from "./analytics";

type urlType = {
  id: number;
  url: string; //
  shortenedUrl: string; //
  totalClicks: number; //
  expiry: string;
  rateRemaining: number; //
  isActive: boolean;
};

export default function Accordian({ url }: { url: urlType }) {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className="w-full bg-[#0c1562] flex flex-col gap-2 rounded-md p-4">
      <div className="flex flex-col gap-2">
        <div className="flex items-center px-2 justify-between">
          <div className=" flex items-center gap-2">
            <IoIosLink size={20} />
            <span className="w-72 truncate" title={url.url}>
              <a href={url.url}>{url.url}</a>

            </span>
          </div>
          <p
            className=" underline hover:cursor-pointer"
            onClick={() => setIsOpen(!isOpen)}
          >
            Show Analytics
          </p>
        </div>
        <div className="flex items-center gap-6 justify-between px-2">
          <div className=" flex items-center gap-2">
            <IoLinkOutline
              size={20}
              style={{ color: url.isActive ? "green" : "red" }}
            />
            <span>
              {"shrinkr.com"}/{url.shortenedUrl}
            </span>
          </div>
          <div className=" flex gap-8">
            <p>Visits : {url.totalClicks}</p>
            <p>ClicksRemaining : {url.rateRemaining}</p>
            <span>Expires on : {url.expiry}</span>
          </div>
        </div>
      </div>
      <div className="h-full w-full pt-4">
        {isOpen && <Analytics url={url.shortenedUrl} />}
      </div>
    </div>
  );
}
