-- Database: toko_online_noobee

-- DROP DATABASE IF EXISTS toko_online_noobee;

CREATE DATABASE toko_online_noobee
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

select * from users;
	
INSERT INTO users (username, emai, nama_lengkap)
VALUES
('user1', 'user1@example.com', 'User Satu'),
('user2', 'user2@example.com', 'User Dua'),
('user3', 'user3@example.com', 'User Tiga');

select * from orders;

INSERT INTO orders (order_id,user_id, status)
VALUES
(1, 1, 'Dalam Pengiriman'),
(2, 2, 'Selesai'),
(3, 3, 'Dibatalkan');

select * from order_items;

INSERT INTO order_items (item_id,order_id, product_name, quantity, harga_per_item)
VALUES
(1,1, 'Produk A', 2, 50000),
(2,1, 'Produk B', 3, 30000),
(3,2, 'Produk C', 1, 750000),
(4,2, 'Produk D', 2, 60000);

select orders.order_id, users.username
from orders
join users on orders.user_id = users.user_id;

SELECT orders.order_id, users.username, SUM(order_items.quantity * order_items.harga_per_item) AS total_harga
FROM orders
JOIN users ON orders.user_id = users.user_id
JOIN order_items ON orders.order_id = order_items.order_id
GROUP BY orders.order_id, users.username;


