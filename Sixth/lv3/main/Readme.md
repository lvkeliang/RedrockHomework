## 数据库结构

## 数据库名：redrocksixth

| Tables_in_redrocksixth |
| --- |
| active_users |
| messages |
| users |

## 表的结构

### 表名：active_users

| Field | Type | Null | Key | Default | Extra |
| --- | --- | --- | --- | --- | --- |
| id | bigint | NO | PRI | NULL | auto_increment |
| cookie | varchar(20) | YES  |  |  |  |

### 表名：messages

| Field   | Type         | Null | Key | Default | Extra          |
| --- | --- | --- | --- | --- | --- |
| id | bigint | NO | PRI | NULL | auto_increment |
| name | varchar(20)  | YES  |  |  |  |
| message | varchar(100) | YES  |  | NULL |  |
| time    | datetime     | YES  |  | NULL |  |

### 表名：users

| Field | Type | Null | Key | Default | Extra |
| --- | --- | --- | --- | --- | --- |
| id | bigint | NO | PRI | NULL | auto_increment |
| name | varchar(20) | YES |  |  |  |
| password | varchar(20) | YES |  |  |  |
| question1 | varchar(20) | YES |  |  |  |
| answer1 | varchar(20) | YES |  |  |  |
