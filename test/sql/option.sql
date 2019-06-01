/**
 * This database file is used to test
 *
 * (c) MIT 2019
 */

-- --------------------------------
-- Table structure for option
-- --------------------------------
DROP TABLE IF EXISTS `option`;
CREATE TABLE `option` (
    `option_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `option_name` varchar(255) UNIQUE NOT NULL COMMENT 'The name of option',
    `option_value` longtext NOT NULL COMMENT 'The value of option',
    PRIMARY KEY (`option_id`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;
-- --------------------------------
-- Records of article
-- --------------------------------
INSERT INTO `option` VALUES (1, "site_url", "https://xn--02f.com");
INSERT INTO `option` VALUES (2, "blog_name", "xn-02f Lab");
INSERT INTO `option` VALUES (3, "blog_description", "✍ The platform for publishing and running your blog.");
INSERT INTO `option` VALUES (4, "users_can_register", "0");
