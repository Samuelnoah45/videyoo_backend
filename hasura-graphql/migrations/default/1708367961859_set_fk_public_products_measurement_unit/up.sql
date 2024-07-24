alter table "public"."products"
  add constraint "products_measurement_unit_fkey"
  foreign key ("measurement_unit")
  references "basic"."measurment_units"
  ("unit") on update restrict on delete restrict;
