CREATE TABLE `db`.`items` (`id` INT NOT NULL AUTO_INCREMENT, `product_name` VARCHAR(45) NOT NULL, PRIMARY KEY (`id`));
CREATE TABLE `db`.`receipts` (`id` INT NOT NULL AUTO_INCREMENT, `created_on` DATETIME NULL, PRIMARY KEY (`id`));
CREATE TABLE `db`.`receipt_product` (`receipt_id` INT NOT NULL, `product_id` INT NOT NULL, PRIMARY KEY(`receipt_id`, `product_id`));