import { Binoculars, Trash } from "phosphor-react";
import { useContext } from "react";
import { TransactionContext } from "../../contexts/TransactionsContext";
import { api } from "../../lib/api";
import { dateFormatter, priceFormatter } from "../../utils/formatter";

export function Transactions() {
  const {transactions, remove} = useContext(TransactionContext)

  async function handleRemoveTransaction(ID: string) {
    try {
      const token = localStorage.getItem("@gofinanceTokenString")
      const response = await api.delete(`/transaction/delete/${ID}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      const {success, message} = response.data

      if (success === true) {
        remove(ID)
      } else {
        throw new Error(message)
      }
    } catch (error) {

    }
  }

  return (
    <div className="w-full max-w-6xl mt-8 mb-8 mx-auto px-5">
      {transactions.length > 0 ? (
        <div className="w-full">
          <div className="flex flex-col gap-3">
            {transactions.map(transaction => (
              <div key={transaction.ID} className="bg-gray-800 flex items-center rounded-xl py-5 px-8">
                <div className="w-1/2 flex items-center gap-4">
                  <button
                    className="p-2 rounded-3xl bg-gray-700 hover:bg-gray-800 border-2 hover:text-red-500 hover:border-red-500 border-gray-700 transition-all delay-75 duration-200 ease-linear"
                    onClick={() => handleRemoveTransaction(transaction.ID)}
                  >
                    <Trash size={16}/>
                  </button>
                  {transaction.description}
                </div>
                <div className="w-1/2 flex justify-between">
                  <div className={transaction.type === 'income' ? 'text-green-300': 'text-rose-200'}>
                    {transaction.type === 'outcome' && "-"}
                    {priceFormatter.format(transaction.number)}
                  </div>
                  <div className="font-bold">{transaction.category}</div>
                  <div>{dateFormatter.format(new Date(transaction.createdAt))}</div>
                </div>
              </div>
            ))}
          </div>
       </div>
      ) : (
        <div className="w-full">
          <div className="flex flex-col bg-gray-800 shadow-green-400 shadow-lg py-6 rounded-md items-center">
            <div className="flex flex-row">
              <Binoculars size={60} className="text-green-500 animate-pulse" />
            </div>
            <span className="flex text-xl font-semibold mb-3">
              We're <p className="mx-2 text-green-500 animate-pulse">looking</p> for...
            </span>
            <span className="text-lg font-semibold">But it seems you have 0 transactions matching</span>
          </div>
        </div>
      )}
    </div>
  )
}
