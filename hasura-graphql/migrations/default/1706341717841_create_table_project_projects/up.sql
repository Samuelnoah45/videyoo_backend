CREATE TABLE "project"."projects" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "project_manager_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "status" Text NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("project_manager_id") REFERENCES "user"."users"("id") ON UPDATE restrict ON DELETE restrict);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
