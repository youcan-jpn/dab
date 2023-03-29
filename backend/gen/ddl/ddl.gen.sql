-- Project Name : dab-database
-- Date/Time    : 2023/03/29 21:44:36
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

-- categories
--* RestoreFromTempTable
CREATE TABLE categories (
  category_id INT NOT NULL auto_increment COMMENT 'category_id'
  , category_name VARCHAR(15) NOT NULL COMMENT 'category_name'
  , CONSTRAINT categories_PKC PRIMARY KEY (category_id)
) COMMENT 'categories' ;

-- shops
--* RestoreFromTempTable
CREATE TABLE shops (
  shop_id INT NOT NULL auto_increment COMMENT 'shop_id'
  , shop_name VARCHAR(100) NOT NULL COMMENT 'shop_name'
  , modified_at TIMESTAMP DEFAULT current_timestamp on update current_timestamp NOT NULL COMMENT 'modified_at'
  , created_at TIMESTAMP DEFAULT current_timestamp NOT NULL COMMENT 'created_at'
  , CONSTRAINT shops_PKC PRIMARY KEY (shop_id)
) COMMENT 'shops' ;

-- currencies
--* RestoreFromTempTable
CREATE TABLE currencies (
  currency_id INT NOT NULL auto_increment COMMENT 'currency_id'
  , currency_name CHAR(3) NOT NULL COMMENT 'currency_name'
  , in_yen FLOAT NOT NULL COMMENT 'in_yen'
  , modified_at TIMESTAMP DEFAULT current_timestamp on update current_timestamp NOT NULL COMMENT 'modified_at'
  , created_at TIMESTAMP DEFAULT current_timestamp NOT NULL COMMENT 'created_at'
  , CONSTRAINT currencies_PKC PRIMARY KEY (currency_id)
) COMMENT 'currencies' ;

-- products
--* RestoreFromTempTable
CREATE TABLE products (
  receipt_id INT NOT NULL COMMENT 'receipt_id'
  , product_id INT NOT NULL COMMENT 'product_id'
  , product_name VARCHAR(100) NOT NULL COMMENT 'product_name'
  , price FLOAT NOT NULL COMMENT 'price'
  , CONSTRAINT products_PKC PRIMARY KEY (receipt_id,product_id)
) COMMENT 'products' ;

-- receipts
--* RestoreFromTempTable
CREATE TABLE receipts (
  receipt_id INT NOT NULL auto_increment COMMENT 'receipt_id'
  , shop_id INT NOT NULL COMMENT 'shop_id'
  , currency_id INT NOT NULL COMMENT 'currency_id'
  , category_id INT NOT NULL COMMENT 'category_id'
  , total_price FLOAT NOT NULL COMMENT 'total_price'
  , purchase_date DATE NOT NULL COMMENT 'purchase_date'
  , modified_at TIMESTAMP DEFAULT current_timestamp on update current_timestamp NOT NULL COMMENT 'modified_at'
  , created_at TIMESTAMP DEFAULT current_timestamp NOT NULL COMMENT 'created_at'
  , CONSTRAINT receipts_PKC PRIMARY KEY (receipt_id)
) COMMENT 'receipts' ;

CREATE INDEX receipts_IX1
  ON receipts(total_price);

CREATE INDEX receipts_IX2
  ON receipts(purchase_date);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK1 FOREIGN KEY (category_id) REFERENCES categories(category_id);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK2 FOREIGN KEY (currency_id) REFERENCES currencies(currency_id);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK3 FOREIGN KEY (shop_id) REFERENCES shops(shop_id);

ALTER TABLE products
  ADD CONSTRAINT products_FK1 FOREIGN KEY (receipt_id) REFERENCES receipts(receipt_id);

