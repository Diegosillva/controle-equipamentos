const BASE_URL = "http://localhost:8080";

export async function GetEquipamentos() {
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/produto/list`);
    if (!response.ok) {
        throw new Error("Error ao buscar equipamentos");
    }
    return await response.json();
}

export async function GetByProdutoEquipamentos() {
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/produto`);
    if (!response.ok) {
        throw new Error("Error ao buscar equipamentos");
    }
    return await response.json();
}

export async function CriarEquipamentos(payload) {
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/produto/create`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
    });
    if (!response.ok) {
        throw new Error("Erro ao criar equipamento");
    }
    return await response.json();
}

export async function AtualizarEquipamentos(id, dados) {
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/produto/edit`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id, ...dados})
    });
    if (!response.ok) {
        const errorText = await response.text();
        console.error("Erro detalhado:", errorText)
        throw new Error("Erro ao atualizar equipamento.");
    }
    return await response.json(); 
}

export async function DeletarEquipamentos(id) {
    try {
        const response = await fetch(`${BASE_URL}/equipamentos/api/v1/produto/delete?id=${id}`, {
            method: "DELETE",
        })
        if (!response.ok) {
            throw new Error("Erro ao deletar equipamento");
        }
        return true;

    } catch (err) {
        console.error("Erro ao deletar da API", err);
        throw err;
    }
}

