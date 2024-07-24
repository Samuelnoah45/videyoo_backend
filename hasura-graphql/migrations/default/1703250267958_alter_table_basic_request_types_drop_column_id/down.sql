alter table "basic"."request_types" alter column "id" set default gen_random_uuid();
alter table "basic"."request_types" alter column "id" drop not null;
alter table "basic"."request_types" add column "id" uuid;
