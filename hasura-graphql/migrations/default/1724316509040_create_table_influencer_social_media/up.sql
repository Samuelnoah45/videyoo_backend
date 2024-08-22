CREATE TABLE "influencer"."social_media" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "icon_url" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") );
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
CREATE TRIGGER "set_influencer_social_media_updated_at"
BEFORE UPDATE ON "influencer"."social_media"
FOR EACH ROW
EXECUTE PROCEDURE "influencer"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_influencer_social_media_updated_at" ON "influencer"."social_media"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
