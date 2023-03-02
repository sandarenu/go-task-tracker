# Simple application to learn GO

## Functionality

Simple CLI based task tracking application

* List tasks
* Add new task
* Remove task
* Complete task

### Task

* Task Id
* Title
* Status
* Created date
* Due date

## Database

```sql
DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks
(
    id     INT AUTO_INCREMENT NOT NULL,
    title  VARCHAR(128) NOT NULL,
    status INT          NOT NULL,
    PRIMARY KEY (`id`)
);
```

```sql
insert into tasks(title, status)
values ("Integrate with mysql db", 1),
       ("Fetch tasks from DB and show", 1),
       ("Mark task as completed", 1);
```

## Reference

* GO DB access - https://go.dev/doc/tutorial/database-access
* GO SQL Driver - https://github.com/go-sql-driver/mysql/