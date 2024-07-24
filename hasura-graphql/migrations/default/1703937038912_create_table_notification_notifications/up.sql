CREATE TABLE "notification"."notifications" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "subject" text NOT NULL, "message" text NOT NULL, "user_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("user_id") REFERENCES "user"."users"("id") ON UPDATE restrict ON DELETE restrict);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
