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

insert into users(user_id,username,email,nama_lengkap) 
values( 2,'user2','user2@example.com','user Dua'),( 3,'user3','user3@example.com','user Tiga')

insert into orders (order_id,user_id,tanggal_pemesana ,status) 
values( 1,1,'2023-10-04 20:59:50.815398','Dalam Pengiriman'),
( 2,2,'2023-10-04 20:59:50.815398','Selesai'),
( 3,3,'2023-10-04 21:14:42.578515','Dibatalkan')

insert into order_items (item_id,order_id,product_name,quantity,harga_per_item) 
values( 1,1,'Produk A','2',50000),
( 2,1,'Produk B','3',30000),
( 3,2,'Produk C','1',75000),
( 4,2,'Produk D','2',60000)


select user_id,username from users

select order_id,users.username,sum(harga_per_item)as total_harga from order_items
inner join users on users.user_id = order_items.order_id 
GROUP BY order_items.order_id, users.username;


