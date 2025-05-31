import { useParams } from "@tanstack/react-router";

export const Hello = () => {
  const { helloId } = useParams({ from: "/$helloId" });
  return (
    <div>
      <h1>Hello World</h1>
      <p>Click on the links above to see the code splitting in action.</p>
      <p>
        The current helloId is: <strong>{helloId}</strong>
      </p>
    </div>
  );
};