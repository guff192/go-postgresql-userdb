-- auto-generated definition
create table users
(
    id        integer,
    name      varchar(80),
    lastname  varchar(80),
    age       integer,
    birthdate date
);

alter table users
    owner to admin;

