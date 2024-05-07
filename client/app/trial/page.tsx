import Button from "./components/button";
import ClientRender from "./components/clientrender";
import ServerRender from "./components/serverrender";

export default function Trial() {
  return (
    <ClientRender>
      <ServerRender>
        <Button />
      </ServerRender>
    </ClientRender>
  );
}
