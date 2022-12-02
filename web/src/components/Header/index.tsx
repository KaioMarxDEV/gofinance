import { Dialog, Transition } from "@headlessui/react"
import { ChartPie, XCircle } from "phosphor-react"
import { Fragment, useState } from "react"

export function Header() {
  let [isOpen, setIsOpen] = useState(false)

  function closeModal() {
    setIsOpen(false)
  }

  function openModal() {
    setIsOpen(true)
  }

  function handleSubmit(e: { preventDefault: () => void }) {
    e.preventDefault()
    console.log("entrei")
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
            <nav className=" flex flex-row mr-8 gap-4 px-4 py-2 rounded-md bg-gray-800">
              <button
                type="button"
                className="flex flex-col items-center text-base text-green-500/80 hover:text-green-400 rounded-md cursor-pointer transition-all ease-in delay-75 duration-200 hover:translate-y-1"
                // TODO: create charts page and route to there!
                onClick={() => alert("go to page charts mf!")}
              >
                <ChartPie size={28} />
                charts
              </button>
            </nav>
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
                  <form className="mt-4 flex flex-col justify-center">
                    <input
                      className="inline-flex text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Description"
                      type="text"
                      required
                    />
                    <input
                      className="inline-flex mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Number"
                      type="number"
                      required
                    />
                    <input
                      className="inline-flex mt-4 text-gray-900 bg-gray-200 p-4 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                      placeholder="Category"
                      type="text"
                      required
                    />
                    <div className="mt-4 gap-4 inline-flex">
                      <button className="flex-1 p-4 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2">Income</button>
                      <button className="flex-1 p-4 bg-gray-200 rounded-md focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2">Outcome</button>
                    </div>
                    <div className="mt-4 flex justify-center">
                      <button
                        type="submit"
                        className="inline-flex justify-center rounded-md border border-transparent bg-green-100 px-8 py-5 text-sm font-medium text-gray-900 hover:text-green-900 hover:bg-green-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-green-500 focus-visible:ring-offset-2"
                        onClick={handleSubmit}
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
