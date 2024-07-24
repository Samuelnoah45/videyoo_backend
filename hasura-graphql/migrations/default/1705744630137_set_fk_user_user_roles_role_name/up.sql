alter table "user"."user_roles"
  add constraint "user_roles_role_name_fkey"
  foreign key ("role_name")
  references "user"."roles"
  ("role_name") on update restrict on delete restrict;
