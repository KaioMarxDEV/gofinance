import { Banner } from "../components/Banner";
import { Header } from "../components/Header";
import { SearchBar } from "../components/SearchBar";
import { Transactions } from "../components/Transactions";

export function Home() {
  return (
    <main>
      <Header />
      <Banner />
      <div className="w-full max-w-6xl mx-auto px-5 flex flex-row mb-8">
        <h2 className="font-bold text-4xl">
          Your Transaction<span className="text-green-500">$</span> :
        </h2>
      </div>
      <SearchBar />
      <Transactions />
    </main>
  )
}
