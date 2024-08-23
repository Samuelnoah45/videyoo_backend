CREATE TABLE "influencer"."influencer_reviews" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "user_id" uuid, "influencer_id" uuid NOT NULL, "rate" numeric NOT NULL, "content" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("influencer_id") REFERENCES "influencer"."influencers"("id") ON UPDATE restrict ON DELETE restrict, FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict);
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
CREATE TRIGGER "set_influencer_influencer_reviews_updated_at"
BEFORE UPDATE ON "influencer"."influencer_reviews"
FOR EACH ROW
EXECUTE PROCEDURE "influencer"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_influencer_influencer_reviews_updated_at" ON "influencer"."influencer_reviews"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
