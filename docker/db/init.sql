CREATE SCHEMA IF NOT EXISTS bank;

CREATE TABLE IF NOT EXISTS bank.accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    balance NUMERIC(10, 2) NOT NULL DEFAULT 0.00
);

INSERT INTO bank.accounts (balance) VALUES (1000.00);
INSERT INTO bank.accounts (balance) VALUES (12341.32);
INSERT INTO bank.accounts (balance) VALUES (2341.45);
INSERT INTO bank.accounts (balance) VALUES (99090.12);
INSERT INTO bank.accounts (balance) VALUES (654643.05);

CREATE USER api_user WITH PASSWORD 'api_password';

GRANT USAGE ON SCHEMA bank TO api_user;

GRANT SELECT, UPDATE ON bank.accounts TO api_user;