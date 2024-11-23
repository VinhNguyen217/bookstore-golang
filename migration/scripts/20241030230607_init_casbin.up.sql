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

-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'USER', '/api/v1/auth/sign-in', 'POST');
-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'admin', '/api/v1/books', 'PUT');
-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'admin', '/api/v1/books', 'GET');
-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'user', '/api/v1/books', 'GET');
-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'user', '/api/v1/carts', 'POST');
-- INSERT INTO casbin_rule(ptype, v0, v1, v2)
-- VALUES ('p', 'user', '/api/v1/carts', 'PUT');