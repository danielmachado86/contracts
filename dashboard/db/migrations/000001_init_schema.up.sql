CREATE TYPE "period_units" AS ENUM (
  'days',
  'months',
  'years'
);

CREATE TABLE "contracts" (
  "id" bigserial PRIMARY KEY,
  "template" int NOT NULL DEFAULT 1
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
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "parties" (
  "user_id" bigint NOT NULL,
  "contract_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "contract_id")
);

CREATE INDEX ON "period_params" ("contract_id");

CREATE INDEX ON "price_params" ("contract_id");

CREATE INDEX ON "time_params" ("contract_id");

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "parties" ("user_id");

CREATE INDEX ON "parties" ("contract_id");

ALTER TABLE "period_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "price_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "time_params" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "parties" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "parties" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");
