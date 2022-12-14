import { createContext, ReactNode, useEffect, useState } from "react";
import { api } from "../lib/api";

interface ResponseDTO {
  success: boolean;
  message: string;
  data: Transaction[];
}

export interface Transaction {
  ID: string;
  description: string;
  number: number;
  category: string;
  type: 'income' | 'outcome';
  createdAt: string;
}

interface TransactionContextType {
  transactions: Transaction[];
  update: (data: Transaction) => {};
  remove: (ID: string) => {};
  updateSearchedTransactions: (data: Transaction[]) => {}
}
interface TransactionProviderProps {
  children: ReactNode;
}

export const TransactionContext = createContext({} as TransactionContextType)


export function TransactionProvider({ children }: TransactionProviderProps) {
  const [transactions, setTransactions] = useState<Transaction[]>([])

  async function updateSearchedTransactions(data: Transaction[]) {
    setTransactions(data)
  }

  async function updateTransactions(data: Transaction) {
    setTransactions([...transactions, data])
  }
  async function removeTransactions(ID: string) {
    const removedArray = transactions.filter(transaction => transaction.ID !== ID)
    setTransactions(removedArray)
  }

  async function loadTransactions() {
    try {
      const token = localStorage.getItem("@gofinanceTokenString") as string
      if (!token) {
      // TODO: toastify here, error handling is missing
        console.log("There an error here should be toastified")
      }

      const response = await api.get(
        "/transaction/all",
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      )

      const {success, message, data} = response.data as ResponseDTO

      if (success === true) {
        setTransactions(data)
      } else {
        throw new Error(message)
      }
    } catch (error) {
      // TODO: toastify this error on loading transactions
      console.log(error)
    }
  }

  useEffect(() => {
    loadTransactions()
  }, [])

  return (
    <TransactionContext.Provider
      value={{
        transactions,
        update: updateTransactions,
        remove: removeTransactions,
        updateSearchedTransactions
      }}
    >
      {children}
    </TransactionContext.Provider>
  )
}
