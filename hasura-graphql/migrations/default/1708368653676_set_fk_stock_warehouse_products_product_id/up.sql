alter table "stock"."warehouse_products"
  add constraint "warehouse_products_product_id_fkey"
  foreign key ("product_id")
  references "public"."products"
  ("id") on update restrict on delete restrict;
