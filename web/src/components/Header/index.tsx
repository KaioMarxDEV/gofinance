export function Header() {
  return (
    <div className="pt-10 pb-28 bg-gray-900 shadow-green-500 shadow-xl">
      <div className="w-full max-w-6xl px-6 mx-auto flex flex-row justify-between items-center">
        <div>
          <h1 className="text-center text-2xl font-light">
            GoFinance
            <span className="text-green-500">
              $
            </span>
          </h1>
        </div>
        <div>
          <button className="rounded-md bg-gray-900 border border-green-500 border-b-4 border-b-green-500 hover:translate-y-1 hover:bg-green-500 transition-all ease-in-out delay-75 duration-300 h-full w-full px-5 py-2">
            <strong className="text-base">
              New Transaction +
            </strong>
          </button>
        </div>
      </div>
    </div>
  )
}
