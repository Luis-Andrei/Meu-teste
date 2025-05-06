// Criar usuário root
db.createUser({
  user: process.env.MONGODB_ROOT_USERNAME,
  pwd: process.env.MONGODB_ROOT_PASSWORD,
  roles: [
    { role: 'root', db: 'admin' }
  ]
});

// Criar banco de dados auth
db = db.getSiblingDB('auth');
db.createUser({
  user: 'auth_user',
  pwd: 'auth_password',
  roles: [
    { role: 'readWrite', db: 'auth' }
  ]
});

// Criar banco de dados ecommerce
db = db.getSiblingDB('ecommerce');
db.createUser({
  user: 'ecommerce_user',
  pwd: 'ecommerce_password',
  roles: [
    { role: 'readWrite', db: 'ecommerce' }
  ]
});

// Inicializar replica set
rs.initiate({
  _id: 'rs0',
  members: [
    { _id: 0, host: 'localhost:27017' }
  ]
}); 