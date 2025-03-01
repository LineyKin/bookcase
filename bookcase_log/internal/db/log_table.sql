CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		user_id INT,
		producer_ts TIMESTAMP,
		consumer_ts TIMESTAMP,
		topic VARCHAR(256) NOT NULL DEFAULT '',
		message TEXT NOT NULL DEFAULT ''
	);