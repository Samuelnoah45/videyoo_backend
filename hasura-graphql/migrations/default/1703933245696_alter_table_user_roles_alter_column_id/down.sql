alter table "user"."roles" rename column "description" to "id";
alter table "user"."roles" alter column "id" set not null;
alter table "user"."roles" alter column "id" set default gen_random_uuid();
ALTER TABLE "user"."roles" ALTER COLUMN "id" TYPE uuid;
