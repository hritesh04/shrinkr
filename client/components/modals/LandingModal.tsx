"use client";
import Image from "next/image";
import { useState } from "react";
import { IoClose } from "react-icons/io5";
import Icon from "../../public/Icons/mainIcon.png";
import { useSetRecoilState } from "recoil";
import { modalToggle } from "@/store/modal";
type SignOnTypes = "REGISTER" | "SIGNIN";

export default function LandingModal() {
  const setModalState = useSetRecoilState(modalToggle);
  const [signon, setSignOn] = useState<SignOnTypes>("REGISTER");
  return (
    <div className="m-auto fixed top-10 rounded-md text-center bg-black p-8 border left-[33rem] z-10 h-fit">
      <div className=" flex h-28 mb-8 justify-between w-full">
        <Image src={Icon} alt="Shrinkr" className=" object-contain" />
        <button className=" absolute right-5">
          <IoClose size={25} onClick={() => setModalState(false)} />
        </button>
      </div>
      <div className=" mb-6">
        <h1 className=" text-2xl">Create Your Account</h1>
        <p className=" text-sm">
          Unlock exclusive perk
          <br />
          SignUP for special offers!
        </p>
      </div>
      <div className=" text-start flex flex-col gap-2">
        <div>
          {signon === "REGISTER" && (
            <>
              <label>Name</label>
              <input
                className="text-black w-full rounded-md p-2"
                placeholder="Name"
                type="text"
              />
            </>
          )}
        </div>
        <div>
          <label>Email</label>
          <input
            className="text-black w-full rounded-md p-2"
            placeholder="Email"
            type="text"
          />
        </div>
        <div>
          <label>Password</label>
          <input
            className="text-black w-full rounded-md p-2"
            placeholder="Password"
            type="password"
          />
        </div>
        <button className="w-full mt-2 p-2 rounded-md bg-[#2336f9]">
          {signon === "REGISTER" ? "Register" : "SIGNIN"}
        </button>
        <p className=" mt-2">
          {signon === "REGISTER"
            ? "Already have an account ?"
            : "New here ? let's get you started "}

          <span
            className="pl-2 underline text-[#4655f0]"
            onClick={() =>
              setSignOn((prev) => {
                return prev === "REGISTER" ? "SIGNIN" : "REGISTER";
              })
            }
          >
            {signon === "REGISTER" ? "SignIN" : "SignUP"}
          </span>
        </p>
      </div>
    </div>
  );
}
