CREATE TABLE categories (
	name 			varchar(50) PRIMARY KEY
);

COPY categories (name) FROM '/etc/postgres/data/categories.txt';

INSERT INTO categories (name) VALUES ('other');
CREATE RULE non_upd_other AS ON UPDATE TO categories WHERE old.name = 'other' DO INSTEAD NOTHING;
CREATE RULE non_del_other AS ON DELETE TO categories WHERE old.name = 'other' DO INSTEAD NOTHING;

CREATE TABLE types (
	name			varchar(50) PRIMARY KEY
);

COPY types (name) FROM '/etc/postgres/data/types.txt';

CREATE TABLE transactions (
	id 				uuid PRIMARY KEY,
	user_id 		uuid NOT NULL,

	type 			varchar(50) NOT NULL,
	target 			varchar(50) DEFAULT '',
	description 	text DEFAULT '',
	category 		varchar(50) NOT NULL DEFAULT 'other',
	cost 			float8 NOT NULL CHECK (cost > 0),

	timestamp 		timestamp DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (type) REFERENCES types (name) ON DELETE RESTRICT ON UPDATE CASCADE,
	FOREIGN KEY (category) REFERENCES categories (name) ON DELETE SET DEFAULT ON UPDATE CASCADE
);
