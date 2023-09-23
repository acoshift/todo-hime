create table todos (
    id uuid default gen_random_uuid(),
    content varchar not null,
    done boolean not null default false,
    created_at timestamp not null default now(),
    primary key (id)
);
