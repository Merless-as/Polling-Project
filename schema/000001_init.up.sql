CREATE TABLE polls
(
    id serial not null unique,
    name varchar(256) not null
);

CREATE TABLE choices
(
    poll_id int references polls (id) on delete cascade not null unique,
    choice_id int not null,
    choice varchar(256),
    call int not null
);