CREATE TABLE IF NOT EXISTS "temp" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    "data" text,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);
