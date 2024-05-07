import { atom, selector, useRecoilState } from "recoil";

export const modalToggle = atom({
  key: "modalState",
  default: false,
});

export const getModalState = selector({
  key: "modalToggle",
  get: ({ get }: { get: any }) => {
    const state = get(modalToggle);
    return state;
  },
  set: ({ set }: { set: any }) => {
    const [state, setState] = useRecoilState(modalToggle);
    set(!state);
  },
});
