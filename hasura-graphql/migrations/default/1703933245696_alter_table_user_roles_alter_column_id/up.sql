ALTER TABLE "user"."roles" ALTER COLUMN "id" TYPE text;
ALTER TABLE "user"."roles" ALTER COLUMN "id" drop default;
alter table "user"."roles" alter column "id" drop not null;
alter table "user"."roles" rename column "id" to "description";
