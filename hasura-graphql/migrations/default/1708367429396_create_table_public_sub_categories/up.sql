CREATE TABLE "public"."sub_categories" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "category" uuid NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("category") REFERENCES "public"."categories"("id") ON UPDATE restrict ON DELETE restrict);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
