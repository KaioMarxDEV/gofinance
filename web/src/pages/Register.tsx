import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { api } from "../lib/api";

interface ResponseDTO {
  success: boolean;
  data: string;
}

export function Register() {
  const navFunction = useNavigate()
  const [email, setEmail] = useState("")
  const [username, setUsername] = useState("")
  const [password, setPass] = useState("")


  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    try {
      e.preventDefault()
      const response = await api.post("/user/add", {
        username,
        password,
        email
      })

      const {success, data} = response.data as ResponseDTO

      if (success === true) {
        navFunction("/", {
          state: data
        })
      } else {
        throw new Error("failed to register new user")
      }
    } catch (error) {
      console.log(error)
    }
  }

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
            <form onSubmit={handleSubmit}>
              {/* <!-- Username input --> */}
              <div className="mb-6">
                <input
                  type="text"
                  onChange={(e) => setUsername(e.target.value)}
                  className="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-green-500 focus:outline-none"
                  placeholder="Username"
                  required
                />
              </div>

              {/* <!-- Email input --> */}
              <div className="mb-6">
                <input
                  type="text"
                  onChange={(e) => setEmail(e.target.value)}
                  className="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-green-500 focus:outline-none"
                  placeholder="Email address"
                  required
                />
              </div>

              {/* <!-- Password input --> */}
              <div className="mb-6">
                <input
                  type="password"
                  onChange={(e) => setPass(e.target.value)}
                  className="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-green-500 focus:outline-none"
                  placeholder="Password"
                  required
                />
              </div>
              {/* <!-- Submit button --> */}
              <button
                type="submit"
                className="inline-block px-7 py-3 bg-green-600 text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:bg-green-700 hover:shadow-lg focus:bg-green-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg transition duration-150 ease-in-out w-full"
                data-mdb-ripple="true"
                data-mdb-ripple-color="light"
              >
                Sign up
              </button>

            </form>
            <div className="flex items-center py-4">
                <div className="flex-grow h-px bg-white"></div>

                <span className="flex-shrink text-lg text-gray-100 px-6 font-light">OR</span>

                <div className="flex-grow h-px bg-white"></div>
            </div>
            <Link to="/">
              <button
                className="inline-block px-7 py-3 bg-blue-600 text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg transition duration-150 ease-in-out w-full"
                data-mdb-ripple="true"
                data-mdb-ripple-color="light"
              >
                Login existing account
              </button>
            </Link>
          </div>
        </div>
      </div>

      {/* Bottom waves */}
      <div>
        <svg className="relative w-full -mb-2 min-h-[100px] max-h-36" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
        viewBox="0 24 150 28" preserveAspectRatio="none" shapeRendering="auto">
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
