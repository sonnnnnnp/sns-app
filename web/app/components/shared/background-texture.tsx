export function BackgroundTexture() {
  return (
    <div className="fixed inset-0 pointer-events-none z-[99]">
      <div className="w-full h-full bg-[url('/misc/background-texture.png')] bg-[length:109px] bg-repeat opacity-[0.06]" />
    </div>
  );
}
