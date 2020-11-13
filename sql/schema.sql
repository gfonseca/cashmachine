create table cashmachine_schema.accounts
(
	id serial not null,
	balance float8 default 0 not null
);

create unique index accounts_id_uindex
	on cashmachine_schema.accounts (id);

alter table cashmachine_schema.accounts
	add constraint accounts_pk
		primary key (id);

