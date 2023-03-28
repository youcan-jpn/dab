create table receipts (
  receipt_id INT not null auto_increment comment 'receipt_id',
  shop_id INT not null comment 'shop_id',
  currency_id INT not null comment 'currency_id',
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

