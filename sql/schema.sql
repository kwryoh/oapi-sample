---
create table items (
    id serial primary key not null,
    name varchar(256) not null,
    code varchar(256) not null,
    unit varchar(256) not null,
    cost DECIMAL(13, 4) not null,
    created timestamp not null default now(),
    updated timestamp not null default now()
);