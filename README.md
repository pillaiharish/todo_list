# todo_list
todo list web interface with data being pushed to sql/mongodb

## setting up postgres
```bash
brew install postgresql@14 
harish $ brew services list

Name          Status  User              File
dbus          none                      
dnsmasq       none                      
postgresql@14 started harishkumarpillai ~/Library/LaunchAgents/homebrew.mxcl.postgresql@14.plist
unbound       none   

harish $ sudo -u harishkumarpillai createdb tododb # create database named tododb passing the root user

harish $ psql -d tododb # Access the PostgreSQL database:
psql (14.12 (Homebrew))
Type "help" for help.

tododb=# CREATE USER todouser WITH ENCRYPTED PASSWORD 'password';
CREATE ROLE
tododb=# GRANT ALL PRIVILEGES ON DATABASE tododb TO todouser;
GRANT

```

list all the tables
```
tododb-# \dt
Did not find any relations.
```

List tables in a specific schema:
``` 
tododb-# \dt schema_name.*
Did not find any relation named "schema_name.*".
```

describe a specific table
```
tododb-# \d table_name
Did not find any relation named "table_name".
```
## create table users
```bash
harish $ psql -d tododb
psql (14.12 (Homebrew))
Type "help" for help.

tododb=# CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);
CREATE TABLE
tododb=# \t
Tuples only is on.
tododb=# \d users
 id       | integer                |           | not null | nextval('users_id_seq'::regclass)
 username | character varying(50)  |           | not null | 
 password | character varying(100) |           | not null | 
 email    | character varying(100) |           | not null | 
```

### some postgres commands on brew
While brew services start postgresql@14 starts the PostgreSQL service, there are other commands for service management:



brew services stop postgresql@14: Stops the PostgreSQL service.
brew services restart postgresql@14: Restarts the PostgreSQL service.
brew services info postgresql@14: Displays information about the PostgreSQL service, including its status and logs.
brew services list


## GORM 
GORM is ORM library for Golang https://v1.gorm.io/
