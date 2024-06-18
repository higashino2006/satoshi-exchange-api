CREATE TABLE `users` (
  `id` VARCHAR(255) PRIMARY KEY NOT NULL,
  `jpy_balance` DOUBLE(10, 3) NOT NULL DEFAULT 0,
  `satoshi_balance` DOUBLE(10, 3) NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL
);

CREATE TABLE `trade_records` (
  `id` INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  `user_id` VARCHAR(255) NOT NULL,
  `type` VARCHAR(255) NOT NULL,
  `jpy` DOUBLE(10, 3) NOT NULL,
  `satoshi` DOUBLE(10, 3) NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);