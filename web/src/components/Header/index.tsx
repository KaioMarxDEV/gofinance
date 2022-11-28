export function Header() {
  return (
    <div className="h-52 bg-black shadow-green-700 shadow-xl">
      <div className="max-w-4xl mx-auto pt-10">
        <div className="flex flex-row justify-between items-center">
          <div>
            <h1 className="text-center text-2xl font-light">
              GoFinance
              <span className="text-green-500">
                $
              </span>
            </h1>
          </div>
          <div>
            <button className="border border-green-700 rounded-md hover:bg-green-500 transition-all ease-in-out delay-75 duration-500 h-full w-full px-5 py-2">
              <strong className="text-base">
                New Transaction
              </strong>
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
