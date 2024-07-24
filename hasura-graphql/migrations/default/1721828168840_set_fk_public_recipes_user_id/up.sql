alter table "public"."recipes"
  add constraint "recipes_user_id_fkey"
  foreign key ("user_id")
  references "public"."recipes"
  ("id") on update cascade on delete cascade;
