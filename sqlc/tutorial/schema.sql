create table actor
(
    actor_id    integer   default nextval('actor_actor_id_seq'::regclass) not null
        primary key,
    first_name  varchar(45)                                               not null,
    last_name   varchar(45)                                               not null,
    last_update timestamp default now()                                   not null
);