import axios from "axios";
import { Binoculars } from "phosphor-react";
import { useContext } from "react";
import { TransactionContext } from "../../contexts/TransactionsContext";
import { dateFormatter, priceFormatter } from "../../utils/formatter";

export function Transactions() {
  const {transactions, remove} = useContext(TransactionContext)

  async function handleRemoveTransaction(ID: string) {
    try {
      const token = localStorage.getItem("@gofinanceTokenString")
      const response = await axios.delete(`http://localhost:3000/api/v1/transaction/delete/${ID}`, {
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
        <table className="w-full">
          <tbody className="flex flex-col gap-3">
            {transactions.map(transaction => (
              <tr key={transaction.ID} className="bg-gray-800 flex items-center rounded-xl py-5 px-8">

                <td className="w-1/2">{transaction.description}</td>
                <div className="w-1/2 flex justify-between ">
                  <td className={transaction.type === 'income' ? 'text-green-300': 'text-rose-200'}>
                    {transaction.type === 'outcome' && "-"}
                    {priceFormatter.format(transaction.number)}
                  </td>
                  <td>{transaction.category}</td>
                  <td>{dateFormatter.format(new Date(transaction.createdAt))}</td>
                </div>
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
