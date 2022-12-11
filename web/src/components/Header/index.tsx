import { Dialog, Transition } from "@headlessui/react";
import * as RadioGroup from '@radix-ui/react-radio-group';
import axios from "axios";
import jwt_decode from 'jwt-decode';
import { ArrowDown, ArrowUp, CurrencyDollarSimple, User, XCircle } from "phosphor-react";
import { FormEvent, Fragment, useEffect, useState } from "react";

interface ResponseDTO {
  success: boolean;
  message: string;
  data: {
    description: string;
    number: number;
    category: string;
    type: 'income' | 'outcome';
  };
}

interface Token {
  exp: number;
  user_id: string;
  username: string;
}

export function Header() {
  const [userID, setUserID] = useState("")
  const [username, setUsername] = useState("")

  const [description, setDescription] = useState("")
  const [number, setNumber] = useState(0)
  const [category, setCategory] = useState("")
  let [isOpen, setIsOpen] = useState(false)

  useEffect(() => {
    async function loadUserStoragedData() {
      const token = localStorage.getItem("@gofinanceTokenString") as string
      const { username, user_id } = jwt_decode(token) as Token
      setUsername(username)
      setUserID(user_id)
    }

    loadUserStoragedData()
  }, [])

  function closeModal() {
    setIsOpen(false)
  }

  function openModal() {
    setIsOpen(true)
  }

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    const token = localStorage.getItem("@gofinanceTokenString")
    const authHeader = `Authentication: Bearer ${token}`

    // TODO: create query add transaction to database passing userID
    const response = await axios.post(
      "http://localhost:3000/api/v1/transaction/add",
      {
        description,
        category,
        number,
        // TODO: add type from radio item
        userID
      },
      {
        headers: {
          authHeader
        }
      }
    )

    const {success, message, data} = response.data as ResponseDTO

    if (success == true) {
      // TODO: add context here to reflect new data on transactions component
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
                        className="inline-flex justify-center rounded-md border border-transparent bg-red-100 px-2 py-1 text-red-500 hover:bg-red-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
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
                  <form onSubmit={handleSubmit} className="mt-4 flex flex-col justify-center">
                    <input
                      className="inline-flex text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Description"
                      type="text"
                      onChange={(e) => setDescription(e.target.value)}
                      required
                    />
                    <input
                      className="inline-flex mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Number"
                      type="number"
                      onChange={(e) => setNumber(e.target.valueAsNumber)}
                      required
                    />
                    <input
                      className="inline-flex mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Category"
                      type="text"
                      // TODO: start here
                      onChange={(e) => setCategory(e.target.value)}
                      required
                    />
                    <RadioGroup.Root className="mt-4 gap-4 inline-flex">
                      <RadioGroup.Item value="income" className="aria-checked:text-white gap-1 flex-1 flex items-center flex-row justify-center p-4 text-gray-900 aria-checked:bg-green-500 hover:bg-gray-300 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2">
                        Income
                        <ArrowUp size={16} />
                        <CurrencyDollarSimple size={16} />
                      </RadioGroup.Item>
                      <RadioGroup.Item value="outcome" className="aria-checked:text-white gap-1 flex-1 flex items-center flex-row justify-center p-4 text-gray-900 aria-checked:bg-red-500 hover:bg-gray-300 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2">
                        Outcome
                        <ArrowDown size={16} />
                        <CurrencyDollarSimple size={16} />
                      </RadioGroup.Item>
                    </RadioGroup.Root>
                    <div className="mt-4 flex justify-center">
                      <button
                        type="submit"
                        className="inline-flex justify-center rounded-md border border-transparent bg-green-100 px-8 py-5 text-sm font-medium text-gray-900 hover:text-green-900 hover:bg-green-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      >
                        Submit Transaction
                      </button>
                    </div>
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
