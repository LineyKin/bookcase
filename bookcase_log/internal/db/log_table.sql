CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		producer_ts TIMESTAMP,
		consumer_ts TIMESTAMP,
		message_id INTEGER,
		topic VARCHAR(256) NOT NULL DEFAULT '',
		message TEXT NOT NULL DEFAULT ''
	);