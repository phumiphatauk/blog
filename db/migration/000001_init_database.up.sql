CREATE TABLE permission_group (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL
);

CREATE TABLE permission (
    "id" BIGSERIAL PRIMARY KEY,
    "code" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "permission_group_id" BIGINT NOT NULL
);

ALTER TABLE "permission"
ADD CONSTRAINT fk_permission_permission_group FOREIGN KEY (permission_group_id) REFERENCES "permission_group" (id);

CREATE TABLE role (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NULL,
  "deleted" BOOL NOT NULL DEFAULT false
);

CREATE TABLE role_permission (
    "id" BIGSERIAL PRIMARY KEY,
    "role_id" BIGINT NOT NULL,
    "permission_id" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NULL,
    "deleted" BOOL NOT NULL DEFAULT false
);

ALTER TABLE "role_permission"
ADD CONSTRAINT fk_role_permission_role FOREIGN KEY (role_id) REFERENCES "role" (id);

ALTER TABLE "role_permission"
ADD CONSTRAINT fk_role_permission_permission FOREIGN KEY (permission_id) REFERENCES "permission" (id);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "code" VARCHAR NOT NULL,
  "username" VARCHAR NOT NULL,
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "description" TEXT NULL,
  "hashed_password" VARCHAR NOT NULL,
  "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ,
  "deleted" BOOL NOT NULL DEFAULT false
);

CREATE TABLE "sessions" (
  "id" UUID PRIMARY KEY,
  "user_id" BIGSERIAL NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "user_agent" VARCHAR NOT NULL,
  "client_ip" VARCHAR NOT NULL,
  "is_blocked" BOOLean NOT NULL DEFAULT false,
  "expires_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" 
ADD CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES "users" (id);

CREATE TABLE user_role (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "role_id" BIGINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NULL,
  "deleted" BOOL NOT NULL DEFAULT false
);

ALTER TABLE "user_role"
ADD CONSTRAINT fk_user_role_user FOREIGN KEY (user_id) REFERENCES "users" (id);

ALTER TABLE "user_role"
ADD CONSTRAINT fk_user_role_role FOREIGN KEY (role_id) REFERENCES "role" (id);
