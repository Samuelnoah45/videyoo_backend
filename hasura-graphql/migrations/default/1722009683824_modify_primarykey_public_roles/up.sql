BEGIN TRANSACTION;
ALTER TABLE "public"."roles" DROP CONSTRAINT "roles_pkey";

ALTER TABLE "public"."roles"
    ADD CONSTRAINT "roles_pkey" PRIMARY KEY ("id", "role_name");
COMMIT TRANSACTION;
