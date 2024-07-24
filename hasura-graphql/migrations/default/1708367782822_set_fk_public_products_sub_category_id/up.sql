alter table "public"."products"
  add constraint "products_sub_category_id_fkey"
  foreign key ("sub_category_id")
  references "public"."sub_categories"
  ("id") on update restrict on delete restrict;
