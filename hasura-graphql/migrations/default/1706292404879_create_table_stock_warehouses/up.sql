CREATE TABLE "stock"."warehouses" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "location" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") );
CREATE EXTENSION IF NOT EXISTS pgcrypto;
