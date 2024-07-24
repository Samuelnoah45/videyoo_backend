CREATE TABLE "public"."ingratiates" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "quantity" Integer NOT NULL, "unit" text NOT NULL, "recipe_id" uuid NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("recipe_id") REFERENCES "public"."recipes"("id") ON UPDATE cascade ON DELETE cascade);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
