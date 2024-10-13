"use client";

import * as VisuallyHidden from "@radix-ui/react-visually-hidden";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import { CircleXIcon } from "lucide-react";
import { Button } from "../ui/button";

type Props = {
  open: boolean;
  message: string;
  onConfirm?: () => void;
};

export function ActionErrorDialog({ open, message, onConfirm }: Props) {
  const handleConfirm = () => {
    onConfirm && onConfirm();
  };

  return (
    <Dialog open={open}>
      <DialogContent hideCloseButton={true} className="max-w-sm">
        <VisuallyHidden.Root>
          <DialogHeader>
            <DialogTitle />
            <DialogDescription />
          </DialogHeader>
        </VisuallyHidden.Root>
        <div className="flex flex-col space-y-4 items-center justify-center">
          <CircleXIcon className="text-red-600" />
          <p>{message}</p>
          <div className="flex flex-wrap justify-center gap-2">
            <Button className="min-w-24 rounded-lg" onClick={handleConfirm}>
              OK
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
