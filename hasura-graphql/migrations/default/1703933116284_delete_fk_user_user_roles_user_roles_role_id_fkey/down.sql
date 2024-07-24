alter table "user"."user_roles"
  add constraint "user_roles_role_id_fkey"
  foreign key ("role_id")
  references "user"."roles"
  ("id") on update cascade on delete cascade;
