
CREATE TABLE "contracts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "template" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hashed" varchar NOT NULL,
  "changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "parties" (
  "username" varchar NOT NULL,
  "role" varchar NOT NULL,
  "contract_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "contract_id")
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "parties" ("username");

CREATE INDEX ON "parties" ("contract_id");

ALTER TABLE "parties" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "parties" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id") ON DELETE CASCADE;
