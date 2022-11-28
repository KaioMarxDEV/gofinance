import { MagnifyingGlass } from "phosphor-react";

export function SearchBar() {
  return (
    <div className="w-full max-w-6xl mx-auto px-5 flex flex-row">
      <input
        className="flex-1 mr-4 bg-gray-800 p-4 rounded-md focus:shadow-lg focus:shadow-green-400 outline-none ring-0"
        placeholder="Search by transaction name..."
      />
      <div>
        <button className="flex flex-row justify-between items-center w-full rounded-md px-8 py-4 bg-gray-900 hover:bg-gray-900/50 border border-gray-800 hover:border-green-500 transition-all ease-in delay-75 duration-200">
          <MagnifyingGlass className="text-green-500" size={16} />
          <span className="ml-2 text-base text-green-500">
            Search
          </span>
        </button>
      </div>
    </div>
  )
}
