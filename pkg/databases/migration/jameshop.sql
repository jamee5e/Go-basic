CREATE TABLE "users" (
  "id" verchar PRIMARY KEY,
  "username" verchar UNIQUE,
  "password" verchar,
  "email" verchar UNIQUE,
  "role_id" int,
  "creted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "oauth" (
  "id" verchar PRIMARY KEY,
  "user_id" verchar,
  "access_token" varchar,
  "refresh_token" varchar,
  "creted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "products" (
  "id" varchar PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "price" float,
  "creted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "images" (
  "id" varchar PRIMARY KEY,
  "filename" varchar,
  "url" varchar,
  "product_id" varchar,
  "creted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" varchar PRIMARY KEY,
  "product_id" varchar,
  "category_id" int
);

CREATE TABLE "categories" (
  "id" int PRIMARY KEY,
  "title" varchar UNIQUE
);

CREATE TABLE "oders" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "contact" varchar,
  "address" varchar,
  "tranfer_slip" jsonb,
  "status" varchar,
  "creted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "products_oders" (
  "id" varchar PRIMARY KEY,
  "oders_id" varchar,
  "qty" int,
  "product" jsonb
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "oders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_oders" ADD FOREIGN KEY ("oders_id") REFERENCES "oders" ("id");
