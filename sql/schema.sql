CREATE TYPE "func" AS ENUM (
  'Vendedor',
  'Tecnico',
  'Supervisor'
);

CREATE TYPE "tipo_servico" AS ENUM (
  'Manutencao',
  'Instalacao'
);

CREATE TABLE "Cliente" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "CPF" char(11)
);

CREATE TABLE "Funcionario" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar,
  "codigo" char(5),
  "salario" float,
  "funcao" func,
  "filial_id" int
);

CREATE TABLE "Filial" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "cidade" varchar(45) NOT NULL
);

CREATE TABLE "Produto" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "nome" varchar(45) NOT NULL,
  "codigo" char(5) NOT NULL
);

CREATE TABLE "Estoque" (
  "filial_id" int,
  "produto_id" int,
  "quantidate" int DEFAULT 0,
  "preco" float,
  PRIMARY KEY ("filial_id", "produto_id")
);

CREATE TABLE "Venda" (
  "id" int PRIMARY KEY,
  "filial_id" int,
  "produto_id" int,
  "cliente_id" int,
  "funcionario_id" int,
  "contrato_id" int,
  "dataHora" datetime
);

CREATE TABLE "Contrato" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "data" date,
  "servico" tipo_servico,
  "preco" float,
  "data_inicio" date,
  "cliente_id" int
);

CREATE TABLE "Servico" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "data" date,
  "relatorio_id" int
);

CREATE TABLE "Foto" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "relatorio_id" int
);

CREATE TABLE "Documento" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "relatorio_id" int
);

CREATE TABLE "Relatorio" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "funcionario_id" int
);

CREATE TABLE "Funcionarios_Servicos" (
  "funcionario_id" int,
  "servico_id" int
);

ALTER TABLE "Venda" ADD FOREIGN KEY ("cliente_id") REFERENCES "Cliente" ("id");

ALTER TABLE "Venda" ADD FOREIGN KEY ("funcionario_id") REFERENCES "Funcionario" ("id");

ALTER TABLE "Estoque" ADD FOREIGN KEY ("produto_id") REFERENCES "Produto" ("id");

ALTER TABLE "Estoque" ADD FOREIGN KEY ("filial_id") REFERENCES "Filial" ("id");

ALTER TABLE "Foto" ADD FOREIGN KEY ("relatorio_id") REFERENCES "Relatorio" ("id");

ALTER TABLE "Documento" ADD FOREIGN KEY ("relatorio_id") REFERENCES "Relatorio" ("id");

ALTER TABLE "Servico" ADD FOREIGN KEY ("relatorio_id") REFERENCES "Relatorio" ("id");

ALTER TABLE "Funcionarios_Servicos" ADD FOREIGN KEY ("servico_id") REFERENCES "Servico" ("id");

ALTER TABLE "Funcionarios_Servicos" ADD FOREIGN KEY ("funcionario_id") REFERENCES "Funcionario" ("id");

ALTER TABLE "Contrato" ADD FOREIGN KEY ("cliente_id") REFERENCES "Cliente" ("id");

ALTER TABLE "Relatorio" ADD FOREIGN KEY ("funcionario_id") REFERENCES "Funcionario" ("id");

ALTER TABLE "Funcionario" ADD FOREIGN KEY ("filial_id") REFERENCES "Filial" ("id");

ALTER TABLE "Venda" ADD FOREIGN KEY ("contrato_id") REFERENCES "Contrato" ("id");

ALTER TABLE "Venda" ADD FOREIGN KEY ("filial_id", "produto_id") REFERENCES "Estoque" ("filial_id", "produto_id");

CREATE INDEX ON "Venda" ("filial_id", "produto_id");
