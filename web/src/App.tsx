import { useState } from 'react'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <div className="flex justify-center align-top mt-8 bg-gray-400">
        <p className='text-xs text-slate-800'>{count}</p>
        <button className='p-3 rounded-xl' onClick={() => setCount(count+1)}>
          +1
        </button>
      </div>
    </div>
  )
}

export default App
