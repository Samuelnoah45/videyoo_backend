CREATE TABLE "content"."blogs" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "media_id" uuid NOT NULL, "date" timestamptz NOT NULL, "title" text NOT NULL, "content" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("media_id") REFERENCES "public"."media"("id") ON UPDATE restrict ON DELETE restrict);
CREATE OR REPLACE FUNCTION "content"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_content_blogs_updated_at"
BEFORE UPDATE ON "content"."blogs"
FOR EACH ROW
EXECUTE PROCEDURE "content"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_content_blogs_updated_at" ON "content"."blogs"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
