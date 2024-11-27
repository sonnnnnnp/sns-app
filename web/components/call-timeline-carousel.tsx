import * as DialogPrimitive from "@radix-ui/react-dialog";
import Link from "next/link";
import React from "react";
import { Avatar, AvatarImage } from "./ui/avatar";
import { Button } from "./ui/button";
import { Card, CardContent } from "./ui/card";
import { Carousel, CarouselContent, CarouselItem } from "./ui/carousel";
import { Dialog, DialogOverlay, DialogPortal } from "./ui/dialog";

export function CallTimelineCarousel() {
  const [isDialogOpen, setDialogOpen] = React.useState(true);

  const handleJoinCall = async () => {
    // documentPictureInPicture.requestWindow();
  };

  return (
    <Dialog
      open={isDialogOpen}
      onOpenChange={() => setDialogOpen(!isDialogOpen)}
    >
      <DialogPortal>
        <DialogOverlay />
        <DialogPrimitive.Content className="fixed left-[50%] top-[50%] z-50 w-full translate-x-[-50%] translate-y-[-50%] duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%]">
          <Carousel
            opts={{ containScroll: false }}
            className="flex w-full items-center"
          >
            <CarouselContent>
              {Array.from({ length: 5 }).map((_, index) => (
                <CarouselItem key={index} className="basis-1/1 lg:basis-1/3">
                  <Card>
                    <CardContent className="flex flex-col items-center p-4 gap-2">
                      <span className="text-xs">
                        英語やります。あんまり話しません。
                      </span>
                      <div className="flex flex-wrap gap-2">
                        <Link href={`/`}>
                          <Avatar>
                            <AvatarImage src="/users/placeholder-profile.svg" />
                          </Avatar>
                        </Link>
                        <Link href={`/`}>
                          <Avatar>
                            <AvatarImage src="/users/placeholder-profile.svg" />
                          </Avatar>
                        </Link>
                        <Link href={`/`}>
                          <Avatar>
                            <AvatarImage src="/users/placeholder-profile.svg" />
                          </Avatar>
                        </Link>
                        <Link href={`/`}>
                          <Avatar>
                            <AvatarImage src="/users/placeholder-profile.svg" />
                          </Avatar>
                        </Link>
                        <Link href={`/`}>
                          <Avatar>
                            <AvatarImage src="/users/placeholder-profile.svg" />
                          </Avatar>
                        </Link>
                      </div>
                      <span className="text-xs">ずばりの通話</span>
                      <Button className="w-full" onClick={handleJoinCall}>
                        通話に参加
                      </Button>
                    </CardContent>
                  </Card>
                </CarouselItem>
              ))}
            </CarouselContent>
          </Carousel>
        </DialogPrimitive.Content>
      </DialogPortal>
    </Dialog>
  );
}
