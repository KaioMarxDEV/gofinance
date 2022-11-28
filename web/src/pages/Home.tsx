import { Banner } from "../components/Banner";
import { Header } from "../components/Header";
import { Transactions } from "../components/Transactions";

export function Home() {
  return (
    <main className="h-screen">
      <Header />
      <Banner />
      <Transactions />
    </main>
  )
}
