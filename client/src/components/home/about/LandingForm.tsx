import FormIndicator from "./FormIndicator";
export default function LandingForm({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="w-1/2">
      <div className="w-full">
        <div className="flex gap-2">
          <FormIndicator />
        </div>
        {children}
      </div>
    </div>
  );
}
