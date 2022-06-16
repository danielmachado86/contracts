CREATE TYPE "templates" AS ENUM (
  'rental',
  'freelance',
  'services'
);

CREATE TYPE "period_units" AS ENUM (
  'days',
  'months',
  'years'
);

CREATE TABLE "contracts" (
  "id" bigserial PRIMARY KEY,
  "template" templates NOT NULL
);

CREATE TABLE "period_params" (
  "id" bigserial PRIMARY KEY,
  "contract_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "value" int NOT NULL,
  "units" period_units NOT NULL
);

CREATE TABLE "price_params" (
  "id" bigserial PRIMARY KEY,
  "contract_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "value" float8 NOT NULL,
  "currency" varchar NOT NULL
);

CREATE TABLE "time_params" (
  "id" bigserial PRIMARY KEY,
  "contract_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "value" timestamptz NOT NULL
);

CREATE TABLE "users" (
  "name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "parties" (
  "username" varchar NOT NULL,
  "contract_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "contract_id")
);

CREATE INDEX ON "period_params" ("contract_id");

CREATE INDEX ON "price_params" ("contract_id");

CREATE INDEX ON "time_params" ("contract_id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "parties" ("username");

CREATE INDEX ON "parties" ("contract_id");

ALTER TABLE "period_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "price_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "time_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "parties" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "parties" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");