export default function AboutDetails({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
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
      {children}
    </div>
  );
}
