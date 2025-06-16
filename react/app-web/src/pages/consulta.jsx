import { useEffect, useState } from "react"
import { GetEquipamentos } from "../api"
import { Link } from "react-router-dom";

export default function Consulta() {
    const [equipamentos, setEquipamentos] = useState([])
    const [buscaModelo, setBuscaModelo] = useState("")

    useEffect(() => {
        GetEquipamentos()
            .then((res) => setEquipamentos(res))
            .catch((err) => console.error("Erro ao busca equipamentos:", err));
    }, [])

    const buscaProduto = equipamentos.filter((item) =>
        item.produto.toLowerCase().includes(buscaModelo.toLowerCase())
    );

    const deletarEquipamento = (id) => {
        const novaLista = equipamentos.filter(item => item.id !== id);
        setEquipamentos(novaLista)
    }


    return <>
        <div className="relative flex justify-center items-center  py-4">
            <h1 className="font-(family-name:--font-primary) 
                text-5xl text-gray-700 dark:text-gray font-bold text-center ">
                Equipamentos Cadastrado
            </h1>
            <Link to="/cadastro">
                <button className="absolute right-6 top-1/2 cursor-pointer
                    -translate-y-1/2 bg-blue-600 hover:bg-blue-700 text-white
                    text-base font-semibold py-2 px-4 rounded-lg shadow"
                    type="button">
                    Cadastrar
                </button>
            </Link>
        </div>
        <div className="overflow-x-auto p-4 dark:bg-gray-200">
            <table className="min-w-full table-auto boder border-gray-200
                dark:border-gray-700">
                <thead className="bg-gray-100 dark:bg-gray-700">
                    <tr className="text-left text-sm font-semibold text-gray-700
                        text-gray-700 dark:text-gray-100">
                        <th className="px-4 py-2">Produto</th>
                        <th className="px-4 py-2">Equipamento</th>
                        <th className="px-4 py-2">Modelo</th>
                        <th className="px-4 py-2">Numero Serie</th>
                        <th className="px-4 py-2">Serial DSP</th>
                        <th className="px-4 py-2">Localização</th>
                        <th className="px-4 py-2">Status</th>
                        <th className="px-4 py-2">Descrição</th>
                        <th>
                            <input
                                className="bg-write border rounded px-1 py-1 w-full -ml-2"
                                placeholder="buscar:"
                                value={buscaModelo}
                                onChange={(e) => setBuscaModelo(e.target.value)}
                            >
                            </input>
                        </th>
                    </tr>
                </thead>
                <tbody className="text-sm">
                    {buscaProduto.map((item) => (
                        <tr key={item.id}
                            className="border-t border-gray-200 
                            dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-300">
                            <td className="px-4 py-2">{item.produto}</td>
                            <td className="px-4 py-2">{item.equipamento}</td>
                            <td className="px-4 py-2">{item.modelo}</td>
                            <td className="px-4 py-2">{item.numero_serie}</td>
                            <td className="px-4 py-2">{item.serial_dsp}</td>
                            <td className="px-4 py-2">{item.localizacao}</td>
                            <td className="px-4 py-2">{item.status}</td>
                            <td className="px-4 py-2">{item.descricao}</td>
                            <td>
                                <button 
                                    type="button"
                                    onClick={() => deletarEquipamento(item.id)}
                                    className="dark:bg-gray-600 hover:bg-gray-50
                                    rounded px-1 py-1 cursor-pointer">
                                    Deletar
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    </>
}
