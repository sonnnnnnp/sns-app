import { Card, CardBody } from "@nextui-org/card";

export function Widget() {
  return (
    <Card
      shadow="none"
      className="w-full h-36"
      classNames={{ base: "border bg-background" }}
    >
      <CardBody>
        <p>Widget</p>
      </CardBody>
    </Card>
  );
}
