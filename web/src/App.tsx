import { BrowserRouter } from 'react-router-dom'
import { Router } from './Router'

function App() {
  return (
    <main className="App">
      <BrowserRouter>
        <Router />
      </BrowserRouter>
    </main>
  )
}

export default App
