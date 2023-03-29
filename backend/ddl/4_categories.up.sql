CREATE TABLE categories (
  category_id INT NOT NULL auto_increment COMMENT 'category_id',
  category_name VARCHAR(50) NOT NULL COMMENT 'category_name',
  CONSTRAINT categories_PKC PRIMARY KEY (category_id),
  INDEX (category_name)  -- 不要だがxoのために指定
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
