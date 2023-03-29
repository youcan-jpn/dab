create table shops (
  shop_id INT not null auto_increment PRIMARY KEY comment 'shop_id',
  shop_name VARCHAR(100) not null comment 'shop_name',
  modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at',
  created_at TIMESTAMP default current_timestamp not null comment 'created_at',
  INDEX (shop_name)  -- 不要だがxoのために指定
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
