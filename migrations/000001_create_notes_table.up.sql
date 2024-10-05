CREATE TABLE IF NOT EXISTS notes (
	note_id bigserial PRIMARY KEY,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	note text NOT NULL,
	version integer NOT NULL DEFAULT 1
);
