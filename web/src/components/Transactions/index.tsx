import { Binoculars } from "phosphor-react";
import { useContext } from "react";
import { TransactionContext } from "../../contexts/TransactionsContext";

export function Transactions() {
  const {transactions} = useContext(TransactionContext)

  return (
    <div className="w-full max-w-6xl mt-8 mb-8 mx-auto px-5">
      {transactions.length > 0 ? (
        <table className="w-full">
        <tbody className="flex flex-col gap-3">
          {transactions.map(transaction => (
            <tr className="bg-gray-800 flex justify-between rounded-xl py-5 px-8">
              <td className="w-1/2">{transaction.description}</td>
              <td className={transaction.type === 'income' ? 'text-green-300': 'text-rose-200'}>${transaction.number}</td>
              <td>{transaction.category}</td>
            </tr>
          ))}
        </tbody>
      </table>
      ) : (
        <div className="w-full">
          <div className="flex flex-col bg-gray-800 shadow-green-400 shadow-lg py-6 rounded-md items-center">
            <div className="flex flex-row">
              <Binoculars size={60} className="text-green-500 animate-pulse" />
            </div>
            <span className="flex text-xl font-semibold mb-3">
              We're <p className="mx-2 text-green-500 animate-pulse">looking</p> for...
            </span>
            <span className="text-lg font-semibold">But it seems you don't have any transactions</span>
          </div>
        </div>
      )}
    </div>
  )
}
