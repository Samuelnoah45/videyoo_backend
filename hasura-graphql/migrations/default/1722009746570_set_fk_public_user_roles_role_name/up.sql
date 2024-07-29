alter table "public"."user_roles"
  add constraint "user_roles_role_name_fkey"
  foreign key ("role_name")
  references "public"."roles"
  ("role_name") on update no action on delete no action;
