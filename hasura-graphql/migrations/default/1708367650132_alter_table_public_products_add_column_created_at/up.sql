alter table "public"."products" add column "created_at" timestamptz
 not null default now();
