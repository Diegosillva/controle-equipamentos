import { useEffect, useState } from "react"
import { DeletarEquipamentos, GetEquipamentos } from "../api"
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

    const deleteEquipamento = async (id) => {
        try {
            await DeletarEquipamentos(id);
            const novaLista = equipamentos.filter(item => item.id !== id);
            setEquipamentos(novaLista)
        } catch (err) {
            alert("Erro ao deletar equipamento")
        }
    }

    return <>
        <div className="relative flex justify-center py-4 px-4">
            <h1 className="font-(family-name:--font-primary) text-3xl sm:text-4xl font-bold  text-center">
                Equipamentos Cadastrados
            </h1>
            <Link to="/cadastro">
                <button className="absolute right-4 top-1/2 -translate-y-1/2 
                    bg-blue-600 hover:bg-blue-700 text-white text-base 
                    font-semibold py-2 px-4 rounded-lg shadow">
                    Cadastrar
                </button>
            </Link>
        </div>

        <div className="w-full overflow-x-auto px-2 sm:px-4 py-2 dark:bg-gray-200">
            <table className="min-w-[800px] sm:min-w-full table-auto border border-gray-200 dark:border-gray-700">
                <thead className="bg-gray-100 dark:bg-gray-700">
                    <tr className="text-left text-sm font-semibold text-gray-700 dark:text-gray-100">
                        <th className="px-4 py-2">Produto</th>
                        <th className="px-4 py-2">Equipamento</th>
                        <th className="px-4 py-2">Modelo</th>
                        <th className="px-4 py-2">Número Série</th>
                        <th className="px-4 py-2">Serial DSP</th>
                        <th className="px-4 py-2">Localização</th>
                        <th className="px-4 py-2">Status</th>
                        <th className="px-4 py-2">Descrição</th>
                        <th className="px-4 py-2">
                            <input
                                className="bg-white border text-black rounded px-2 py-1 w-full sm:w-32"
                                placeholder="Buscar"
                                value={buscaModelo}
                                onChange={(e) => setBuscaModelo(e.target.value)}
                                aria-label="Buscar por produto"
                            />
                        </th>
                    </tr>
                </thead>
                <tbody className="text-sm">
                    {buscaProduto.map((item) => (
                        <tr key={item.id}
                            className="border-t border-gray-200 dark:border-gray-600
                            hover:bg-gray-50 dark:hover:bg-gray-300"
                        >
                            <td className="px-4 py-2">{item.produto}</td>
                            <td className="px-4 py-2">{item.equipamento}</td>
                            <td className="px-4 py-2">{item.modelo}</td>
                            <td className="px-4 py-2">{item.numero_serie}</td>
                            <td className="px-4 py-2">{item.serial_dsp}</td>
                            <td className="px-4 py-2">{item.localizacao}</td>
                            <td className="px-4 py-2">{item.status}</td>
                            <td className="px-4 py-2">{item.descricao}</td>
                            <td className="px-4 py-2">
                                <div className="flex flex-col sm:flex-row gap-2">
                                    <button
                                        type="button"
                                        onClick={() => deleteEquipamento(item.id)}
                                        className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
                                        aria-label="Deletar equipamento"
                                    >
                                        Deletar
                                    </button>
                                    <Link to="/cadastro" state={{ item }}>
                                        <button
                                            type="button"
                                            className="bg-yellow-500 hover:bg-yellow-600 text-white px-3 py-1 rounded"
                                            aria-label="Editar equipamento"
                                        >
                                            Editar
                                        </button>
                                    </Link>
                                </div>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    </>
}
