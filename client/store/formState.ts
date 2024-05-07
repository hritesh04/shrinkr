import { atom } from "recoil";
type formStateType = "SHORT" | "QR";
export const formState = atom<formStateType>({
  key: "formState",
  default: "SHORT",
});
