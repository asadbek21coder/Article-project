CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "article" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    "title" varchar,
    "body" text,
    "author_id" UUID,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS "unique_title_on_article" ON "article" ("title");

CREATE TABLE IF NOT EXISTS "author" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    "firstname" varchar,
    "lastname" varchar,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS "unique_firstname_lastname_on_author" ON "author" ("firstname", "lastname");

ALTER TABLE "article" ADD CONSTRAINT "article_author_id_fkey" FOREIGN KEY ("author_id") REFERENCES "author" ("id") ON DELETE CASCADE;