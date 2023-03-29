create table receipts (
  receipt_id INT not null auto_increment comment 'receipt_id',
  shop_id INT not null comment 'shop_id',
  currency_id INT not null comment 'currency_id',
  category_id INT not null comment 'category_id',
  total_price FLOAT not null comment 'total_price',
  purchase_date DATE not null comment 'purchase_date',
  modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at',
  created_at TIMESTAMP default current_timestamp not null comment 'created_at',
  constraint receipts_PKC primary key (receipt_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


create index receipts_IX1
  on receipts(total_price);

create index receipts_IX2
  on receipts(purchase_date);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK1 FOREIGN KEY (category_id) REFERENCES categories(category_id);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK2 FOREIGN KEY (currency_id) REFERENCES currencies(currency_id);

ALTER TABLE receipts
  ADD CONSTRAINT receipts_FK3 FOREIGN KEY (shop_id) REFERENCES shops(shop_id);
