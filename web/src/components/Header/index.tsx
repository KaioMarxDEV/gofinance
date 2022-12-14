import { Dialog, Transition } from "@headlessui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import * as RadioGroup from '@radix-ui/react-radio-group';
import jwt_decode from 'jwt-decode';
import { ArrowDown, ArrowUp, CurrencyDollarSimple, SpinnerGap, User, XCircle } from "phosphor-react";
import { Fragment, useContext, useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import * as z from 'zod';
import { TransactionContext } from "../../contexts/TransactionsContext";
import { api } from "../../lib/api";

interface ResponseDTO {
  success: boolean;
  message: string;
  data: {
    ID: string;
    category: string;
    description: string;
    number: number;
    type: 'income' | 'outcome'
    createdAt: string
  };
}

interface Token {
  exp: number;
  user_id: string;
  username: string;
}

const newModalTransactionSchema = z.object({
  description: z.string(),
  number: z.number(),
  category: z.string(),
  type: z.enum(['income', 'outcome'])
})

type newModalTransactionsInputs = z.infer<typeof newModalTransactionSchema>

export function Header() {
  const { update } = useContext(TransactionContext)

  const [username, setUsername] = useState("")
  let [isOpen, setIsOpen] = useState(false)


  useEffect(() => {
    async function loadUserStoragedData() {
      const token = localStorage.getItem("@gofinanceTokenString") as string
      const { username } = jwt_decode(token) as Token
      setUsername(username)
    }

    loadUserStoragedData()
  }, [])

  function closeModal() {
    setIsOpen(false)
  }

  function openModal() {
    setIsOpen(true)
  }

  const {
    register,
    handleSubmit,
    reset,
    formState: {
      isSubmitting,
    },
    control
  } = useForm<newModalTransactionsInputs>({
    resolver: zodResolver(newModalTransactionSchema)
  })

  async function handleNewTransactionModal(inputData: newModalTransactionsInputs) {
    const token = localStorage.getItem("@gofinanceTokenString") as string

    const response = await api.post(
      "transaction/add",
      inputData,
      {
        headers: {
          Authorization: `Bearer ${token}`
        }
      }
    )

    const {success, message, data} = response.data as ResponseDTO
    if (success == true) {
      update(data)
      reset()
      closeModal()
    } else {
      throw new Error(message)
    }
  }

  return (
    <>
      {/* Header Content */}
      <div className="pt-10 pb-28 bg-gray-900 shadow-green-500 shadow-xl">
        <div className="w-full max-w-6xl px-6 mx-auto flex flex-row justify-between items-center">
          {/* Title */}
          <div>
            <h1 className="text-center text-2xl font-light">
              GoFinance
              <span className="text-green-500">
                $
              </span>
            </h1>
          </div>
          <div className="flex items-center">
            <div className="mr-14">
              <div className="flex flex-col items-center">
                <User size={28}/>
                <span className="text-lg font-semibold">
                  {username}
                </span>
              </div>
            </div>
            <button
              type="button"
              onClick={openModal}
              className="inline-flex justify-center rounded-md  px-8 py-5 text-base bg-green-600/80 border-2 border-green-500 text-white hover:bg-green-500 hover:border-green-500 transition-all ease-in delay-75 duration-200"
              >
              Add Transaction
            </button>
          </div>
        </div>
      </div>
      {/* Modal */}
      <Transition appear show={isOpen} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={closeModal}>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-black bg-opacity-25" />
          </Transition.Child>

          <div className="fixed inset-0 overflow-y-auto">
            <div className="flex min-h-full items-center justify-center p-4 text-center">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 scale-95"
                enterTo="opacity-100 scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 scale-100"
                leaveTo="opacity-0 scale-95"
              >
                <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                  {/* Modal Header and Close button */}
                  <div>
                    <div className="flex-1 flex flex-row justify-end">
                      <button
                        className="inline-flex justify-center rounded-md border border-transparent bg-red-100 px-2 py-1 text-red-500 hover:bg-red-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2 transition-all delay-75 duration-300"
                        onClick={closeModal}
                      >
                        <XCircle size={32}/>
                      </button>
                    </div>
                    <Dialog.Title
                      as="h3"
                      className="mt-1 text-lg text-center underline underline-offset-2 decoration-green-500 font-medium leading-6 text-gray-900"
                      >
                      New Transaction
                    </Dialog.Title>
                    <div className="mt-2">
                      <p className="text-sm text-center text-gray-500">
                        Your new transaction will be registered in the table
                        on dashboard if submit succeeds.
                      </p>
                    </div>
                  </div>

                  {/* New transaction form */}
                  <form onSubmit={handleSubmit(handleNewTransactionModal)} className="mt-4 flex flex-col justify-center">
                    <input
                      className="flex w-full text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Description"
                      type="text"
                      required
                      {...register('description')}
                    />
                    <input
                      className="flex w-full mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Number"
                      type="number"
                      required
                      {...register('number', { valueAsNumber: true })}
                    />
                    <input
                      className="flex w-full mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Category"
                      type="text"
                      required
                      {...register('category')}
                    />


                    <Controller
                      control={control}
                      name="type"
                      render={({field}) => {
                         return (
                          <RadioGroup.Root onValueChange={field.onChange} value={field.value} className="mt-4 gap-4 inline-flex">
                            <RadioGroup.Item value="income" className="gap-1 flex-1 flex items-center flex-row justify-center p-4 aria-checked:bg-green-500 hover:bg-gray-300 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2 transition-all delay-75 duration-300">
                              <span className="text-gray-900">
                                Income
                              </span>
                              <ArrowUp className="text-gray-900" size={16} />
                              <CurrencyDollarSimple className="text-gray-900" size={16} />
                            </RadioGroup.Item>
                            <RadioGroup.Item value="outcome" className="gap-1 flex-1 flex items-center flex-row justify-center p-4 aria-checked:bg-red-500 hover:bg-gray-300 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2 transition-all delay-75 duration-300">
                              <span className="text-gray-900">
                                Outcome
                              </span>
                              <ArrowDown className="text-gray-900" size={16} />
                              <CurrencyDollarSimple className="text-gray-900" size={16} />
                            </RadioGroup.Item>
                          </RadioGroup.Root>
                        )
                      }}
                    />

                    <button
                      type="submit"
                      className="disabled:opacity-40 mt-4 inline-flex justify-center rounded-md border border-transparent bg-blue-600 mx-2 py-5 hover:scale-105 text-sm font-medium text-gray-100 hover:bg-blue-500 focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2 transition-all ease-linear delay-75 duration-300"
                      disabled={isSubmitting}
                    >
                      {
                        isSubmitting
                        ? <SpinnerGap className="animate-spin" size={16} />
                        : "Submit Transaction"
                      }
                    </button>
                  </form>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition>
    </>
  )
}
