create table products (
  receipt_id INT not null comment 'receipt_id'
  , product_id INT not null comment 'product_id'
  , product_name VARCHAR(100) not null comment 'product_name'
  , price FLOAT not null comment 'price'
  , constraint products_PKC primary key (receipt_id,product_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
