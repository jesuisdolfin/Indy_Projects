CREATE TABLE "user_data" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "weight" decimal NOT NULL,
  "height" decimal NOT NULL,
  "age" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "liftentries" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "weight_lifted" decimal NOT NULL,
  "reps" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_data" ("name");

CREATE INDEX ON "liftentries" ("user_id");

COMMENT ON COLUMN "user_data"."weight" IS 'must be positive';

COMMENT ON COLUMN "user_data"."height" IS 'must be positive';

COMMENT ON COLUMN "user_data"."age" IS 'must be positive';

COMMENT ON COLUMN "liftentries"."weight_lifted" IS 'must be positive';

COMMENT ON COLUMN "liftentries"."reps" IS 'must be positive';

ALTER TABLE "liftentries" ADD FOREIGN KEY ("user_id") REFERENCES "user_data" ("id");