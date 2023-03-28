-- Project Name : dab-database
-- Date/Time    : 2023/03/28 23:17:25
-- Author       : Yuta Ono
-- RDBMS Type   : MySQL
-- Application  : A5:SQL Mk-2

/*
  << 注意！！ >>
  BackupToTempTable, RestoreFromTempTable疑似命令が付加されています。
  これにより、drop table, create table 後もデータが残ります。
  この機能は一時的に $$TableName のような一時テーブルを作成します。
  この機能は A5:SQL Mk-2でのみ有効であることに注意してください。
*/

-- shops
--* BackupToTempTable
drop table if exists shops cascade;

--* RestoreFromTempTable
create table shops (
  shop_id INT not null auto_increment comment 'shop_id'
  , shop_name VARCHAR(100) not null comment 'shop_name'
  , modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at'
  , created_at TIMESTAMP default current_timestamp not null comment 'created_at'
  , constraint shops_PKC primary key (shop_id)
) comment 'shops' ;

-- currencies
--* BackupToTempTable
drop table if exists currencies cascade;

--* RestoreFromTempTable
create table currencies (
  currency_id INT not null auto_increment comment 'currency_id'
  , currency_name CHAR(3) not null comment 'currency_name'
  , in_yen FLOAT not null comment 'in_yen'
  , modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at'
  , created_at TIMESTAMP default current_timestamp not null comment 'created_at'
  , constraint currencies_PKC primary key (currency_id)
) comment 'currencies' ;

-- products
--* BackupToTempTable
drop table if exists products cascade;

--* RestoreFromTempTable
create table products (
  receipt_id INT not null comment 'receipt_id'
  , product_id INT not null comment 'product_id'
  , product_name VARCHAR(100) not null comment 'product_name'
  , price FLOAT not null comment 'price'
  , constraint products_PKC primary key (receipt_id,product_id)
) comment 'products' ;

-- receipts
--* BackupToTempTable
drop table if exists receipts cascade;

--* RestoreFromTempTable
create table receipts (
  receipt_id INT not null auto_increment comment 'receipt_id'
  , shop_id INT not null comment 'shop_id'
  , currency_id INT not null comment 'currency_id'
  , total_price FLOAT not null comment 'total_price'
  , purchase_date DATE not null comment 'purchase_date'
  , modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at'
  , created_at TIMESTAMP default current_timestamp not null comment 'created_at'
  , constraint receipts_PKC primary key (receipt_id)
) comment 'receipts' ;

create index receipts_IX1
  on receipts(total_price);

create index receipts_IX2
  on receipts(purchase_date);

alter table receipts
  add constraint receipts_FK1 foreign key (currency_id) references currencies(currency_id);

alter table receipts
  add constraint receipts_FK2 foreign key (shop_id) references shops(shop_id);

alter table products
  add constraint products_FK1 foreign key (receipt_id) references receipts(receipt_id);

