CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE products
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    image       text         not null,
    price       int          not null
);

CREATE TABLE categories
(
    id          serial       not null unique,
    title       varchar(255) not null
);

CREATE TABLE categories_products
(
    id      serial                                           not null unique,
    product_id int references products (id) on delete cascade not null,
    category_id int references categories (id) on delete cascade not null
);

CREATE TABLE comments
(
    id serial not null unique,
    title varchar(255) not null,
    body text not null,
    user_id int references users (id) on delete cascade not null
);

CREATE TABLE products_comments
(
    id serial not null unique,
    comment_id int references comments (id) on delete cascade not null,
    product_id int references products (id) on delete cascade not null
);

CREATE TABLE cart
(
    id serial not null unique,
    product_id int references products (id) on delete cascade not null, 
    user_id int not null
);

-- references products (id) on delete cascade