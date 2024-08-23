CREATE TABLE "content"."FAQ" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "question" text NOT NULL, "answer" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") );
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
CREATE TRIGGER "set_content_FAQ_updated_at"
BEFORE UPDATE ON "content"."FAQ"
FOR EACH ROW
EXECUTE PROCEDURE "content"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_content_FAQ_updated_at" ON "content"."FAQ"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
