import { useState } from 'react'
import { Login } from './pages/Login'

function App() {
  const [count, setCount] = useState(0)

  return (
    <main className="App">
      <Login />
    </main>
  )
}

export default App
