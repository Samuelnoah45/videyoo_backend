alter table "basic"."transaction_types" alter column "id" set default gen_random_uuid();
alter table "basic"."transaction_types" alter column "id" drop not null;
alter table "basic"."transaction_types" add column "id" uuid;
