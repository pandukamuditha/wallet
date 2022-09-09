create table "user" (
	"id" bigserial,
	"fname" varchar(50) not null,
	"lname" varchar(50) not null,
	primary key ("id")
)

create table "wallet" (
	"id" bigserial,
	"user_id" bigserial not null,
	"balance" decimal(10,2) default 0 not null,
	primary key ("id"),
	foreign key ("user_id") references "user"("id")
	on update cascade
	on delete cascade
)

create table "transaction" (
	"id" bigserial,
	"from_wallet" bigserial not null,
	"to_wallet" bigserial not null,
	"amount" decimal(10,2) default 0 not null,
	"created_date" timestamp with time zone default now() not null,
	"updated_date" timestamp with time zone default now() not null,
	primary key ("id"),
	foreign key ("from_wallet") references "wallet"("id") on update cascade on delete cascade,
	foreign key ("to_wallet") references "wallet"("id") on update cascade on delete cascade
)

CREATE INDEX user_fname_idx ON "user" ("fname");
CREATE INDEX user_lname_idx ON "user" ("lname");

CREATE INDEX wallet_user_id_idx ON "wallet" ("user_id");

CREATE INDEX transaction_from_wallet_idx ON "transaction" ("from_wallet");
CREATE INDEX transaction_to_wallet_idx ON "transaction" ("to_wallet");
CREATE INDEX transaction_from_to_wallet_idx ON "transaction" ("from_wallet", "to_wallet");
CREATE INDEX transaction_created_date_idx ON "transaction" ("created_date");
CREATE INDEX transaction_updated_date_idx ON "transaction" ("updated_date");
