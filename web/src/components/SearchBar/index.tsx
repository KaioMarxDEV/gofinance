import { zodResolver } from '@hookform/resolvers/zod';
import { CircleNotch, MagnifyingGlass } from "phosphor-react";
import { useContext, useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import * as z from 'zod';
import { Transaction, TransactionContext } from '../../contexts/TransactionsContext';
import { api } from '../../lib/api';

interface ResponseDTO {
  success: boolean;
  message: string;
  data:    Transaction[];
}

const searchFormSchema = z.object({
  query: z.string(),
})

type SearchFormInputs = z.infer<typeof searchFormSchema>

export function SearchBar() {
  const {transactions, updateSearchedTransactions} = useContext(TransactionContext)
  const [disableStatus, setDisableStatus] = useState(false)

  useEffect(() => {
    transactions.length > 0 ? setDisableStatus(false) : setDisableStatus(true)
  }, [transactions])

  const { register, handleSubmit, formState: { isSubmitting } } = useForm<SearchFormInputs>({
    resolver: zodResolver(searchFormSchema)
  })

  async function handleSearchFormSubmit({ query }: SearchFormInputs) {
    try {
      // TODO: Create axios folder lib to abstract baseURL instance
      const token = localStorage.getItem("@gofinanceTokenString")
      const response = await api.get(`/transaction/search?q=${query}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      const {success, message, data} = response.data as ResponseDTO

      if (success === true) {
        updateSearchedTransactions(data)
      } else {
        throw new Error(message)
      }
    } catch (error) {
      // toastify
      console.log(error)
    }
  }

  return (
    <form
      className="w-full max-w-6xl mx-auto px-5 flex flex-row"
      onSubmit={handleSubmit(handleSearchFormSubmit)}
    >
      <input
        className="disabled:cursor-not-allowed flex-1 transition-all delay-100 duration-300 mr-4 bg-gray-800 p-4 rounded-md focus:shadow-lg focus:shadow-green-400 outline-none ring-0"
        placeholder="Search by transaction name or category name if any..."
        {...register('query')}
      />
      <div>
        <button
          className="flex flex-row justify-between items-center w-full rounded-md px-8 py-4 bg-gray-900 disabled:opacity-40 hover:bg-gray-900/50 border border-gray-800 hover:border-green-500 transition-all ease-in delay-75 duration-200"
          type="submit"
          disabled={isSubmitting}
        >
          {
            isSubmitting
            ? <CircleNotch className="text-green-500 animate-spin" size={16} />
            : <MagnifyingGlass className="text-green-500" size={16} />
          }
          <span className="ml-2 text-base text-green-500">
            Search
          </span>
        </button>
      </div>
    </form>
  )
}
