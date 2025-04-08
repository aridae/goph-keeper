package database

const schemaDDL = `
create table if not exists users (
    id serial primary key,
    username text unique not null,
    password_hash bytea not null,
    created_at timestamp not null,
    is_deleted boolean not null default false
);

create table if not exists secrets (
    id serial primary key,
    key text not null,
    data bytea not null,
    owner_username text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    is_deleted boolean not null default false 
);

create unique index if not exists idx_unq_secrets__owner_username__key
    on secrets(owner_username, key)
    where is_deleted=false;
`

const (
	UsersTable   = "users"
	SecretsTable = "secrets"
)
