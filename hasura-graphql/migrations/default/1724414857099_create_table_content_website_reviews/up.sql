CREATE TABLE "content"."website_reviews" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "rate_count" numeric NOT NULL, "comment" text NOT NULL, "name" text NOT NULL, "address" text NOT NULL, "social_media_id" uuid NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("social_media_id") REFERENCES "influencer"."social_media"("id") ON UPDATE restrict ON DELETE restrict);
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
CREATE TRIGGER "set_content_website_reviews_updated_at"
BEFORE UPDATE ON "content"."website_reviews"
FOR EACH ROW
EXECUTE PROCEDURE "content"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_content_website_reviews_updated_at" ON "content"."website_reviews"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
