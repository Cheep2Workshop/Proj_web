# proj-web

## Note

### migrate

```sh
migrate -source file://./ -database "mysql://root:QMKAJNNjNK9vBO88@tcp(localhost:3306)/dashboard" up 1
```

## Describe

## use

- golang
- gin
- gorm
- mysql
- redis
- grpc
- docker

### Step 1

- 登入畫面
    - 使用者登入
        - 顯示本身資訊 (Name,Email) 
        - 登出
    - 管理者登入
        - 顯示本身資訊 (Name,Email) 
        - 檢視後台登入紀錄 (Name,Email,Time) 
        - 增刪改查使用者資訊 (Name,Email) All
        - 登出

### Step 2
- 使用者登入
    - 自身的登入紀錄

## Arch Page

### Step 1
- login
    - homepage
    - myinfo
    - logout
    - (admin) alluser
    - (admin) dashboard.log

## API Features

| Functions              | Detail                                            | URL                         |
| :--------------------: | ------------------------------------------------- | --------------------------- |
| Sign up | User can sign up an account by inputting name, email, password | /signup |
| Log in | User can log in using registered email | /signin |
| Log out | User can log out of an account | /users/logout |
| View all users | Admin can view all users after log in | /admin/users |
| Edit a user | Admin can update user's role after log in | /admin/users/:id |

## SQL desgin

### Web

```sql=mysql
-- table 使用者 
CREATE TABLE user
(id int,
name varchar(50),
account varchar(50),
password varchar(50),
admin bool,
created_at datetime);

-- table 後台登入紀錄
CREATE TABLE dashboard_login_log
(id int,
name varchar(50),
email varchar(50),
created_at datetime);
```

### Game

```sql=mysql
-- table 玩家資料
CREATE TABLE app_user
(id int,
name varchar(50),
email varchar(50),
password varchar(50),
created_at datetime);

-- table 遊戲登入紀錄
CREATE TABLE app_login_log
(app_id int, 
user_id int, 
created_at datetime);

-- table 遊戲版本紀錄
CREATE TABLE app_version_info
(app_id int,
app_name char(50), 
app_version char(20), 
created_at datetime, 
updated_at datetime);

-- table 遊戲系統紀錄
CREATE TABLE app_system_log
(log_id int, 
app_id int, 
log_tags varchar(50), 
log_msg char(256), 
created_at datetime);
```

## how to run

### run via docker

```shell
# docker run mysql
make mysql-up 
```
