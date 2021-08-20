ALTER TABLE `products`
ADD COLUMN `delete_at` BIGINT;

ALTER TABLE `users`
DROP COLUMN `test3`;