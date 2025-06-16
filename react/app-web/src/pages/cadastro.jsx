import { useState } from "react";
import { CriarEquipamentos } from "../api";
import { Link } from "react-router-dom";

export default function Cadastro() {
    const [form, setForm] = useState({
        produto: "",
        equipamento: "",
        modelo: "",
        numero_serie: "",
        serial_dsp: "",
        localizacao: "",
        status: "",
        descricao: "",
    });

    const [msg, setMsg] = useState("")

    const handlerValue = (e) => {
        const { name, value } = e.target;
        setForm((prev) => ({ ...prev, [name]: value.toUpperCase() }));
    }


    const click = async (e) => {
        e.preventDefault();
        try {
            const resposta = CriarEquipamentos(form)
            console.log("Salvo com sucesso.", resposta)
            alert("Equipamento cadastrado com sucesso.")
            setForm({
                produto: "",
                equipamento: "",
                modelo: "",
                numero_serie: "",
                serial_dsp: "",
                localizacao: "",
                status: "",
                descricao: "",
            });
        } catch (err) {
            console.error("Erro ao salvar", err)
            setError("Erro ao salvar Equipamento")
        }
        console.log("Formulario Enviado")
    };


    return (
        <div className="flex items-center justify-center h-screen bg-(--color-body)">
            <div className="bg-white dark:bg-gray-800 rounded-lg px-6 py-8 
                        ring shadow-xl ring-gray-900/5 ">
                <h1 className="font-(family-name:--font-primary) text-5xl text-white font-bold text-center p-2 pb-6">
                    Controle de Equipamentos
                </h1>
                {msg && <div className="flex justify-center text-green-300 text-2xl mb-4 items-center">
                    Sucesso: {msg}
                </div>}

                <form
                    onSubmit={click}
                    className="dark:text-white flex flex-col gap-4"
                >
                    <input
                        name="produto"
                        value={form.produto}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Produto"
                    />
                    <input
                        name="equipamento"
                        value={form.equipamento}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Equipamento"
                    />
                    <input
                        name="modelo"
                        value={form.modelo}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Modelo"
                    />
                    <input
                        name="numero_serie"
                        value={form.numero_serie}
                        onChange={handlerValue}
                        required
                        maxLength={14}
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Numero de Serie"
                    />
                    <input
                        name="serial_dsp"
                        value={form.serial_dsp}
                        onChange={handlerValue}
                        required
                        maxLength={4}
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Numero DSP"
                    />
                    <input
                        name="localizacao"
                        value={form.localizacao}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Localização"
                    />
                    <input
                        name="status"
                        value={form.status}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Status"
                    />
                    <input
                        name="descricao"
                        value={form.descricao}
                        onChange={handlerValue}
                        required
                        className="border rounded px-2 py-1"
                        type="text"
                        placeholder="Descrição"
                    />
                    <div className="flex justify-center gap-4 mt-4">
                        <button className="bg-green-600 hover:bg-green-700 
                            text-white font-semibold py-3 px-6 rounded-xl 
                            shadow text-lg cursor-pointer"
                            type="submit"
                        >
                            Cadastrar
                        </button>

                        <Link to="/">
                            <button className=" bg-blue-600 hover:bg-blue-700 text-white
                    text-xl font-semibold py-3 px-6 rounded-xl shadow"
                                type="button">
                                Consultar
                            </button>
                        </Link>
                    </div>
                </form>
            </div>
        </div>
    )
}



