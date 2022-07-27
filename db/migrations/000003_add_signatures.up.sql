CREATE TABLE "signatures" (
  "username" varchar NOT NULL,
  "contract_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "contract_id")
);

CREATE INDEX ON "signatures" ("username");
CREATE INDEX ON "signatures" ("contract_id");
ALTER TABLE "signatures" ADD FOREIGN KEY ("username", "contract_id") REFERENCES "parties" ("username", "contract_id");