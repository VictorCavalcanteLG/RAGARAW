CREATE TYPE "func" AS ENUM (
  'Vendedor',
  'Tecnico',
  'Supervisor'
);

CREATE TYPE "tipo_servico" AS ENUM (
  'Manutencao',
  'Instalacao'
);

CREATE TABLE "cliente" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "CPF" char(11)
);

CREATE TABLE "funcionario" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar,
  "codigo" varchar UNIQUE,
  "salario" float,
  "funcao" func,
  "senha" varchar,
  "filial_id" int
);

CREATE TABLE "filial" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "cidade" varchar(45) NOT NULL
);

CREATE TABLE "produto" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "codigo" varchar UNIQUE NOT NULL
);

CREATE TABLE "estoque" (
  "filial_id" int,
  "produto_id" int,
  "quantidate" int DEFAULT 0,
  "preco" float,
  PRIMARY KEY ("filial_id", "produto_id")
);

CREATE TABLE "venda" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "filial_id" int,
  "produto_id" int,
  "cliente_id" int,
  "funcionario_id" int,
  "contrato_id" int,
  "dataHora" datetime
);

CREATE TABLE "contrato" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "data" date,
  "servico" tipo_servico,
  "preco" float,
  "data_inicio" date,
  "cliente_id" int
);

CREATE TABLE "servico" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "data" date,
  "relatorio_id" int
);

CREATE TABLE "foto" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "relatorio_id" int
);

CREATE TABLE "documento" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "relatorio_id" int
);

CREATE TABLE "relatorio" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "funcionario_id" int
);

CREATE TABLE "funcionarios_servicos" (
  "funcionario_id" int,
  "servico_id" int
);

CREATE INDEX ON "venda" ("filial_id", "produto_id");

ALTER TABLE "venda" ADD FOREIGN KEY ("cliente_id") REFERENCES "cliente" ("id");

ALTER TABLE "venda" ADD FOREIGN KEY ("funcionario_id") REFERENCES "funcionario" ("id");

ALTER TABLE "estoque" ADD FOREIGN KEY ("produto_id") REFERENCES "produto" ("id");

ALTER TABLE "estoque" ADD FOREIGN KEY ("filial_id") REFERENCES "filial" ("id");

ALTER TABLE "foto" ADD FOREIGN KEY ("relatorio_id") REFERENCES "relatorio" ("id");

ALTER TABLE "documento" ADD FOREIGN KEY ("relatorio_id") REFERENCES "relatorio" ("id");

ALTER TABLE "servico" ADD FOREIGN KEY ("relatorio_id") REFERENCES "relatorio" ("id");

ALTER TABLE "funcionarios_servicos" ADD FOREIGN KEY ("servico_id") REFERENCES "servico" ("id");

ALTER TABLE "funcionarios_servicos" ADD FOREIGN KEY ("funcionario_id") REFERENCES "funcionario" ("id");

ALTER TABLE "contrato" ADD FOREIGN KEY ("cliente_id") REFERENCES "cliente" ("id");

ALTER TABLE "relatorio" ADD FOREIGN KEY ("funcionario_id") REFERENCES "funcionario" ("id");

ALTER TABLE "funcionario" ADD FOREIGN KEY ("filial_id") REFERENCES "filial" ("id");

ALTER TABLE "venda" ADD FOREIGN KEY ("contrato_id") REFERENCES "contrato" ("id");

ALTER TABLE "venda" ADD FOREIGN KEY ("filial_id", "produto_id") REFERENCES "estoque" ("filial_id", "produto_id");