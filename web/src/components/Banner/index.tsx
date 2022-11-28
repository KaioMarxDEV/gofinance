import { ArrowDown, ArrowUp, Bank, CurrencyDollar } from "phosphor-react";

export function Banner() {
  return (
    <div className="w-full">
      <div className="relative -top-20 max-w-6xl px-6 mx-auto flex flex-row items-center justify-between">
        <div className="bg-gray-800 flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:-translate-y-1 border-2 border-green-500 h-36 p-6 rounded-lg w-full">
          <header className="flex flex-row justify-between items-center">
            <span className="text-bold text-base">Deposits</span>
            <div className="flex flex-row">
              <CurrencyDollar className="text-green-500" size={28} />
              <ArrowUp className="text-green-500" size={28} />
            </div>
          </header>
          <div>
            <span className="text-3xl">R$ 15.000,00</span>
          </div>
        </div>
        <div className="bg-gray-800 flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:-translate-y-1 border-2 border-red-500 h-36 p-6 rounded-lg mx-8 w-full">
          <header className="flex flex-row justify-between items-center">
            <span className="text-bold text-base">Withdraws</span>
            <div className="flex flex-row">
              <CurrencyDollar className="text-red-500" size={28} />
              <ArrowDown className="text-red-500" size={28} />
            </div>
          </header>
          <div>
            <span className="text-3xl">-R$ 1.000,00</span>
          </div>
        </div>
        <div className="bg-gradient-to-b from-[#543ab7] to-[#00acc1] flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:-translate-y-1  border-4 h-36 p-6 rounded-lg w-full">
          <div className="flex flex-row justify-between items-center">
            <span>Total</span>
            <Bank size={28} />
          </div>
          <div>
            <span className="text-3xl">R$ 14.000,00</span>
          </div>
        </div>
      </div>
    </div>
  )
}
