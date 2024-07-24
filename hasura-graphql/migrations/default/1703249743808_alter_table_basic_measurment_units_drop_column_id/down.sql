alter table "basic"."measurment_units" alter column "id" set default gen_random_uuid();
alter table "basic"."measurment_units" alter column "id" drop not null;
alter table "basic"."measurment_units" add column "id" uuid;
