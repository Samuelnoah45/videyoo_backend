CREATE TABLE "influencer"."works" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "influencer_id" uuid NOT NULL, "work_type" text NOT NULL, "media_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "created_by" uuid NOT NULL, "updated_by" uuid NOT NULL, PRIMARY KEY ("id") );
CREATE OR REPLACE FUNCTION "influencer"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_influencer_works_updated_at"
BEFORE UPDATE ON "influencer"."works"
FOR EACH ROW
EXECUTE PROCEDURE "influencer"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_influencer_works_updated_at" ON "influencer"."works"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
