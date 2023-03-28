create table shops (
  shop_id INT not null auto_increment comment 'shop_id'
  , shop_name VARCHAR(100) not null comment 'shop_name'
  , modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at'
  , created_at TIMESTAMP default current_timestamp not null comment 'created_at'
  , constraint shops_PKC primary key (shop_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
