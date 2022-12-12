TRUNCATE TABLE user_roles RESTART IDENTITY;

INSERT INTO user_roles(name) values('admin');
INSERT INTO user_roles(name) values('user');
