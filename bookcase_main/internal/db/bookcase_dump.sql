-- ТАБЛИЦА АВТОРОВ
CREATE TABLE IF NOT EXISTS authors (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL DEFAULT '',
	father_name VARCHAR(256) NOT NULL DEFAULT '',
	last_name VARCHAR(256) NOT NULL DEFAULT ''
);

COMMENT ON TABLE authors IS 'таблица авторов: писателей, поэтов, публицистов'

ALTER SEQUENCE authors_id_seq RESTART WITH 21;

INSERT INTO "authors" ("id", "name", "father_name", "last_name") VALUES
	(1, 'Николай', 'Алексеевич', 'Островский'),
	(2, 'Вера', 'Викторовна', 'Камша'),
	(3, 'Дмитрий', 'Алексеевич', 'Глуховский'),
	(4, 'Карл', '', 'Маркс'),
	(5, 'Фридрих', '', 'Энгельс'),
	(9, 'Михаил', 'Афанасьевич', 'Булгаков'),
	(10, 'Никколо', '', 'Макиавелли'),
	(11, 'Сьюзен', '', 'Коллинз'),
	(12, 'Стивен', '', 'Кинг'),
	(13, 'Джордж', '', 'Мартин'),
	(14, 'Алан', 'А.А.', 'Донован'),
	(15, 'Брайан', 'У.', 'Керниган'),
	(16, 'Алексей', 'Александрович', 'Набебин'),
	(17, 'Зигмунд', '', 'Фрейд'),
	(18, 'Джордж', '', 'Оруэлл'),
	(19, 'Эндрю', '', 'Таненбаум'),
	(20, 'Херберт', '', 'Бос');


-- ТАБЛИЦА ХУДОЖЕСТВЕННЫХ ПРОИЗВЕДЕНИЙ
CREATE TABLE IF NOT EXISTS literary_work (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL DEFAULT '',
	authors INTEGER[]
);

COMMENT ON TABLE literary_work IS 'таблица художественных произведений (без привязки к физической книге)';
COMMENT ON COLUMN literary_work.authors IS 'список айдишек авторов';


-- ТАБЛИЦА ФИЗИЧЕСКИХ КНИГ
CREATE TABLE IF NOT EXISTS book (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	year_of_publication INTEGER,
	publishing_house_id INTEGER,
	literary_works INTEGER[]
);

COMMENT ON TABLE book IS 'таблица физических книг';
COMMENT ON COLUMN book.literary_works IS 'список айдишек художественных произведений и работ';


-- ТАБЛИЦА ИЗДАТЕЛЬСТВ
CREATE TABLE IF NOT EXISTS publishing_house (
		id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL DEFAULT ''
	);
ALTER SEQUENCE publishing_house_id_seq RESTART WITH 10;

COMMENT ON TABLE publishing_house IS 'таблица издательств';
INSERT INTO "publishing_house" ("id", "name") VALUES
	(1, 'Эксмо'),
	(2, 'АСТ'),
	(3, 'Политиздат'),
	(4, 'Мартин'),
	(5, 'РИПОЛ классик'),
	(6, 'Диалектика'),
	(7, 'Научный мир'),
	(8, 'Астрель'),
	(9, 'Питер');


-- ТАБЛИЦА С ПОЛЬЗОВАТЕЛЯМИ
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	login VARCHAR(256) NOT NULL,
	password CHAR(64) NOT NULL
);

COMMENT ON TABLE users IS 'таблица пользователей приложения';

ALTER SEQUENCE users_id_seq RESTART WITH 2;
CREATE UNIQUE INDEX idx_login_unique ON users (login);

INSERT INTO "users" ("id", "login", "password") 
VALUES (1, 'Sam', '51b21d529c47d8a88cc39d267fbddd704f19fdb353f5c5b3ca85080c5755715b');

