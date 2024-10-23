export function Page() {
  return (
    <div className="flex min-h-dvh">
      <nav>
        <div className="sticky inset-y-0 w-64 h-dvh bg-emerald-700 hidden xl:block">
          navigation
        </div>
        <div className="sticky inset-y-0 w-16 h-dvh bg-emerald-700 hidden sm:block xl:hidden">
          nav
        </div>
        <div className="fixed bottom-0 w-full h-14 border-t-[1px] bg-emerald-700 sm:hidden">
          navigation
        </div>
      </nav>
      <main className="flex-grow bg-gray-500 sm:border-x-[1px]">
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
      </main>
      <div className="sticky inset-y-0 w-80 h-dvh bg-sky-900 hidden lg:block">
        widget contents
      </div>
    </div>
  );
}

export default Page;
