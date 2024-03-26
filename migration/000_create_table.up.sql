CREATE TABLE "results" (
  "id" uuid,
  "position" integer,
  "club_name" varchar,
  "played" integer,
  "won" integer,
  "draw" integer,
  "lost" integer,
  "points" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "matches" (
  "id" uuid PRIMARY KEY,
  "result_id" uuid,
  "match_date" timestamp,
  "attendance" integer,
  "referee" varchar,
  "club_id" uuid,
  "against_club_id" uuid,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "clubs" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE,
  "stadium" varchar,
  "established" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "goals" (
  "id" uuid,
  "match_id" uuid,
  "player_name" varchar,
  "time" timestamp
);

ALTER TABLE "goals" ADD FOREIGN KEY ("match_id") REFERENCES "matches" ("id");

ALTER TABLE "matches" ADD FOREIGN KEY ("result_id") REFERENCES "results" ("id");

ALTER TABLE "matches" ADD FOREIGN KEY ("club_id") REFERENCES "clubs" ("id");

ALTER TABLE "matches" ADD FOREIGN KEY ("against_club_id") REFERENCES "clubs" ("id");