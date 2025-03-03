-- ТАБЛИЦА АВТОРОВ
CREATE TABLE IF NOT EXISTS authors (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL DEFAULT '',
	father_name VARCHAR(256) NOT NULL DEFAULT '',
	last_name VARCHAR(256) NOT NULL DEFAULT ''
);

COMMENT ON TABLE authors IS 'таблица авторов: писателей, поэтов, публицистов';

ALTER SEQUENCE authors_id_seq RESTART WITH 24;

INSERT INTO "authors" ("id", "name", "father_name", "last_name") VALUES
(1,  'Николай', 'Алексеевич','Островский'),
(2,  'Вера', 'Викторовна', 'Камша'),
(3,  'Дмитрий', 'Алексеевич','Глуховский'),
(4,  'Карл', '', 'Маркс'),
(5,  'Фридрих', '', 'Энгельс'),
(9,  'Михаил', 'Афанасьевич', 'Булгаков'),
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
(20, 'Херберт', '', 'Бос'),
(21, 'Агата', '', 'Кристи'),
(22, 'Адитья', '', 'Бхаргава'),
(23, 'Олдос', '', 'Хаксли');


-- ТАБЛИЦА ХУДОЖЕСТВЕННЫХ ПРОИЗВЕДЕНИЙ
CREATE TABLE IF NOT EXISTS literary_work (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL DEFAULT '',
	authors INTEGER[]
);
ALTER SEQUENCE literary_work_id_seq RESTART WITH 27;

INSERT INTO literary_work (id, name, authors) VALUES
(1,   'Как закалялась сталь','{1}'),
(2,   'Язык программирования Go','{14,15}'),
(3,   'Красное на красном','{2}'),
(4,   'От войны до войны','{2}'),
(5,   'Манифест коммунистической партии','{4,5}'),
(6,   'Современные операционные системы','{19,20}'),
(7,   'Вилла Белый конь','{21}'),
(8,   'Н или М' ,'{21}'),
(9,   'Чаепитие в Хантербери','{21}'),
(10,  'Убить легко' ,'{21}'),
(11,  'Грокаем алгоритмы','{22}'),
(12,  'Дискретная математика','{16}'),
(13,  'Будущее','{3}'),
(14,  'Государь','{10}'),
(15,  'Баллада о змеях и певчих птицах','{11}'),
(16,  '1984','{18}'),
(17,  'Скотный двор','{18}'),
(18,  'Происхождение семьи, частной собственности и государства','{5}'),
(19,  'Введение в психоанализ','{17}'),
(20,  'О дивный новый мир','{23}'),
(21,  'Мастер и Маргарита','{9}'),
(22,  'Игра престолов','{13}'),
(23,  'Битва королей','{13}'),
(24,  'Стрелок','{12}'),
(25,  'Извлечение троих','{12}'),
(26,  'Лик победы','{2}');

-- индекс для более быстрых join-ов с таблицей literary_work по полю authors
CREATE INDEX IF NOT EXISTS idx_lw_authors ON literary_work USING GIN (authors);

COMMENT ON TABLE literary_work IS 'таблица художественных произведений (без привязки к физической книге)';
COMMENT ON COLUMN literary_work.authors IS 'список айдишек авторов';
COMMENT ON INDEX idx_lw_authors IS 'индекс для более быстрых join-ов с таблицей literary_work по полю authors';


-- ТАБЛИЦА ФИЗИЧЕСКИХ КНИГ
CREATE TABLE IF NOT EXISTS book (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	year_of_publication INTEGER,
	publishing_house_id INTEGER,
	literary_works INTEGER[]
);
ALTER SEQUENCE book_id_seq RESTART WITH 23;

INSERT INTO book (id, user_id, year_of_publication, publishing_house_id, literary_works) VALUES
(1,   1,   2023,    1,   '{1}'),
(2,   1,   2020,    2,   '{2}'),
(3,   1,   2022,    1,   '{3}'),
(4,   1,   2022,    1,   '{4}'),
(5,   1,   1980,    3,   '{5}'),
(6,   1,   2024,    4,   '{6}'),
(7,   1,   1991,    5,   '{7,8,9,10}'),
(8,   1,   2024,    4,   '{11}'),
(9,   1,   2010,    6,   '{12}'),
(10,  1,   2013,    7,   '{13}'),
(11,  1,   2015,    8,   '{14}'),
(12,  1,   2020,    7,   '{15}'),
(13,  1,   2012,    9,   '{16,17}'),
(14,  1,   2019,    7,   '{18}'),
(15,  1,   2017,    7,   '{19}'),
(16,  1,   2014,    7,   '{20}'),
(17,  1,   2011,    10,  '{21}'),
(18,  1,   2016,    7,   '{22}'),
(19,  1,   2016,    7,   '{23}'),
(20,  1,   2018,    7,   '{24}'),
(21,  1,   2017,    7,   '{25}'),
(22,  1,   2022,    1,   '{26}');

-- индекс для более быстрых join-ов с таблицей book по полю literary_works
CREATE INDEX IF NOT EXISTS idx_book_lw ON book USING GIN (literary_works);

COMMENT ON TABLE book IS 'таблица физических книг';
COMMENT ON COLUMN book.literary_works IS 'список айдишек художественных произведений и работ';
COMMENT ON INDEX idx_book_lw IS 'индекс для более быстрых join-ов с таблицей book по полю literary_works';


-- ТАБЛИЦА ИЗДАТЕЛЬСТВ
CREATE TABLE IF NOT EXISTS publishing_house (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL DEFAULT ''
);
ALTER SEQUENCE publishing_house_id_seq RESTART WITH 11;


INSERT INTO publishing_house (id, name) VALUES
(1,   'Эксмо'),
(2,   'Диалектика'),
(3,   'Политиздат'),
(4,   'Питер'),
(5,   'Корона-принт'),
(6,   'Научный мир'),
(7,   'АСТ'),
(8,   'РИПОЛ классик'),
(9,   'Астрель'),
(10,  'Мартин');

COMMENT ON TABLE publishing_house IS 'таблица издательств';


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