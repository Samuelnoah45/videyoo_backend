alter table "public"."products"
  add constraint "products_created_by_fkey"
  foreign key ("created_by")
  references "user"."users"
  ("id") on update restrict on delete restrict;
