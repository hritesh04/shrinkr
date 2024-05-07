import { formState } from "@/store/formState";
import { cn } from "@/utils/clsx";
import { IoIosLink } from "react-icons/io";
import { IoQrCodeOutline } from "react-icons/io5";
import { useRecoilState } from "recoil";

export default function FormIndicator() {
  const [form, setForm] = useRecoilState(formState);
  return (
    <>
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
    </>
  );
}
