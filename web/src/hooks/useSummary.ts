import { Transaction } from "../contexts/TransactionsContext"

export function useSummary(transactions: Transaction[]) {
  let income = 0, outcome = 0, total = 0

  transactions.map(transaction => {
    if (transaction.type === 'income') {
      income += transaction.number
      total += transaction.number
    } else {
      outcome += transaction.number
      total -= transaction.number
    }
  })

  return {income, outcome, total}
}
