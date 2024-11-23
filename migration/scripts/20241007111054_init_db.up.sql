DROP TABLE IF EXISTS book;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS bill;
DROP TABLE IF EXISTS cart;
DROP TABLE IF EXISTS bill_detail;

CREATE TABLE `book`
(
    `id`           bigint NOT NULL AUTO_INCREMENT,
    `name`         longtext,
    `quantity`     bigint DEFAULT NULL,
    `sold`         bigint DEFAULT NULL,
    `price`        bigint DEFAULT NULL,
    `publish_date` datetime(3) DEFAULT NULL,
    `description`  longtext,
    `created_at`   datetime(3) DEFAULT NULL,
    `updated_at`   datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `user`
(
    `id`         bigint NOT NULL AUTO_INCREMENT,
    `name`       longtext,
    `username`   longtext,
    `password`   longtext,
    `salt`       longtext,
    `role`       longtext,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `cart`
(
    `id`         bigint NOT NULL AUTO_INCREMENT,
    `user_id`    bigint DEFAULT NULL,
    `book_id`    bigint DEFAULT NULL,
    `quantity`   bigint DEFAULT NULL,
    `price`      bigint DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `bill`
(
    `id`           bigint NOT NULL AUTO_INCREMENT,
    `receiver`     longtext,
    `phone`        longtext,
    `address`      longtext,
    `email`        longtext,
    `note`         longtext,
    `total`        bigint DEFAULT NULL,
    `status`       longtext,
    `payment`      longtext,
    `confirm_date` datetime(3) DEFAULT NULL,
    `created_at`   datetime(3) DEFAULT NULL,
    `updated_at`   datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `bill_detail`
(
    `id`         bigint NOT NULL AUTO_INCREMENT,
    `bill_id`    bigint DEFAULT NULL,
    `book_id`    bigint DEFAULT NULL,
    `quantity`   bigint DEFAULT NULL,
    `price`      bigint DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO user(name,username, password, salt,role,created_at,updated_at)
VALUES ('Admin', 'admin',
        'f4ee55daf6dbb929ec8194f9e1bdc232cfd68f7ae1d9164855bb7810039a9345b85babb7eadc4ad5c4cc000c6adb7ccc6955da4218508f94db6e3794872ead4a',
        'msvcAWgtMVFAiiBQOkzD',
        'ADMIN',
        '2024-11-21 17:15:20.910',
    '2024-11-21 17:15:20.910');