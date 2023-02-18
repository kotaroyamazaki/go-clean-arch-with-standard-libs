CREATE TABLE `users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `nickname` VARCHAR(255) NOT NULL,
  `birth_date` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB;

CREATE TABLE `books` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `price` INT NOT NULL,
  `description` VARCHAR(1024) NOT NULL,
  `imageURL` text,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB;

CREATE TABLE `user_favorite_books` (
  `user_id` INT NOT NULL,
  `book_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT `fk_user_favorite_books_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_user_favorite_books_book_id` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
  PRIMARY KEY (`user_id`,`book_id`)
)ENGINE=InnoDB;

