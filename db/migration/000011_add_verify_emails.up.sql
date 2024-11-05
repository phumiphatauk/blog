CREATE TABLE "verify_emails" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "email" VARCHAR NOT NULL,
  "secret_code" VARCHAR NOT NULL,
  "is_used" BOOL NOT NULL DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "expired_at" TIMESTAMPTZ NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE verify_emails
ADD CONSTRAINT fk_verify_emails_user FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE "users"
ADD COLUMN "is_email_verified" bool NOT NULL DEFAULT false;
