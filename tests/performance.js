import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 }, // Ramp-up para 20 usuários
    { duration: '1m', target: 20 },  // Manter 20 usuários por 1 minuto
    { duration: '30s', target: 0 },  // Ramp-down para 0 usuários
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% das requisições devem ser mais rápidas que 500ms
    http_req_failed: ['rate<0.01'],   // Taxa de falha deve ser menor que 1%
  },
};

export default function () {
  // Teste de criação de conta
  const createAccountRes = http.post('http://localhost:8080/accounts', JSON.stringify({
    name: 'Test User',
    balance: 1000,
  }), {
    headers: { 'Content-Type': 'application/json' },
  });
  check(createAccountRes, {
    'status is 201': (r) => r.status === 201,
  });

  // Extrair o ID da conta criada
  const accountId = createAccountRes.json('id');

  // Teste de depósito
  const depositRes = http.post(`http://localhost:8080/accounts/${accountId}/deposit`, JSON.stringify({
    amount: 100,
  }), {
    headers: { 'Content-Type': 'application/json' },
  });
  check(depositRes, {
    'status is 200': (r) => r.status === 200,
  });

  // Teste de consulta de saldo
  const balanceRes = http.get(`http://localhost:8080/accounts/${accountId}/balance`);
  check(balanceRes, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(1); // Espera 1 segundo entre as iterações
} 