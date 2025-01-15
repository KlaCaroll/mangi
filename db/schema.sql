DROP DATABASE IF EXISTS mangi;

CREATE DATABASE mangi;

USE mangi;

create table if not exists user (
	id integer not null AUTO_INCREMENT,
	email varchar(50) not null UNIQUE,
	name text not null,
	password varchar(70) not null,
	is_admin boolean DEFAULT 0,
	primary key (id)
);

create table if not exists home (
	id integer not null AUTO_INCREMENT,
	owner_id integer,
	name varchar(50),
	primary key (id),
	foreign key (owner_id) REFERENCES user(id) ON DELETE CASCADE
);

create table if not exists user_home (
	user_id integer,
	home_id integer,
	primary key (user_id, home_id),
	foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE,
	foreign key (home_id) REFERENCES home(id) ON DELETE CASCADE
);

create table if not exists preference (
	id integer not null AUTO_INCREMENT,
	name varchar(15),
	primary key (id)
);

create table if not exists user_preference (
	user_id integer,
	preference_id integer,
	exist boolean DEFAULT 0,
	primary key (user_id, preference_id),
	foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE,
	foreign key (preference_id) REFERENCES preference(id) ON DELETE CASCADE
);

create table if not exists ustensil (
	id integer not null AUTO_INCREMENT,
	name varchar(15),
	primary key (id)
);

create table if not exists user_ustensil (
	user_id integer,
	ustensil_id integer,
	exist boolean DEFAULT 1,
	primary key (user_id, ustensil_id),
	foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE,
	foreign key (ustensil_id) REFERENCES ustensil(id) ON DELETE CASCADE
);

create table if not exists shopping_list (
	id integer not null AUTO_INCREMENT,
	user_id integer not null,
	food_name varchar(100),
	food_quantity float,
	food_unit varchar(5),
	fromTime datetime,
	toTime datetime,
	name varchar(100),
	home_id integer,
	primary key (id),
	foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE,
	foreign key (home_id) REFERENCES home(id)
);

create table if not exists meal (
	id integer not null AUTO_INCREMENT,
	owner_id integer not null,
	planned_at datetime not null,
	guests integer not null,
	primary key (id),
	foreign key (owner_id) REFERENCES user(id) ON DELETE CASCADE
);

create table if not exists recipe (
	id integer not null AUTO_INCREMENT,
	name text not null,
	owner_id integer,
	preparation_time integer,
	total_time integer,
	description text not null,
	is_public boolean DEFAULT 0,
	primary key (id),
	foreign key (owner_id) REFERENCES user(id) ON DELETE CASCADE
);

create table if not exists ustensil_recipe (
	ustensil_id integer not null,
	recipe_id integer not null,
	primary key (ustensil_id, recipe_id),
	foreign key (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE
);

create table if not exists meal_recipe (
	meal_id integer not null,
	recipe_id integer not null,
	primary key (meal_id, recipe_id),
	foreign key (meal_id) REFERENCES meal(id) ON DELETE CASCADE,
	foreign key (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE
);

create table if not exists user_recipe_favorite (
	user_id integer,
	recipe_id integer,
	primary key (user_id, recipe_id),
	foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE,
	foreign key (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE
);
create table if not exists food (
	id integer not null AUTO_INCREMENT,
	name varchar(100),
	family varchar(50),
	sub_family varchar(30),
	primary key (id)
);

create table if not exists food_preference (
	food_id integer,
	preference_id integer,
	primary key (food_id, preference_id),
	foreign key (food_id) REFERENCES food(id) ON DELETE CASCADE
);

create table if not exists recipe_food (
	recipe_id integer not null,
	food_id integer not null,
	quantity float not null,
	unit enum('u', 'cl', 'g', 'cc', 'cs', 'cm'),
	primary key (recipe_id, food_id)
);

create table if not exists food_month (
	food_id integer,
	month_id integer,
	primary key (food_id, month_id),
	foreign key (food_id) REFERENCES food(id) ON DELETE CASCADE
);

create table if not exists month (
	id integer not null AUTO_INCREMENT,
	name varchar(15),
	season enum('spring', 'summer', 'fall', 'winter'),
	primary key (id)
);

create table if not exists category (
	id integer not null AUTO_INCREMENT,
	name varchar(15),
	exist boolean DEFAULT 1,
	primary key (id)
);

create table if not exists recipe_category (
	recipe_id integer,
	category_id integer,
	primary key (recipe_id, category_id),
	foreign key (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE,
	foreign key (category_id) REFERENCES category(id) ON DELETE CASCADE
);

create table if not exists recipe_comment (
	recipe_id integer,
	comment_id integer,
	primary key (recipe_id, comment_id),
	foreign key (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE
);

create table if not exists comment (
	id integer not null AUTO_INCREMENT,
	owner_id integer not null,
	description text,
	parent_id integer,
	primary key (id)
);