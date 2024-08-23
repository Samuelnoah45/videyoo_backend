alter table "influencer"."works"
  add constraint "works_media_id_fkey"
  foreign key ("media_id")
  references "public"."media"
  ("id") on update cascade on delete cascade;
