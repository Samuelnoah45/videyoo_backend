CREATE TABLE "content"."page_videos" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "order_number" numeric NOT NULL, "media_id" uuid NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("media_id") REFERENCES "public"."media"("id") ON UPDATE restrict ON DELETE restrict);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
