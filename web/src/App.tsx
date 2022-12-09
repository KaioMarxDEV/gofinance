import { BrowserRouter } from 'react-router-dom'
import { TransactionProvider } from './contexts/TransactionsContext'
import { Router } from './Router'

function App() {
  return (
    <main className="App">
      <BrowserRouter>
        <TransactionProvider>
          <Router />
        </TransactionProvider>
      </BrowserRouter>
    </main>
  )
}

export default App
