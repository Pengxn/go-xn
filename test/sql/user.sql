/**
 * This database file is used to test
 *
 * (c) MIT 2021
 */

-- --------------------------------
-- Table structure for user
-- --------------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(18) NOT NULL COMMENT 'The name of user',
    `nick_name` varchar(25) NOT NULL COMMENT 'The nick name of user',
    `password` varchar(20) NOT NULL COMMENT 'The password of user',
    `status` tinyint(4) NOT NULL COMMENT '0: not verified; 1: verified; 2: deleted',
    `email` varchar(50) NOT NULL COMMENT 'The email of user',
    `role` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0: guest; 1: admin; 2: owner',
    `create_time` datetime NOT NULL COMMENT 'The time that user is created',
    `update_time` datetime NOT NULL COMMENT 'The time that user is updated',
    `delete_time` datetime NOT NULL COMMENT 'The time that user is deleted',
    PRIMARY KEY (`ID`),
    KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;
-- --------------------------------
-- Records of user
-- --------------------------------
INSERT INTO `user` VALUES (
    '1',
    'fyj',
    'Feng.YJ',
    'qw123456',
    '1',
    'i@huiyifyj.cn',
    '1',
    '2017-04-07 10:17:50',
    '2019-04-12 16:17:10',
    '0000-00-00 00:00:00'
);
INSERT INTO `user` VALUES (
    '2',
    'test',
    'Test',
    'password',
    '0',
    'i@test.com',
    '0',
    '2017-04-13 10:17:50',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
