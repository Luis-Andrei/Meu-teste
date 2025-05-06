# Meu Projeto - Sistema de Autenticação

Sistema completo de autenticação com React, Node.js, MongoDB e TypeScript.

## 🚀 Tecnologias

### Frontend
- React 18
- TypeScript
- Vite
- React Router DOM
- React Hook Form
- Yup
- Axios
- React Hot Toast
- Tailwind CSS
- Vitest (Testes)
- React Testing Library

### Backend
- Node.js
- Express
- TypeScript
- MongoDB
- Mongoose
- JWT
- Bcrypt
- Yup
- Jest (Testes)
- Supertest

## 📦 Instalação

### Pré-requisitos
- Node.js 18+
- MongoDB
- Git

### Backend
```bash
# Entrar no diretório do backend
cd server

# Instalar dependências
npm install

# Configurar variáveis de ambiente
cp .env.example .env
# Edite o arquivo .env com suas configurações

# Iniciar em desenvolvimento
npm run dev

# Rodar testes
npm test

# Rodar lint
npm run lint
```

### Frontend
```bash
# Entrar no diretório do frontend
cd client

# Instalar dependências
npm install

# Iniciar em desenvolvimento
npm run dev

# Rodar testes
npm test

# Rodar testes com cobertura
npm run test:coverage

# Rodar lint
npm run lint
```

## 🧪 Testes

### Backend
Os testes do backend são escritos com Jest e Supertest. Para executar:

```bash
cd server
npm test
```

### Frontend
Os testes do frontend são escritos com Vitest e React Testing Library. Para executar:

```bash
cd client
npm test
```

Para ver a cobertura de testes:
```bash
npm run test:coverage
```

## 🔄 CI/CD

O projeto usa GitHub Actions para CI/CD. O pipeline inclui:

1. **Testes e Build**
   - Instalação de dependências
   - Execução de testes (frontend e backend)
   - Build do projeto
   - Linting

2. **Deploy** (apenas na branch main)
   - Deploy do backend
   - Deploy do frontend

### Secrets Necessárias
Configure as seguintes secrets no GitHub:
- `MONGODB_URI_TEST`: URI do MongoDB para testes
- `JWT_SECRET_TEST`: Chave secreta para JWT em testes
- `SSH_PRIVATE_KEY`: Chave SSH para deploy
- `SSH_HOST`: Host do servidor de produção

## 📝 Estrutura do Projeto

```
meu-projeto/
├── client/              → Frontend (React + Vite + Tailwind)
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   │   └── ProtectedRoute.tsx
│   │   ├── pages/
│   │   │   ├── Login.tsx
│   │   │   ├── Signup.tsx
│   │   │   └── Dashboard.tsx
│   │   ├── test/
│   │   │   └── setup.ts
│   │   ├── App.tsx
│   │   ├── main.tsx
│   │   └── index.css
│   ├── tailwind.config.js
│   ├── postcss.config.js
│   ├── vite.config.ts
│   ├── vitest.config.ts
│   ├── tsconfig.json
│   └── package.json
│
├── server/              → Backend (Node + Express + MongoDB)
│   ├── controllers/
│   │   └── authController.js
│   ├── middleware/
│   │   ├── authMiddleware.js
│   │   ├── errorMiddleware.js
│   │   └── validationMiddleware.js
│   ├── models/
│   │   └── User.js
│   ├── routes/
│   │   └── authRoutes.js
│   ├── validations/
│   │   └── authValidation.js
│   ├── .env
│   ├── app.js
│   ├── server.js
│   └── package.json
│
└── .github/
    └── workflows/
        └── ci-cd.yml
```

## 🔒 Segurança

- Validação de dados com Yup
- Autenticação com JWT
- Senhas criptografadas com Bcrypt
- Proteção contra CSRF
- Rate limiting
- Sanitização de inputs

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes. 