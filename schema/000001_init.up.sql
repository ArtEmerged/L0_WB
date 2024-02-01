CREATE TABLE orders 
(
    order_uid VARCHAR(20) NOT NULL UNIQUE,
    track_number VARCHAR(12) NOT NULL,
    entry VARCHAR(4) NOT NULL,
    locale VARCHAR(3) NOT NULL,
    internal_signature VARCHAR(15),
    customer_id VARCHAR(5) NOT NULL,
    delivery_service VARCHAR(10) NOT NULL,
    shardkey VARCHAR(2) NOT NULL,
    sm_id  SMALLSERIAL NOT NULL, 
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR(1) NOT NULL
);

CREATE TABLE delivery 
(
    order_uid VARCHAR(20) REFERENCES orders (order_uid) ON DELETE CASCADE NOT NULL,
    name VARCHAR(10) NOT NULL,
    phone VARCHAR(12) NOT NULL,
    zip VARCHAR(7) NOT NULL,
    city VARCHAR(20) NOT NULL,
    addres VARCHAR(40) NOT NULL,
    region VARCHAR(20) NOT NULL,
    email VARCHAR(40) NOT NULL
);

CREATE TABLE payments
(
    order_uid VARCHAR(20) REFERENCES orders (order_uid) ON DELETE CASCADE NOT NULL,
    transaction VARCHAR(20) NOT NULL,
    request_id VARCHAR(10) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    provider VARCHAR(6) NOT NULL,
    amount SERIAL NOT NULL,
    payment_dt BIGSERIAL NOT NULL,
    bank VARCHAR(10) NOT NULL,
    delivery_cost SERIAL NOT NULL,
    goods_total SERIAL NOT NULL,
    custom_fee SERIAL 
); 

CREATE TABLE items
(
    order_uid VARCHAR(20) REFERENCES orders (order_uid) ON DELETE CASCADE NOT NULL,
    chrt_id SERIAL NOT NULL,
    track_number VARCHAR(20) NOT NULL,
    price SERIAL NOT NULL,
    rid VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL,
    sale SMALLSERIAL,
    size VARCHAR(10),
    total_price SERIAL NOT NULL,
    nm_id SERIAL,
    brand VARCHAR(30) NOT NULL,
    status SMALLSERIAL NOT NULL
);
