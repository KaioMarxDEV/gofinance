import axios from "axios";
import { createContext, ReactNode, useEffect, useState } from "react";

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
    const {data} = await axios.get("http://localhost:3000/api/v1/transaction/all")

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
