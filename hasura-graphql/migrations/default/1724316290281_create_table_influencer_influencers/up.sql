CREATE TABLE "influencer"."influencers" ("first_name" text NOT NULL, "last_name" text NOT NULL, "bio" text NOT NULL, "gender" text NOT NULL, "age" text, "location" text NOT NULL, "badge_number" text NOT NULL, "id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "created_by" uuid, "updated_by" uuid, PRIMARY KEY ("id") , FOREIGN KEY ("created_by") REFERENCES "public"."users"("id") ON UPDATE set null ON DELETE set null, FOREIGN KEY ("updated_by") REFERENCES "public"."users"("id") ON UPDATE set null ON DELETE set null);
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
CREATE TRIGGER "set_influencer_influencers_updated_at"
BEFORE UPDATE ON "influencer"."influencers"
FOR EACH ROW
EXECUTE PROCEDURE "influencer"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_influencer_influencers_updated_at" ON "influencer"."influencers"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
