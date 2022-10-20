CREATE TABLE "user" (
	"id" bigserial PRIMARY KEY,
	"fname" varchar NOT NULL,
	"lname" varchar NOT NULL
);

CREATE TABLE "wallet" (
	"id" bigserial PRIMARY KEY,
	"user_id" bigserial NOT NULL,
	"balance" integer default 0 NOT NULL
);

ALTER TABLE "wallet" add foreign key ("user_id") references "user" ("id");

CREATE TABLE "transaction" (
	"id" bigserial PRIMARY KEY,
	"from_wallet" bigserial NOT NULL,
	"to_wallet" bigserial NOT NULL,
	"amount" integer default 0 NOT NULL,
	"created_date" timestamp with time zone default now() NOT NULL,
	"updated_date" timestamp with time zone default now() NOT NULL
);

ALTER TABLE "transaction" add foreign key ("from_wallet") references "wallet" ("id");
ALTER TABLE "transaction" add foreign key ("to_wallet") references "wallet" ("id");

CREATE INDEX user_fname_idx ON "user" ("fname");
CREATE INDEX user_lname_idx ON "user" ("lname");

CREATE INDEX wallet_user_id_idx ON "wallet" ("user_id");

CREATE INDEX transaction_from_wallet_idx ON "transaction" ("from_wallet");
CREATE INDEX transaction_to_wallet_idx ON "transaction" ("to_wallet");
CREATE INDEX transaction_from_to_wallet_idx ON "transaction" ("from_wallet", "to_wallet");
CREATE INDEX transaction_created_date_idx ON "transaction" ("created_date");
CREATE INDEX transaction_updated_date_idx ON "transaction" ("updated_date");

CREATE TABLE "auth" (
	"user_id" bigserial UNIQUE NOT NULL,
	"username" varchar PRIMARY KEY,
	"password" varchar NOT NULL,
	"claims" varchar
);

ALTER TABLE "auth" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
