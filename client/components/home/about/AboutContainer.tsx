export default function AboutContainer({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="md:flex-row flex flex-col w-full justify-around px-16 pt-40 pb-30">
      {children}
    </div>
  );
}
