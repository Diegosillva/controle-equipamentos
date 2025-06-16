const BASE_URL = "http://localhost:8080";

export async function GetEquipamentos() {
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/list`);
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
    const response = await fetch(`${BASE_URL}/equipamentos/api/v1/create`,{
        method : "POST",
        headers : { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
    });
    if (!response.ok) {
        throw new Error("Erro ao criar equipamento");
    }
    return await response.json();
}
