CREATE TABLE "user"."user_roles" ("user_id" uuid NOT NULL, "role_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "created_by" uuid NOT NULL, "updated_at" timestamptz NOT NULL DEFAULT now(), "id" uuid NOT NULL DEFAULT gen_random_uuid(), PRIMARY KEY ("id") , FOREIGN KEY ("user_id") REFERENCES "user"."users"("id") ON UPDATE cascade ON DELETE cascade, FOREIGN KEY ("role_id") REFERENCES "user"."roles"("id") ON UPDATE cascade ON DELETE cascade, UNIQUE ("user_id", "role_id"));
CREATE EXTENSION IF NOT EXISTS pgcrypto;
