USE mangi;

insert into user (email, password, name) values 
('foo@foo.com', 'foo', 'foo'),
('bar@bar.com', 'bar', 'bar'),
('fizz@fizz.com', 'fizz', 'fizz'),
('buzz@buzz.com', 'buzz', 'buzz');

insert into food (name) values
('oeuf'),
('lait'),
('beurre'),
('chocolat');

insert into meal (planned_at, guests, user_id) values
('2022-07-01 19:00:00', 2, 1),
('2022-07-02 19:00:00', 2, 1),
('2022-07-01 21:00:00', 4, 2),
('2022-07-02 21:00:00', 4, 2);

insert into recipe (name) values
('gateau'),
('oeuf au plat');

insert into recipe_food (recipe_id, food_id, quantity) values
(1, 1, 1),
(1, 2, 200),
(1, 3, 200),
(1, 4, 200);
