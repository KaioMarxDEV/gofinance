import { Header } from "../components/Header";

export function Home() {
  return (
    <main className="h-screen">
      <Header />
      {/* TODO: START HERE */}
      <div className="relative -top-20 max-w-4xl mx-auto flex flex-row items-center justify-between">
        <div className="bg-gray-800 h-36 p-6 rounded-lg w-full">
          Entradas
        </div>
        <div className="bg-gray-800 h-36 p-6 rounded-lg mx-8 w-full">
          Saidas
        </div>
        <div className="bg-gray-800 h-36 p-6 rounded-lg w-full">
          Total
        </div>
      </div>
    </main>
  )
}
