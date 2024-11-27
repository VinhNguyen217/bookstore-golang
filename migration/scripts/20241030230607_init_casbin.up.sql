CREATE TABLE `casbin_rule`
(
    `id`    bigint NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) DEFAULT NULL,
    `v0`    varchar(100) DEFAULT NULL,
    `v1`    varchar(100) DEFAULT NULL,
    `v2`    varchar(100) DEFAULT NULL,
    `v3`    varchar(100) DEFAULT NULL,
    `v4`    varchar(100) DEFAULT NULL,
    `v5`    varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO casbin_rule(ptype, v0, v1, v2)
VALUES ('p', 'USER', '/api/v1/users', 'PUT'),
       ('p', 'ADMIN', '/api/v1/users', 'GET'),
       ('p', 'USER', '/api/v1/users/my-info', 'GET'),
       ('p', 'ADMIN', '/api/v1/books', 'POST'),
       ('p', 'ADMIN', '/api/v1/books/*', 'PUT'),
       ('p', 'ADMIN', '/api/v1/books/*', 'DELETE'),
       ('p', 'USER', '/api/v1/carts', 'POST'),
       ('p', 'USER', '/api/v1/carts', 'PUT'),
       ('p', 'USER', '/api/v1/carts', 'GET'),
       ('p', 'USER', '/api/v1/carts/*', 'DELETE'),
       ('p', 'USER', '/api/v1/bills', 'POST'),
       ('p', 'USER', '/api/v1/bills/cancel/*', 'PUT'),
       ('p', 'ADMIN', '/api/v1/bills/update-status/*', 'PUT'),
       ('p', 'USER', '/api/v1/bills/user', 'GET'),
       ('p', 'ADMIN', '/api/v1/bills', 'GET');
