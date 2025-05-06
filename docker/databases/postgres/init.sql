-- Criar extensões
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Criar schema do banco
CREATE SCHEMA IF NOT EXISTS bank;

-- Criar tabela de contas
CREATE TABLE IF NOT EXISTS bank.accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0,
    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Criar tabela de transações
CREATE TABLE IF NOT EXISTS bank.transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_account_id UUID NOT NULL REFERENCES bank.accounts(id),
    to_account_id UUID NOT NULL REFERENCES bank.accounts(id),
    amount DECIMAL(15,2) NOT NULL,
    type VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Criar índices
CREATE INDEX IF NOT EXISTS idx_accounts_account_number ON bank.accounts(account_number);
CREATE INDEX IF NOT EXISTS idx_transactions_from_account ON bank.transactions(from_account_id);
CREATE INDEX IF NOT EXISTS idx_transactions_to_account ON bank.transactions(to_account_id);
CREATE INDEX IF NOT EXISTS idx_transactions_created_at ON bank.transactions(created_at);

-- Criar função para atualizar updated_at
CREATE OR REPLACE FUNCTION bank.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Criar triggers
CREATE TRIGGER update_accounts_updated_at
    BEFORE UPDATE ON bank.accounts
    FOR EACH ROW
    EXECUTE FUNCTION bank.update_updated_at_column();

CREATE TRIGGER update_transactions_updated_at
    BEFORE UPDATE ON bank.transactions
    FOR EACH ROW
    EXECUTE FUNCTION bank.update_updated_at_column(); 