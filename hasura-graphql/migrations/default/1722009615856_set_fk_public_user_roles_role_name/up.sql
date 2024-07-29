alter table "public"."user_roles" drop constraint "user_roles_role_id_fkey",
  add constraint "user_roles_role_name_fkey"
  foreign key ("role_name")
  references "public"."roles"
  ("id") on update cascade on delete cascade;
