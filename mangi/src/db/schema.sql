DROP DATABASE IF EXISTS mangi;
CREATE DATABASE mangi;
USE mangi;

create table if not exists user (
    id integer not null AUTO_INCREMENT,
    email varchar(50) not null,
    password varchar(20) not null,
    name text not null,
    primary key (id)
);

create table if not exists meal (
	id integer not null AUTO_INCREMENT,
	user_id integer not null,
	planned_at datetime not null,
	guests integer not null,
	primary key (id)
);

create table if not exists meal_recipe (
	meal_id integer not null,
	recipe_id integer not null,
	primary key (meal_id, recipe_id)
);

create table if not exists recipe (
	id integer not null AUTO_INCREMENT,
	name text not null,
	primary key (id)
);

create table if not exists recipe_food (
	recipe_id integer not null,
	food_id integer not null,
	quantity float not null,
	primary key (recipe_id, food_id)
);

create table if not exists food (
	id integer not null AUTO_INCREMENT,
	name text not null,
	primary key (id)
);


