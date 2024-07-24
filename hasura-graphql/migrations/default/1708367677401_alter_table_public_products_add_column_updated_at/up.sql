alter table "public"."products" add column "updated_at" timestamptz
 not null default now();
