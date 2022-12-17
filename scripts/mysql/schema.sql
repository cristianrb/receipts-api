CREATE TABLE `db`.`items` (`id` INT NOT NULL AUTO_INCREMENT, `product_name` VARCHAR(45) NOT NULL UNIQUE, PRIMARY KEY (`id`));
CREATE TABLE `db`.`receipts` (`id` INT NOT NULL AUTO_INCREMENT, `created_on` DATETIME NULL, PRIMARY KEY (`id`));
CREATE TABLE `db`.`receipt_product` (`receipt_id` INT NOT NULL, `product_id` INT NOT NULL,
                                     PRIMARY KEY(`receipt_id`, `product_id`),
                                     FOREIGN KEY (receipt_id) REFERENCES receipts (id) ON DELETE CASCADE,
                                     FOREIGN KEY (product_id) REFERENCES items (id) ON DELETE CASCADE
);