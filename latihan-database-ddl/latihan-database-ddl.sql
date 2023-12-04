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

create table "users" (
	"user_id" SERIAL PRIMARY KEY,
	"username" varchar(30) unique not null,
	"emai" varchar(30) unique not null,
	"nama_lengkap" varchar(30) unique not null);
	
CREATE TABLE orders (
order_id INT PRIMARY KEY,
user_id INT NOT NULL,
tanggal_pemesanan TIMESTAMP DEFAULT NOW(),
status VARCHAR NOT NULL,
CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE order_items (
item_id INT PRIMARY KEY,
order_id INT NOT NULL,
product_name VARCHAR NOT NULL,
quantity INT NOT NULL,
harga_per_item NUMERIC NOT NULL,
CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

select * from users;