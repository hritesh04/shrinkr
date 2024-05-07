"use client";
import { getModalState, modalToggle } from "@/store/modal";
import { useRecoilValue } from "recoil";

export default function Button() {
  const state = useRecoilValue(getModalState);
  return <button>State is :{state ? "Open" : "false"}</button>;
}
