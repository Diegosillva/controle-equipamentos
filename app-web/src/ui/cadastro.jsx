
function Cadastro (){
    return <>
            <div className="border border-black-500 p-4 m-auto max-w-md mt-10">
                <h1 className="text-3xl font-bold text-center p-2">Controle de Equipamentos</h1>
                <form className="flex flex-col gap-4">
                    <input className="border rounded px-2 py-1" type="text" placeholder="Produto" />
                    <input className="border rounded px-2 py-1" type="text" placeholder="Equipamento"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Modelo"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Numero de Serie"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Numero DSP"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Localização"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Status"/>
                    <input className="border rounded px-2 py-1" type="text" placeholder="Descrição"/>
                    <button type="submit">Cadastrar</button>
                </form>
            </div>
        </>
}
export default Cadastro



