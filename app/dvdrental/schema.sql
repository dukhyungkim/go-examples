create table actor
(
    actor_id    integer   default nextval('actor_actor_id_seq'::regclass) not null
        primary key,
    first_name  varchar(45)                                               not null,
    last_name   varchar(45)                                               not null,
    last_update timestamp default now()                                   not null
);

create table address
(
    address_id  integer   default nextval('address_address_id_seq'::regclass) not null
        primary key,
    address     varchar(50)                                                   not null,
    address2    varchar(50),
    district    varchar(20)                                                   not null,
    city_id     smallint                                                      not null
        constraint fk_address_city
            references city,
    postal_code varchar(10),
    phone       varchar(20)                                                   not null,
    last_update timestamp default now()                                       not null
);

create table category
(
    category_id integer   default nextval('category_category_id_seq'::regclass) not null
        primary key,
    name        varchar(25)                                                     not null,
    last_update timestamp default now()                                         not null
);

create table city
(
    city_id     integer   default nextval('city_city_id_seq'::regclass) not null
        primary key,
    city        varchar(50)                                             not null,
    country_id  smallint                                                not null
        constraint fk_city
            references country,
    last_update timestamp default now()                                 not null
);

create table country
(
    country_id  integer   default nextval('country_country_id_seq'::regclass) not null
        primary key,
    country     varchar(50)                                                   not null,
    last_update timestamp default now()                                       not null
);

create table customer
(
    customer_id integer   default nextval('customer_customer_id_seq'::regclass) not null
        primary key,
    store_id    smallint                                                        not null,
    first_name  varchar(45)                                                     not null,
    last_name   varchar(45)                                                     not null,
    email       varchar(50),
    address_id  smallint                                                        not null
        references address
            on update cascade on delete restrict,
    activebool  boolean   default true                                          not null,
    create_date date      default ('now'::text)::date                           not null,
    last_update timestamp default now(),
    active      integer
);

create table film
(
    film_id          integer       default nextval('film_film_id_seq'::regclass) not null
        primary key,
    title            varchar(255)                                                not null,
    description      text,
    release_year     year,
    language_id      smallint                                                    not null
        references language
            on update cascade on delete restrict,
    rental_duration  smallint      default 3                                     not null,
    rental_rate      numeric(4, 2) default 4.99                                  not null,
    length           smallint,
    replacement_cost numeric(5, 2) default 19.99                                 not null,
    rating           mpaa_rating   default 'G'::mpaa_rating,
    last_update      timestamp     default now()                                 not null,
    special_features text[],
    fulltext         tsvector                                                    not null
);

create table film_actor
(
    actor_id    smallint                not null
        references actor
            on update cascade on delete restrict,
    film_id     smallint                not null
        references film
            on update cascade on delete restrict,
    last_update timestamp default now() not null,
    primary key (actor_id, film_id)
);

create table film_category
(
    film_id     smallint                not null
        references film
            on update cascade on delete restrict,
    category_id smallint                not null
        references category
            on update cascade on delete restrict,
    last_update timestamp default now() not null,
    primary key (film_id, category_id)
);

create table inventory
(
    inventory_id integer   default nextval('inventory_inventory_id_seq'::regclass) not null
        primary key,
    film_id      smallint                                                          not null
        references film
            on update cascade on delete restrict,
    store_id     smallint                                                          not null,
    last_update  timestamp default now()                                           not null
);

create table language
(
    language_id integer   default nextval('language_language_id_seq'::regclass) not null
        primary key,
    name        char(20)                                                        not null,
    last_update timestamp default now()                                         not null
);

create table payment
(
    payment_id   integer default nextval('payment_payment_id_seq'::regclass) not null
        primary key,
    customer_id  smallint                                                    not null
        references customer
            on update cascade on delete restrict,
    staff_id     smallint                                                    not null
        references staff
            on update cascade on delete restrict,
    rental_id    integer                                                     not null
        references rental
            on update cascade on delete set null,
    amount       numeric(5, 2)                                               not null,
    payment_date timestamp                                                   not null
);

create table rental
(
    rental_id    integer   default nextval('rental_rental_id_seq'::regclass) not null
        primary key,
    rental_date  timestamp                                                   not null,
    inventory_id integer                                                     not null
        references inventory
            on update cascade on delete restrict,
    customer_id  smallint                                                    not null
        references customer
            on update cascade on delete restrict,
    return_date  timestamp,
    staff_id     smallint                                                    not null
        constraint rental_staff_id_key
            references staff,
    last_update  timestamp default now()                                     not null
);

create table staff
(
    staff_id    integer   default nextval('staff_staff_id_seq'::regclass) not null
        primary key,
    first_name  varchar(45)                                               not null,
    last_name   varchar(45)                                               not null,
    address_id  smallint                                                  not null
        references address
            on update cascade on delete restrict,
    email       varchar(50),
    store_id    smallint                                                  not null,
    active      boolean   default true                                    not null,
    username    varchar(16)                                               not null,
    password    varchar(40),
    last_update timestamp default now()                                   not null,
    picture     bytea
);

create table store
(
    store_id         integer   default nextval('store_store_id_seq'::regclass) not null
        primary key,
    manager_staff_id smallint                                                  not null
        references staff
            on update cascade on delete restrict,
    address_id       smallint                                                  not null
        references address
            on update cascade on delete restrict,
    last_update      timestamp default now()                                   not null
);
