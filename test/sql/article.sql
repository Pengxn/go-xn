/**
 * This database file is used to test
 *
 * (c) MIT 2018
 */

-- --------------------------------
-- Table structure for article
-- --------------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
    `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `title` text NOT NULL COMMENT 'The title of article',
    `content` longtext NOT NULL COMMENT 'The content of article',
    `article_views` bigint(20) NOT NULL DEFAULT 0 COMMENT 'Number of articles viewed',
    `article_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0: draft; 1: published; 2: deleted',
    `create_time` datetime NOT NULL COMMENT 'The time that article is created',
    `update_time` datetime NOT NULL COMMENT 'The time that article is updated',
    `delete_time` datetime NOT NULL COMMENT 'The time that article is deleted',
    PRIMARY KEY (`ID`),
    KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4;
-- --------------------------------
-- Records of article
-- --------------------------------
INSERT INTO `article` VALUES (
    '1',
    'Everybody listen! We have to put a barrier between us and the world!',
    'At Aspire Themes I use a lot of tools to help me create WordPress, Ghost and Jekyll themes. Tools will range from development, design, servi',
    '0',
    '1',
    '2017-04-07 10:17:50',
    '2019-04-12 16:17:10',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '2',
    'If you have an opportunity to use your voice you should use it',
    'This service is just awesome, I use the Ghost Stack to install Ghost locally. This saves a lot of time and headache installing Ghost, and by',
    '0',
    '1',
    '2017-04-13 10:17:50',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '3',
    'Writing posts with Ghost ✍️',
    'Getting started with the editor is simple, with familiar formatting options in a functional toolbar and the ability to add dynamic content seamlessly.',
    '0',
    '1',
    '2017-06-01 00:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '4',
    'The reason we are gathered here on our God-given, much-naeeded day.',
    'iTerm is a replacement to Mac terminal, and I think most of you are using it for Mac. Tmux is a terminal multilpexer and a pretty great tool',
    '0',
    '1',
    '2017-10-14 08:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '5',
    'Two antarctic penguins took an adorable selfie',
    'Welcome, it is great to have you here. We know that first impressions are important, so we have populated your new site with some initial getting started posts that will help you get familiar with everything in no time.',
    '0',
    '1',
    '2017-12-30 10:17:50',
    '2019-04-12 16:17:10',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '6',
    'Slow cooker honey-dijon glazed carrots',
    'Getting started with the editor is simple, with familiar formatting options in a functional toolbar and the ability to add dynamic content seamlessly.',
    '0',
    '1',
    '2018-01-08 10:17:50',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '7',
    'Managing admin settings',
    'There are a couple of things to do next while you are getting set up: making your site private and inviting your team.',
    '0',
    '1',
    '2018-03-11 00:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '8',
    'Nibh labore ac condimentum sequi ullam',
    'Porro! Mollitia earum congue aliquid? Doloribus. Sociosqu hymenaeos! Ultrices, placerat accusantium iaculis? Irure voluptatibus accumsan odio? Aut, id hymenaeos officia reiciendis dictumst necessitatibus netus, voluptates doloribus porro sodales, eleifend! Mollis.',
    '0',
    '1',
    '2018-04-30 08:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '9',
    'Dicta montes ac doloremque? Exercitation',
    'Welcome, it is great to have you here.We know that first impressions are important, so we have populated your new site with some initial getting started posts that will help you',
    '0',
    '1',
    '2018-10-01 08:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
INSERT INTO `article` VALUES (
    '10',
    'Molestie nostra consequatur. Vivamus aliqua ridiculus',
    'The Ghost editor has everything you need to fully optimise your content. This is where you can add tags and authors, feature a post, or turn a post into a',
    '0',
    '1',
    '2019-01-11 08:12:52',
    '0000-00-00 00:00:00',
    '0000-00-00 00:00:00'
);
