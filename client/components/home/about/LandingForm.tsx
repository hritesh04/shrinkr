"use client";

import { useState } from "react";
import ShortForm from "./ShortForm";
import { IoIosLink } from "react-icons/io";
import { IoQrCodeOutline } from "react-icons/io5";
import { cn } from "@/utils/clsx";
import QrForm from "./QrForm";

type FormType = "SHORT" | "QR";

export default function LandingForm() {
  const [form, setForm] = useState<FormType>("SHORT");
  return (
    <div className="w-full">
      <div className="flex gap-2">
        <button
          className={cn(
            "flex items-center gap-1 p-2",
            form === "SHORT" && "bg-white text-black rounded-t-md"
          )}
          onClick={() => setForm("SHORT")}
        >
          <IoIosLink />
          Short Link
        </button>
        <button
          className={cn(
            "flex items-center gap-1 p-2",
            form === "QR" && "bg-white text-black rounded-t-md"
          )}
          onClick={() => setForm("QR")}
        >
          <IoQrCodeOutline />
          QR Code
        </button>
      </div>
      {form === "SHORT" && <ShortForm />}
      {form === "QR" && <QrForm />}
    </div>
  );
}
