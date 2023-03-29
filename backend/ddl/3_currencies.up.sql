create table currencies (
  currency_id INT not null auto_increment PRIMARY KEY comment 'currency_id'
  , currency_name CHAR(3) not null comment 'currency_name'
  , in_yen FLOAT not null comment 'in_yen'
  , modified_at TIMESTAMP default current_timestamp on update current_timestamp not null comment 'modified_at'
  , created_at TIMESTAMP default current_timestamp not null comment 'created_at',
  INDEX (in_yen)  -- 不要だがxoのために指定
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
