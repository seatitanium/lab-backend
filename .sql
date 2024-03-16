CREATE TABLE `seati_users`
(
    `id`         INT unsigned NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(50)  NOT NULL,
    `nickname`   VARCHAR(50) DEFAULT NULL,
    `email`      VARCHAR(100) NOT NULL,
    `mcid`       VARCHAR(30)  NOT NULL,
    `created_at` TIMESTAMP    NOT NULL,
    `updated_at` TIMESTAMP    NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    `hash`       VARCHAR(512) NOT NULL,
    PRIMARY KEY (`id`)
);