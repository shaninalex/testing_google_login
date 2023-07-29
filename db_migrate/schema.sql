CREATE EXTENSION "uuid-ossp";

CREATE TABLE "providers" (
    "id" SERIAL,
    "name" VARCHAR(50) NOT NULL,
    PRIMARY KEY ("id")
);

-- "local" - default provider for user created in the system by email/password
INSERT INTO
    "providers" ("name")
VALUES
    ('local'),
    ('google'),
    ('facebook'),
    ('twitter');

CREATE TABLE "users" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(256) NOT NULL,
    "email" VARCHAR(256) NOT NULL UNIQUE,
    "password_hash" TEXT,
    "image" TEXT,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY ("id")
);

CREATE TABLE "user_providers" (
    "id" SERIAL,
    "user_id" uuid NOT NULL,
    "provider_id" INTEGER NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("provider_id") REFERENCES "providers" ("id")
);