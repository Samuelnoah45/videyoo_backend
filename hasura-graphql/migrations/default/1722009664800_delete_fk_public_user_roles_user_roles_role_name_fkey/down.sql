alter table "public"."user_roles"
  add constraint "user_roles_role_name_fkey"
  foreign key ("role_name")
  references "public"."roles"
  ("id") on update cascade on delete cascade;
