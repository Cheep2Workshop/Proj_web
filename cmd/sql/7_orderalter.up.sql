ALTER TABLE `order_details`
ADD COLUMN `order_id` BIGINT,
ADD FOREIGN KEY(`order_id`) REFERENCES `orders`(`id`);

ALTER TABLE `orders`
ADD FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

