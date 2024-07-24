CREATE TABLE "public"."roles" ("role_name" text NOT NULL, "id" uuid NOT NULL DEFAULT gen_random_uuid(), "description" text NOT NULL, PRIMARY KEY ("role_name") , UNIQUE ("role_name"));
CREATE EXTENSION IF NOT EXISTS pgcrypto;
