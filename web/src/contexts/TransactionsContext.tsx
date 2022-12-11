import axios from "axios";
import { createContext, ReactNode, useEffect, useState } from "react";

// TODO: add created at column

export interface Transaction {
  description: string;
  number: number;
  category: string;
  type: 'income' | 'outcome';
}

interface TransactionContextType {
  transactions: Transaction[]
}
interface TransactionProviderProps {
  children: ReactNode;
}

export const TransactionContext = createContext({} as TransactionContextType)


export function TransactionProvider({ children }: TransactionProviderProps) {
  const [transactions, setTransactions] = useState<Transaction[]>([])

  async function loadTransactions() {
    const token = localStorage.getItem("@gofinanceTokenString") as string
    if (!token) {
    // TODO: error handling is missing
      console.log("There an error here should be toastified")
    }

    const authHeader = `Authorization: Bearer ${token}`;

    const {data} = await axios.get(
      "http://localhost:3000/api/v1/transaction/all",
      {
        headers: {
          authHeader
        }
      }
    )

    setTransactions(data)
  }

  useEffect(() => {
    loadTransactions()
  }, [])

  return (
    <TransactionContext.Provider value={{transactions}}>
      {children}
    </TransactionContext.Provider>
  )
}
