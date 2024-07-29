alter table "public"."user_roles" drop constraint "user_roles_user_id_fkey",
  add constraint "user_roles_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update cascade on delete cascade;
