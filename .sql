CREATE TABLE `seati_users`
(
    `id`         INT unsigned NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(50)  NOT NULL,
    `nickname`   VARCHAR(50) DEFAULT NULL,
    `email`      VARCHAR(100) NOT NULL,
    `mcid`       VARCHAR(30)  NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `hash`       VARCHAR(512) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `seati_ecs`
(
    `id` INT unsigned NOT NULL AUTO_INCREMENT,
    `instance_id` VARCHAR(50) NOT NULL,
    `trade_price` FLOAT NOT NULL,
    `region_id` VARCHAR(20) NOT NULL,
    `instance_type` VARCHAR(20) NOT NULL,
    `active` BOOL NOT NULL DEFAULT 1,
    # 关于 status:
    # Pending：创建中
    # Running：运行中
    # Starting：启动中
    # Stopping：停止中
    # Stopped：已停止
    `status` VARCHAR(20) NOT NULL DEFAULT 'Pending',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `seati_ecs_actions`
(
    `id` INT unsigned NOT NULL AUTO_INCREMENT,
    `instance_id` VARCHAR(50) DEFAULT "",
    `action_type` VARCHAR(50) NOT NULL,
    `by_username` VARCHAR(50) DEFAULT "",
    `automated` BOOL NOT NULL DEFAULT 0,
    `at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);