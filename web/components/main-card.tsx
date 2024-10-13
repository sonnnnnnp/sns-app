import { Card } from "./ui/card";

type Props = {
  children: React.ReactNode;
};

export function MainCard({ children }: Props) {
  return (
    <Card className="flex text-sm w-full rounded-none border-0 mb-16 bg-muted sm:mb-0 sm:border-r-[1px] md:border-x-[1px] dark:bg-black dark:border-slate-800">
      {children}
    </Card>
  );
}
