import { ArrowDown, ArrowUp, Bank, CurrencyDollar } from "phosphor-react";
import { useContext, useEffect, useState } from "react";
import { Transaction, TransactionContext } from "../../contexts/TransactionsContext";
import { priceFormatter } from "../../utils/formatter";

export function Banner() {
  const { transactions } = useContext(TransactionContext)
  const [income, setIncome] = useState(0)
  const [outcome, setOutcome] = useState(0)
  const [total, setTotal] = useState(0)

  function calculateSummary(transactions: Transaction[]) {
    let income=0, outcome=0, total=0
    transactions.map(transaction => {
      if (transaction.type === 'income') {
        income += transaction.number
        total += transaction.number
      } else {
        outcome += transaction.number
        total -= transaction.number
      }
    })

    return { income, outcome, total }
  }

  useEffect(() => {
    const summary = calculateSummary(transactions);
    setIncome(summary.income)
    setOutcome(summary.outcome)
    setTotal(summary.total)
  }, [transactions])

  return (
    <div className="w-full">
      <div className="relative -top-20 max-w-6xl px-6 mx-auto flex flex-row items-center justify-between">
        <div className="bg-gray-800 flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:scale-105 border-2 border-green-500 h-36 p-6 rounded-lg w-full">
          <header className="flex flex-row justify-between items-center">
            <span className="text-bold text-base">Deposits</span>
            <div className="flex flex-row">
              <CurrencyDollar className="text-green-500" size={28} />
              <ArrowUp className="text-green-500" size={28} />
            </div>
          </header>
          <div>
            <span className="text-3xl">{priceFormatter.format(income)}</span>
          </div>
        </div>
        <div className="bg-gray-800 flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:scale-105 border-2 border-red-500 h-36 p-6 rounded-lg mx-8 w-full">
          <header className="flex flex-row justify-between items-center">
            <span className="text-bold text-base">Withdraws</span>
            <div className="flex flex-row">
              <CurrencyDollar className="text-red-500" size={28} />
              <ArrowDown className="text-red-500" size={28} />
            </div>
          </header>
          <div>
            <span className="text-3xl">-{priceFormatter.format(outcome)}</span>
          </div>
        </div>
        <div className="bg-gradient-to-b from-[#543ab7] to-[#00acc1] flex flex-col justify-between transition-all ease-in delay-75 duration-200 hover:scale-105  h-36 p-6 rounded-lg w-full">
          <div className="flex flex-row justify-between items-center">
            <span>TOTAL</span>
            <Bank size={28} />
          </div>
          <div>
            <span className="text-3xl">{priceFormatter.format(total)}</span>
          </div>
        </div>
      </div>
    </div>
  )
}
