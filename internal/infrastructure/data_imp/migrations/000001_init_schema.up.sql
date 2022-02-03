CREATE TABLE "web_sites" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "url" varchar NOT NULL,
    "status" integer NOT NULL,
    "last_checked" DATE,
    "last_updated" DATE
);