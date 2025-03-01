CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL DEFAULT '',
		father_name VARCHAR(256) NOT NULL DEFAULT '',
		last_name VARCHAR(256) NOT NULL DEFAULT ''
	);

ALTER SEQUENCE authors_id_seq RESTART WITH 22;

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
	(19, 'Николай', 'Николаевич', 'Носов'),
	(20, 'Эндрю', '', 'Таненбаум'),
	(21, 'Херберт', '', 'Бос');

CREATE TABLE IF NOT EXISTS author_and_literary_work (
		author_id INTEGER,
		literary_work_id INTEGER
	);

INSERT INTO "author_and_literary_work" ("author_id", "literary_work_id") VALUES
	(1, 1),
	(2, 2),
	(3, 3),
	(4, 4),
	(5, 4),
	(2, 5),
	(9, 6),
	(5, 7),
	(10, 8),
	(11, 9),
	(12, 10),
	(12, 11),
	(13, 12),
	(13, 13),
	(14, 14),
	(15, 14),
	(16, 15),
	(2, 16),
	(17, 17),
	(18, 18),
	(18, 19),
	(20, 20),
	(21, 20);

CREATE TABLE IF NOT EXISTS book (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		year_of_publication INTEGER,
		publishing_house_id INTEGER
	);

ALTER SEQUENCE book_id_seq RESTART WITH 20;

INSERT INTO "book" ("id", "user_id", "year_of_publication", "publishing_house_id") VALUES
	(1, 1, 2023, 1),
	(2, 1, 2022, 1),
	(3, 1, 2013, 2),
	(4, 1, 1980, 3),
	(5, 1, 2022, 1),
	(6, 1, 2011, 4),
	(7, 1, 2019, 2),
	(8, 1, 2015, 5),
	(9, 1, 2020, 2),
	(10, 1, 2018, 2),
	(11, 1, 2017, 2),
	(12, 1, 2016, 2),
	(13, 1, 2016, 2),
	(14, 1, 2020, 6),
	(15, 1, 2010, 7),
	(16, 1, 2022, 1),
	(17, 1, 2017, 2),
	(18, 1, 2012, 8),
	(19, 1, 2024, 9);

CREATE TABLE IF NOT EXISTS book_and_literary_work (
		literary_work_id INTEGER,
		book_id INTEGER
	);

INSERT INTO "book_and_literary_work" ("literary_work_id", "book_id") VALUES
	(1, 1),
	(2, 2),
	(3, 3),
	(4, 4),
	(5, 5),
	(6, 6),
	(7, 7),
	(8, 8),
	(9, 9),
	(10, 10),
	(11, 11),
	(12, 12),
	(13, 13),
	(14, 14),
	(15, 15),
	(16, 16),
	(17, 17),
	(18, 18),
	(19, 18),
	(20, 19);

CREATE TABLE IF NOT EXISTS literary_work (
		id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL DEFAULT ''
	);
ALTER SEQUENCE literary_work_id_seq RESTART WITH 21;

INSERT INTO "literary_work" ("id", "name") VALUES
	(1, 'Как закалялась сталь'),
	(2, 'От войны до войны'),
	(3, 'Будущее'),
	(4, 'Манифест коммунистической партии'),
	(5, 'Красное на красном'),
	(6, 'Мастер и Маргарита'),
	(7, 'Происхождение семьи, частной собственности и государства'),
	(8, 'Государь'),
	(9, 'Баллада о змеях и певчих птицах'),
	(10, 'Стрелок'),
	(11, 'Извлечение троих'),
	(12, 'Игра Престолов'),
	(13, 'Битва королей'),
	(14, 'Язык программирования Go'),
	(15, 'Дискретная математика'),
	(16, 'Лик победы'),
	(17, 'Введение в психоанализ'),
	(18, '1984'),
	(19, 'Скотный двор'),
	(20, 'Современные операционные системы');

CREATE TABLE IF NOT EXISTS publishing_house (
		id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL DEFAULT ''
	);
ALTER SEQUENCE publishing_house_id_seq RESTART WITH 10;

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

/* таблица с пользователями */
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	login VARCHAR(256) NOT NULL,
	password VARCHAR(256) NOT NULL
);

ALTER SEQUENCE users_id_seq RESTART WITH 2;
CREATE UNIQUE INDEX idx_login_unique ON users (login);

INSERT INTO "users" ("id", "login", "password") 
VALUES (1, 'Sam', '51b21d529c47d8a88cc39d267fbddd704f19fdb353f5c5b3ca85080c5755715b')
