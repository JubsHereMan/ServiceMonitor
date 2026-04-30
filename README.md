# ServiceMonitor

Projeto desenvolvido em Go para monitoramento de serviços e sites, verificando disponibilidade e status HTTP em tempo real.

A aplicação possui uma interface visual simples para acompanhamento no navegador e também uma rota em JSON para consumo via API.

## 🚀 Objetivo

O objetivo do projeto é permitir o monitoramento rápido de serviços externos, retornando informações como:

- Status de disponibilidade
- Código HTTP da resposta
- Sit atual do serviço (ONLINE / OFFLINE)
- Validação de acesso ao endpoint

Além disso, o projeto apresenta essas informações em uma interface web amigável.

---

## 🛠 Tecnologias utilizadas

- Go (Golang)
- Fiber
- HTML
- CSS
- JavaScript
- Fetch API

---

## ⚙️ Funcionalidades

### Monitoramento de serviços

A aplicação realiza verificações automáticas em URLs configuradas e retorna:

- Nome do serviço  
- Status da conexão  
- Código HTTP  
- Resultado da validação  

### Interface Web

A rota principal (`/`) renderiza uma dashboard visual para acompanhar os serviços monitorados.

### API JSON

A rota `/api` retorna todas as informações em formato JSON.

Ideal para:

- dashboards  
- integrações  
- observabilidade  
- automações  

---

## 🔗 Endpoints

### Página visual

```http
GET /
Renderiza a interface web com o monitoramento visual dos serviços.
```
## API JSON

```http
GET /api
[
  {
    "service": "youtube.com",
    "status": "ONLINE",
    "code": 1
  },
  {
    "service": "spotify.com",
    "status": "NÃO POSSÍVEL ACESSAR SITE",
    "code": 0
  }
]
```


##  Como executar

### Clone o repositório

```bash
git clone https://github.com/JubsHereMan/ServiceMonitor.git
```

### Acesse a pasta

```bash
cd ServiceMonitor
```

### Instale as dependências

```bash
go mod tidy
```

### Execute o projeto


```bash
go run cmd/server/main.go
```


### 🌐 Acesse no navegador

```
http://localhost:3000/
```
