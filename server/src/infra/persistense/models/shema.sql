create table if not exists users (
    id uuid primary key,
    username text,
    permissions text[],
    password text
);

create table if not exists planes (
    id uuid primary key,
    name text
);

create table if not exists seats (
    id uuid primary key,
    plane_id uuid references planes (id),
    tag text,
    serial integer,
    class text
);

create table if not exists flights (
    id uuid primary key,
    departure_aip_code text,
    arrival_aip_code text,
    plane_id uuid references planes (id),
    arrival_time timestamp,
    departure_time timestamp
);

create type price as (
    amount bigint,
    currency text
);

create table if not exists flight_seats (
    id uuid primary key,
    is_occupied boolean,
    price price,
    seat_id uuid references seats (id),
    flight_id uuid references flights (id)
);

create table if not exists tickets (
    id uuid primary key,
    buyer_id uuid references users (id),
    price price,
    flight_seat_id uuid references flight_seats (id)
);