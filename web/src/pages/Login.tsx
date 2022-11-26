export function Login() {
  return (
    <main className="relative flex flex-col h-screen text-center bg-gradient-to-r from-[#543ab7] to-[#00acc1]">
      {/* Login Content */}
      <div className="w-full h-full m-0 p-0 flex justify-center items-center text-center">
        <div className="container flex flex-col justify-center items-center">
          <div className="md:w-8/12 lg:w-5/12 lg:ml-20">
            <h1 className="text-center font-thin text-5xl mb-9">
              GoFinance
              <span className="text-green-500">
                $
              </span>
            </h1>
            <form>
              {/* <!-- Email input --> */}
              <div className="mb-6">
                <input
                  type="text"
                  className="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-green-500 focus:outline-none"
                  placeholder="Email address"
                />
              </div>

              {/* <!-- Password input --> */}
              <div className="mb-6">
                <input
                  type="password"
                  className="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-green-500 focus:outline-none"
                  placeholder="Password"
                />
              </div>
              {/* <!-- Submit button --> */}
              <button
                type="submit"
                className="inline-block px-7 py-3 bg-blue-600 text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out w-full"
                data-mdb-ripple="true"
                data-mdb-ripple-color="light"
              >
                Sign in
              </button>

            </form>
          </div>
        </div>
      </div>

      {/* Bottom waves */}
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
