create table movie (
    id uuid primary key,
    title varchar not null,
    duration time not null,
    description text
);