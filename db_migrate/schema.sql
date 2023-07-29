CREATE EXTENSION "uuid-ossp";

CREATE TABLE "users"
(
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(256) NOT NULL,
    "email" VARCHAR(256) NOT NULL UNIQUE,
    "password_hash" TEXT,
    "image" TEXT,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY ("id")
);