DELETE FROM products;
DELETE FROM receipts;
DELETE FROM shops;
DELETE FROM currencies;
DELETE FROM categories;

INSERT INTO shops (
    shop_id,  -- 実際はAUTO_INCREMENT
    shop_name,
    modified_at,  -- 実際は自動挿入
    created_at  -- 実際は自動挿入
)
VALUES (
    1,
    "sample_shop_1",
    '2011-01-01 01:01:01',
    '2011-01-01 01:00:00'
), (
    2,
    "sample_shop_2",
    '2011-01-02 02:02:02',
    '2011-01-02 02:00:00'
);

INSERT INTO currencies (
    currency_id,  -- 実際はAUTO_INCREMENT
    currency_name,
    in_yen
)
VALUES (
    1,
    "JPY",
    1
), (
    2,
    "CHF",
    142.25
);

INSERT INTO categories (
    category_id,
    category_name
)
VALUES (
    1,
    N'食費'
), (
    2,
    N'娯楽費'
);

INSERT INTO receipts (
    receipt_id,  -- 実際はAUTO_INCREMENT
    shop_id,
    currency_id,
    category_id,
    total_price,
    purchase_date
)
VALUES (
    1,
    1,
    1,
    1,
    900,
    "2022-03-01"
), (
    2,
    2,
    2,
    2,
    30,
    "2012-08-29"
);

INSERT INTO products (
    receipt_id,
    product_id,
    product_name,
    price
)
VALUES (
    1,
    1,
    "sample_product_1_1",
    900
), (
    2,
    1,
    "sample_product_2_1",
    10
), (
    2,
    2,
    "sample_product_2_2",
    20
);