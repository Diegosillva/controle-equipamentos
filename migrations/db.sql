CREATE TABLE cadastro_equipamentos (
    id SERIAL PRIMARY KEY,
    produto TEXT,
    equipamento TEXT,
    modelo TEXT,
    numero_de_serie TEXT,
    serial_dsp TEXT,
    descricao TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
