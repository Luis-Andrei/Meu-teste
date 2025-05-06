# Projetos Integrados 🚀

Este repositório integra vários projetos desenvolvidos durante os cursos, criando um ecossistema completo de aplicações.

## 📦 Projetos Integrados

### 1. Sistema de Autenticação (auth-service)
- Frontend: React + TypeScript + Vite
- Backend: Node.js + Express + MongoDB
- Funcionalidades:
  - Autenticação JWT
  - Registro e Login
  - Proteção de rotas
  - Validação de dados

### 2. E-commerce (ecommerce-service)
- Frontend: HTML/CSS/JavaScript
- Backend: Python + Flask
- Funcionalidades:
  - Catálogo de produtos
  - Carrinho de compras
  - Sistema de pagamentos
  - Área do cliente

### 3. Sistema Bancário (bank-service)
- Backend: Go
- Funcionalidades:
  - Criação de contas
  - Transferências
  - Consulta de saldo
  - Histórico de transações

### 4. API de Integração (api-gateway)
- Backend: Go + Node.js
- Funcionalidades:
  - Roteamento de requisições
  - Autenticação centralizada
  - Rate limiting
  - Logging

## 🛠️ Tecnologias Utilizadas

### Frontend
- React 18
- TypeScript
- Vite
- Tailwind CSS
- HTML/CSS/JavaScript

### Backend
- Node.js + Express
- Go
- Python + Flask
- MongoDB
- PostgreSQL

### DevOps
- Docker
- Docker Compose
- GitHub Actions
- Nginx

## 📋 Pré-requisitos

- Node.js 18+
- Go 1.21+
- Python 3.8+
- Docker
- MongoDB
- PostgreSQL

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/projetos-integrados.git
cd projetos-integrados
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

3. Inicie os serviços com Docker Compose:
```bash
docker-compose up -d
```

## 🚀 Executando os Projetos

### Sistema de Autenticação
```bash
cd auth-service
npm install
npm run dev
```

### E-commerce
```bash
cd ecommerce-service
python -m venv .venv
source .venv/bin/activate  # ou .venv\Scripts\activate no Windows
pip install -r requirements.txt
python app.py
```

### Sistema Bancário
```bash
cd bank-service
go mod download
go run main.go
```

### API Gateway
```bash
cd api-gateway
npm install
npm run dev
```

## 📝 Estrutura do Projeto

```
projetos-integrados/
├── auth-service/           → Sistema de Autenticação
│   ├── client/            → Frontend React
│   └── server/            → Backend Node.js
│
├── ecommerce-service/      → E-commerce
│   ├── frontend/          → Frontend HTML/CSS/JS
│   └── backend/           → Backend Python
│
├── bank-service/          → Sistema Bancário
│   ├── handlers/          → Handlers HTTP
│   ├── models/            → Modelos de dados
│   └── db/                → Configuração do banco
│
├── api-gateway/           → API Gateway
│   ├── routes/            → Rotas da API
│   └── middleware/        → Middlewares
│
├── docker/                → Configurações Docker
│   ├── nginx/             → Configuração Nginx
│   └── databases/         → Configurações dos bancos
│
└── docs/                  → Documentação
    ├── api/               → Documentação da API
    └── architecture/      → Arquitetura do sistema
```

## 🔄 CI/CD

O projeto utiliza GitHub Actions para:
- Testes automatizados
- Build dos containers
- Deploy automático
- Análise de código

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes. 