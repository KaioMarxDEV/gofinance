export function Transactions() {
  let variant = 'income'
  return (
    <div className="w-full max-w-6xl mt-10 mx-auto px-5">
      <table className="w-full">
        <tbody className="flex flex-col gap-3">
          <tr className="bg-gray-800 flex justify-between rounded-xl py-5 px-8">
            <td className="w-1/2">Desenvolvimento de Software</td>
            <td className={variant === 'income' ? 'text-green-300': 'text-rose-200'}>-R$ 15.000,00</td>
            <td>Vendas</td>
            <td>18/03/2002</td>
          </tr>
        </tbody>
      </table>
    </div>
  )
}
