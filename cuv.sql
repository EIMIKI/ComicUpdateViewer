DROP TABLE IF EXISTS `comics`;

CREATE TABLE `comics` (
  `comic_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(40) NOT NULL,
  `url` varchar(255) NOT NULL,
  `img` varchar(255) DEFAULT NULL,
  `date` date DEFAULT NULL,
  PRIMARY KEY (`comic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `comics` WRITE;

UNLOCK TABLES;