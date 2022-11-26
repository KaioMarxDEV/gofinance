export function Login() {
  return (
    <main className="relative flex flex-col h-screen text-center bg-gradient-to-r from-[#543ab7] to-[#00acc1]">
      <div className="w-full h-full m-0 p-0 flex justify-center items-center text-center">
        Big Div of FORM
      </div>

      <div>
        <svg className="relative w-full -mb-2 min-h-[100px] max-h-36" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
        viewBox="0 24 150 28" preserveAspectRatio="none" shape-rendering="auto">
        <defs>
          <path id="gentle-wave" d="M-160 44c30 0 58-18 88-18s 58 18 88 18 58-18 88-18 58 18 88 18 v44h-352z" />
        </defs>
          <g>
            <use className="animate-[moveForever_-2s_cubic-bezier(.55,.5,.45,.5)_7s_infinite]" xlinkHref="#gentle-wave" x="48" y="0" fill="rgba(255,255,255,0.7" />
            <use className="animate-[moveForever_-3s_cubic-bezier(.55,.5,.45,.5)_10s_infinite]" xlinkHref="#gentle-wave" x="48" y="3" fill="rgba(255,255,255,0.5)" />
            <use className="animate-[moveForever_-4s_cubic-bezier(.55,.5,.45,.5)_13s_infinite]" xlinkHref="#gentle-wave" x="48" y="5" fill="rgba(255,255,255,0.3)" />
            <use className="animate-[moveForever_-5s_cubic-bezier(.55,.5,.45,.5)_20s_infinite]" xlinkHref="#gentle-wave" x="48" y="7" fill="#fff" />
          </g>
        </svg>
      </div>
    </main>
  )
}
