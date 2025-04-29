# Bank Server API 🏦

![Go Tests](https://github.com/Luis-Andrei/bank-server/actions/workflows/tests-and-deploy.yml/badge.svg)
[![Coverage](https://codecov.io/gh/Luis-Andrei/bank-server/branch/main/graph/badge.svg)](https://codecov.io/gh/Luis-Andrei/bank-server)
[![Deploy Status](https://vercelbadge.vercel.app/Luis-Andrei/bank-server)](https://vercel.com)

API de sistema bancário desenvolvida em Go com testes automatizados e CI/CD.

## 🚀 Funcionalidades

- Criação de contas
- Consulta de saldo
- Depósito
- Saque
- Testes automatizados
- CI/CD com GitHub Actions
- Deploy automático para Vercel
- Containerização com Docker
- Notificações de falha no Slack e Discord
- Testes de Performance com k6

## 🛠️ Tecnologias

- Go 1.21
- PostgreSQL
- GitHub Actions
- Codecov
- Vercel
- Docker
- Slack e Discord (notificações)
- k6 (testes de performance)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Docker (para rodar o PostgreSQL)
- Git
- k6 (para testes de performance)
- Node.js (para o Vercel CLI)

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/Luis-Andrei/bank-server.git
cd bank-server
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
export DB_USER=postgres
export DB_PASS=yourpassword
export DB_NAME=bankdb
export DB_HOST=localhost
```

## 🧪 Testes

Para executar os testes localmente:
```bash
go test ./handlers -v
```

Para ver a cobertura de testes:
```bash
go test ./handlers -cover
```

Para executar testes de performance:
```bash
k6 run tests/performance.js
```

## 🐳 Docker

Para construir a imagem Docker:
```bash
docker build -t bank-server .
```

Para rodar o container:
```bash
docker run -p 8080:8080 bank-server
```

## 📦 CI/CD

O projeto utiliza GitHub Actions para:
- Executar testes automatizados
- Gerar relatório de cobertura
- Fazer deploy automático para Vercel
- Criar e publicar imagem Docker
- Enviar notificações de falha para Slack e Discord
- Executar testes de performance com k6

## 🔐 Secrets Necessários

Para que o CI/CD funcione corretamente, você precisa configurar os seguintes secrets no GitHub:

- `CODECOV_TOKEN`: Token do Codecov
- `VERCEL_TOKEN`: Token do Vercel
- `SLACK_WEBHOOK_URL`: URL do webhook do Slack
- `DISCORD_WEBHOOK_URL`: URL do webhook do Discord

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Configuração do Ambiente

### 1. Banco de Dados PostgreSQL

Você pode rodar o PostgreSQL localmente usando Docker:

```bash
docker run --rm -d -p 5432:5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=yourpassword \
  -e POSTGRES_DB=bankdb \
  postgres:15
```

## Estrutura do Projeto

- `handlers/`: Contém os handlers HTTP e seus testes