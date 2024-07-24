alter table "public"."sub_categories"
  add constraint "sub_categories_category_id_fkey"
  foreign key ("category_id")
  references "public"."categories"
  ("id") on update restrict on delete restrict;
