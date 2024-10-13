import LandingForm from "./LandingForm";
import ShortForm from "./ShortForm";
import { useRecoilState, useRecoilValue } from "recoil";
import AboutDetails from "./AboutDetails";
import AboutContainer from "./AboutContainer";
import QrForm from "./QrForm";import { formState } from "../../../store/formState";
import { modalToggle } from "../../../store/modal";
import { cn } from "../../../utils/clsx";
import LandingModal from "../../modals/LandingModal";
export default function About() {
  const [isOpen, setIsOpen] = useRecoilState(modalToggle);
  const form = useRecoilValue(formState);
  return (
    <>
      <div
        className={cn(
          "w-full min-h-[110vh] flex flex-col items-center",
          isOpen && "blur-sm"
        )}
      >
        <AboutContainer>
          <AboutDetails>
            <button
              className=" w-[57%] mt-4 p-4 rounded-md bg-[#2336f9]"
              onClick={() => setIsOpen(true)}
            >
              Get Started
            </button>
          </AboutDetails>
          <LandingForm>
            {form === "SHORT" && <ShortForm modalToggle={setIsOpen} />}
            {form === "QR" && <QrForm modalToggle={setIsOpen} />}
          </LandingForm>
        </AboutContainer>
      </div>
      {isOpen && <LandingModal />}
    </>
  );
}
