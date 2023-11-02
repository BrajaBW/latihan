create database toko_online_noobe

create table users(
user_id int primary key,
username varchar (50) unique not null,
email varchar (100) unique not null,
nama_lengkap varchar (100) unique not null
);

create table orders(
order_id int primary key,
user_id int,
constraint orders_user_id_fkey foreign key (user_id) references users(user_id),
tanggal_pemesana timestamp,
status varchar not null
);

create table order_items(
item_id int primary key,
order_id int,
constraint order_items_order_id_fkey foreign key (order_id) references orders(order_id),
product_name varchar not null,
quantity int not null,
harga_per_item numeric not null
);