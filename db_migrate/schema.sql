CREATE EXTENSION "uuid-ossp";

CREATE TABLE "providers" (
    "id" SERIAL,
    "name" VARCHAR(10) NOT NULL,
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



-- create user.
-- Arguments:
--      - p_name - user name
--      - p_email - user email
--      - p_image - user image ( avatar or user_pic, depends of providers)
--      - p_provider - string, provider name, one of the providers in "providers" table 
CREATE OR REPLACE PROCEDURE create_user(
    p_name VARCHAR(256),
    p_email VARCHAR(256),
    p_image TEXT
) LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO "users" ("name", "email", "image")
    VALUES (p_name, p_email, p_image);
    -- TODO: insert into user_providers
    COMMIT;
END;$$