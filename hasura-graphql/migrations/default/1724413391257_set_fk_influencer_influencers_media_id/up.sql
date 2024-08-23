alter table "influencer"."influencers"
  add constraint "influencers_media_id_fkey"
  foreign key ("media_id")
  references "public"."media"
  ("id") on update cascade on delete cascade;
