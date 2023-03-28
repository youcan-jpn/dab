DELETE FROM products;
DELETE FROM receipts;
DELETE FROM shops;
DELETE FROM currencies;

INSERT INTO shops (
    shop_id,  -- 実際はAUTO_INCREMENT
    shop_name
)
VALUES (
    1,
    "sample_shop_1"
), (
    2,
    "sample_shop_2"
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

INSERT INTO receipts (
    receipt_id,  -- 実際はAUTO_INCREMENT
    shop_id,
    currency_id,
    total_price,
    purchase_date
)
VALUES (
    1,
    1,
    1,
    900,
    "2022-03-01"
), (
    2,
    2,
    2,
    30,
    "2012-08-29"
);