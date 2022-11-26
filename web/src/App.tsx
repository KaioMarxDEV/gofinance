import { useState } from 'react'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <div>
        {count}
        <button onClick={() => setCount(count+1)}>
          +1
        </button>
      </div>
    </div>
  )
}

export default App
