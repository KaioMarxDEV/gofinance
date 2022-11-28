import { Banner } from "../components/Banner";
import { Header } from "../components/Header";
import { SearchBar } from "../components/SearchBar";
import { Transactions } from "../components/Transactions";

export function Home() {
  return (
    <main>
      <Header />
      <Banner />
      <SearchBar />
      <Transactions />
    </main>
  )
}
