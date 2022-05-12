create table tables (
	id varchar(100) primary key,
	name text not null,
	capacity int not null,
	pos_x float8 not null,
	pos_y float8 not null,
	notes text default ''
);

create table lines (
	id varchar(100) primary key,
	closed boolean default true,
	coords jsonb not null default '{}'::json
);

create table guests (
	id varchar(100) primary key,
	table_id varchar(100) references tables(id) on delete set null,
	first_name text not null,
	last_name text not null,
	people int not null,
	notes text default ''
);

create table settings (
	id text primary key,
	value jsonb not null default '{}'::json
);
