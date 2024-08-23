CREATE TABLE "content"."example_videos" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "media_id" uuid NOT NULL, "example_video_type" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("media_id") REFERENCES "public"."media"("id") ON UPDATE restrict ON DELETE restrict);
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
CREATE TRIGGER "set_content_example_videos_updated_at"
BEFORE UPDATE ON "content"."example_videos"
FOR EACH ROW
EXECUTE PROCEDURE "content"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_content_example_videos_updated_at" ON "content"."example_videos"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
