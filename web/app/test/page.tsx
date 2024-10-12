export function Page() {
  return (
    <div className="flex justify-center min-h-dvh">
      <aside>
        <div className="sticky inset-y-0 w-64 h-dvh bg-emerald-700 hidden lg:block">
          navigation
        </div>
        <div className="sticky inset-y-0 w-16 h-dvh bg-emerald-700 hidden sm:block lg:hidden">
          nav
        </div>
      </aside>
      <main className="flex-grow max-w-[750px] bg-gray-500 sm:border-x-[1px]">
        <div className="w-full bg-sky-600">
          <div className="sticky top-0 z-10 h-14 border-b-[1px] bg-sky-300">
            tabs
          </div>
          <div className="h-32 w-full border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
          <div className="h-32 border-b-[1px]">primary contents</div>
        </div>
        <aside className="sticky bottom-0 h-14 border-t-[1px] bg-emerald-700 sm:hidden">
          navigation
        </aside>
      </main>
      <div className="sticky inset-y-0 w-80 h-dvh bg-sky-900 hidden xl:block">
        widget contents
      </div>
    </div>
  );
}

export default Page;
