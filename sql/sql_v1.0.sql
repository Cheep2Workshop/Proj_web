UPDATE TABLE users
(id int,
name varchar(50),
email varchar(50),
password varchar(50),
admin bool,
created_at datetime);

CREATE TABLE dashboard_login_log
(id int,
user_id int,
created_at datetime);