"use client";

import * as VisuallyHidden from "@radix-ui/react-visually-hidden";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import { CircleHelpIcon } from "lucide-react";
import { Button } from "../ui/button";

type Props = {
  open: boolean;
  message: string;
  onConfirm?: () => void;
  onCancel?: () => void;
};

export function ConfirmActionDialog({
  open,
  message,
  onConfirm,
  onCancel,
}: Props) {
  const handleConfirm = () => {
    onConfirm && onConfirm();
  };

  const handleCancel = () => {
    onCancel && onCancel();
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
          <CircleHelpIcon />
          <p>{message}</p>
          <div className="flex flex-wrap justify-center gap-2">
            <Button className="min-w-24 rounded-lg" onClick={handleConfirm}>
              OK
            </Button>
            <Button
              variant="secondary"
              className="min-w-24 rounded-lg"
              onClick={handleCancel}
            >
              キャンセル
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
