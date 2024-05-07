export default function ServerRender({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div>
      Hello world
      <div>{children}</div>
    </div>
  );
}
