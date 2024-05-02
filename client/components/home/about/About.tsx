import LandingForm from "./LandingForm";

export default function About() {
  return (
    <div className="w-full min-h-[110vh] flex flex-col items-center pt-40 pb-30">
      <div className="md:flex-row flex flex-col w-full justify-around px-16">
        <div className="w-1/2">
          <p className=" bg-[#d8e3ff] text-[#0f2afa] w-fit px-4 p-1 rounded-2xl text-sm mb-4">
            Create link with just one click
          </p>
          <h1 className=" text-5xl mb-5">
            LINK SHORTNER & <br />
            QR GENERATOR <br />
            FOR ANY OF YOUR NEED
          </h1>
          <p>
            Create short links, QR Codes, share them anywhere.
            <br />
            Track what's working, and what's not.
          </p>
          <button className=" w-[57%] mt-4 p-4 rounded-md bg-[#2336f9]">
            Get Started
          </button>
        </div>
        <div className="w-1/2">
          <LandingForm />
        </div>
      </div>
    </div>
  );
}
