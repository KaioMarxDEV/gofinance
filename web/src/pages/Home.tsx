import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Banner } from "../components/Banner";
import { Header } from "../components/Header";
import { SearchBar } from "../components/SearchBar";
import { Transactions } from "../components/Transactions";

export function Home() {
  const navFunction = useNavigate()

  useEffect(() => {
    const token = localStorage.getItem("@gofinanceTokenString")

    if (token === null) {
      navFunction("/")
      // TODO: toastify here that user was redirected
    }
  }, [])

  return (
    <main>
      <Header />
      <Banner />
      <div className="w-full max-w-6xl mx-auto px-5 flex flex-row mb-4">
        <h2 className="font-bold text-4xl">
          Your Transaction<span className="text-green-500">$</span> :
        </h2>
      </div>
      <SearchBar />
      <Transactions />
    </main>
  )
}
