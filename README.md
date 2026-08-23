Backend do projeto **imerscafe**, desenvolvido em **Python com FastAPI**.

O serviço concentra a lógica determinística do jogo, sendo responsável pelo processamento das rodadas, validação das ações e gerenciamento do estado oficial da aplicação.

## Responsabilidades

O backend é responsável por:

- Executar as regras formais do jogo;
- Validar hard skills;
- Gerenciar o estado oficial da rodada;
- Validar os dados recebidos;
- Comparar os ingredientes selecionados com a receita esperada;
- Calcular a pontuação;
- Determinar o resultado oficial da rodada;
- Consolidar os dados processados;
- Orquestrar integrações necessárias ao processamento da rodada.

## Stack

- **Python 3.11+**
- **FastAPI** — framework para construção da API;
- **Pydantic** — validação e modelagem dos dados;
- **Uvicorn** — servidor ASGI;
- **Pytest** — testes automatizados.

## Estrutura

```text
backend/
│
├── app/
│   ├── main.py
│   │
│   ├── api/
│   │   └── routes/
│   │
│   ├── models/
│   │
│   ├── schemas/
│   │
│   ├── services/
│   │
│   └── core/
│
├── tests/
│
├── requirements.txt
├── .env.example
└── README.md
````

### `main.py`

Ponto de entrada da aplicação FastAPI.

Responsável pela criação e configuração da aplicação e pelo registro das rotas.

### `api/`

Contém os endpoints e as rotas utilizadas para comunicação com o backend.

### `models/`

Contém os modelos utilizados pela aplicação para representar os dados e entidades do domínio.

### `schemas/`

Contém os schemas Pydantic utilizados para validação dos dados de entrada e saída da API.

### `services/`

Contém a lógica de negócio e o processamento das regras do jogo.

### `core/`

Contém configurações e componentes centrais da aplicação.

## Fluxo de processamento

O processamento de uma rodada ocorre no backend:

```text
Request
   │
   ▼
FastAPI
   │
   ▼
Validação dos dados
   │
   ▼
Regras do jogo
   │
   ├── Validação da receita
   ├── Validação dos ingredientes
   ├── Cálculo da pontuação
   └── Atualização da rodada
   │
   ▼
Resultado oficial
   │
   ▼
Response
```

O resultado da rodada é determinado exclusivamente pelo backend.

## Execução

### Criar ambiente virtual

```bash
python3 -m venv .venv
```

### Ativar ambiente virtual

```bash
source .venv/bin/activate
```

### Instalar dependências

```bash
pip install -r requirements.txt
```

### Executar a aplicação

```bash
uvicorn app.main:app --reload
```

A API estará disponível em:

```text
http://localhost:8000
```

## Documentação da API

O FastAPI disponibiliza automaticamente a documentação interativa da API.

### Swagger UI

```text
http://localhost:8000/docs
```

### ReDoc

```text
http://localhost:8000/redoc
```

## Testes

Executar os testes com:

```bash
pytest
```

## Princípio do Backend

O backend é a fonte de verdade para todas as regras determinísticas do jogo.

```text
Cliente
   │
   │ Dados da ação
   ▼
FastAPI
   │
   │ Validação
   ▼
Regras de negócio
   │
   │ Resultado determinístico
   ▼
Estado oficial da rodada
```
**O cliente fornece os dados da ação. O backend valida, processa e determina o resultado oficial.**

