import { zodResolver } from '@hookform/resolvers/zod';
import { MagnifyingGlass } from "phosphor-react";
import { useContext, useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import * as z from 'zod';
import { TransactionContext } from '../../contexts/TransactionsContext';

const searchFormSchema = z.object({
  query: z.string(),
})

type SearchFormInputs = z.infer<typeof searchFormSchema>

export function SearchBar() {
  const {transactions} = useContext(TransactionContext)
  const [disableStatus, setDisableStatus] = useState(false)

  useEffect(() => {
    transactions.length > 0 ? setDisableStatus(false) : setDisableStatus(true)
  }, [transactions])

  const { register, handleSubmit } = useForm<SearchFormInputs>({
    resolver: zodResolver(searchFormSchema)
  })

  function handleSearchFormSubmit(data: SearchFormInputs) {
    console.log(data)
  }

  return (
    <form
      className="w-full max-w-6xl mx-auto px-5 flex flex-row"
      onSubmit={handleSubmit(handleSearchFormSubmit)}
    >
      <input
        className="disabled:cursor-not-allowed flex-1 transition-all delay-100 duration-300 mr-4 bg-gray-800 p-4 rounded-md focus:shadow-lg focus:shadow-green-400 outline-none ring-0"
        placeholder="Search by transaction name..."
        {...register('query')}
        disabled={disableStatus}
      />
      <div>
        <button
          className="flex flex-row justify-between items-center w-full rounded-md px-8 py-4 bg-gray-900 hover:bg-gray-900/50 border border-gray-800 hover:border-green-500 transition-all ease-in delay-75 duration-200"
          type="submit"
        >
          <MagnifyingGlass className="text-green-500" size={16} />
          <span className="ml-2 text-base text-green-500">
            Search
          </span>
        </button>
      </div>
    </form>
  )
}
