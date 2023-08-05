/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : subject

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 05/08/2023 11:08:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nation` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `periodical` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `organization` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `create_date` date NULL DEFAULT NULL,
  `technique` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `key_word` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `digest` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram`
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articles
-- ----------------------------
INSERT INTO `articles` VALUES (1, 'China', 'title', '123', '123', '123', '2023-08-01', '123', 'key', 'digest', 'data', 'text', 0, 0, '2023-08-04 15:53:46.035', '2023-08-04 15:53:46.035');
INSERT INTO `articles` VALUES (3, 'China', 'title', '123', '123', '123', '2023-08-01', '123', 'key', 'digest', 'data', 'text', 0, 0, '2023-08-04 15:53:47.173', '2023-08-04 15:53:47.173');

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nation` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `language` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `publisher` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `time` datetime NULL DEFAULT NULL,
  `digest` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram`
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES (1, 'reo', 'China', 'chinese', 'Trump', '123', '2023-08-01 15:30:45', '123', 'text', 0, 0, '2023-08-04 15:52:19.035', '2023-08-04 15:52:19.035');
INSERT INTO `books` VALUES (3, 'reo', 'China', 'chinese', 'Trump', '123', '2023-08-01 15:30:45', '123', 'text', 0, 0, '2023-08-04 15:52:20.354', '2023-08-04 15:52:20.354');

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `foreign_key` bigint NULL DEFAULT NULL,
  `tablee` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `category1` tinyint(1) NULL DEFAULT NULL,
  `category2` tinyint(1) NULL DEFAULT NULL,
  `category3` tinyint(1) NULL DEFAULT NULL,
  `category4` tinyint(1) NULL DEFAULT NULL,
  `category5` tinyint(1) NULL DEFAULT NULL,
  `category6` tinyint(1) NULL DEFAULT NULL,
  `category7` tinyint(1) NULL DEFAULT NULL,
  `category8` tinyint(1) NULL DEFAULT NULL,
  `category9` tinyint(1) NULL DEFAULT NULL,
  `category10` tinyint(1) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES (1, 1, 'unique_databases', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-04 14:45:05.211', '2023-08-04 14:45:05.211');
INSERT INTO `categories` VALUES (3, 3, 'unique_databases', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-04 14:45:06.280', '2023-08-04 14:45:06.280');
INSERT INTO `categories` VALUES (4, 4, 'unique_databases', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-04 14:45:06.737', '2023-08-04 14:45:06.737');
INSERT INTO `categories` VALUES (5, 5, 'unique_databases', 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, '2023-08-04 15:21:07.152', '2023-08-04 15:21:07.152');
INSERT INTO `categories` VALUES (6, 1, 'teamworks', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-08-04 15:48:19.007', '2023-08-04 15:48:19.007');
INSERT INTO `categories` VALUES (8, 3, 'teamworks', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-08-04 15:48:20.278', '2023-08-04 15:48:20.278');
INSERT INTO `categories` VALUES (9, 4, 'teamworks', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-08-04 15:48:20.773', '2023-08-04 15:48:20.773');
INSERT INTO `categories` VALUES (10, 5, 'teamworks', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-08-04 15:48:21.219', '2023-08-04 15:48:21.219');
INSERT INTO `categories` VALUES (11, 1, 'projects', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-04 15:49:12.876', '2023-08-04 15:49:12.876');
INSERT INTO `categories` VALUES (13, 3, 'projects', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-04 15:49:14.000', '2023-08-04 15:49:14.000');
INSERT INTO `categories` VALUES (14, 1, 'institutes', 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, '2023-08-04 15:51:28.935', '2023-08-04 15:51:28.935');
INSERT INTO `categories` VALUES (16, 3, 'institutes', 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, '2023-08-04 15:51:43.225', '2023-08-04 15:51:43.225');
INSERT INTO `categories` VALUES (17, 4, 'institutes', 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, '2023-08-04 15:51:43.582', '2023-08-04 15:51:43.582');
INSERT INTO `categories` VALUES (18, 1, 'books', 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, '2023-08-04 15:52:19.039', '2023-08-04 15:52:19.039');
INSERT INTO `categories` VALUES (20, 3, 'books', 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, '2023-08-04 15:52:20.362', '2023-08-04 15:52:20.362');
INSERT INTO `categories` VALUES (21, 1, 'dissertations', 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, '2023-08-04 15:52:46.109', '2023-08-04 15:52:46.109');
INSERT INTO `categories` VALUES (23, 1, 'articles', 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, '2023-08-04 15:53:46.048', '2023-08-04 15:53:46.048');
INSERT INTO `categories` VALUES (25, 3, 'articles', 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, '2023-08-04 15:53:47.186', '2023-08-04 15:53:47.186');
INSERT INTO `categories` VALUES (26, 0, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:00:36.243', '2023-08-04 16:00:36.243');
INSERT INTO `categories` VALUES (27, 0, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:00:38.313', '2023-08-04 16:00:38.313');
INSERT INTO `categories` VALUES (28, 5, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:01:55.806', '2023-08-04 16:01:55.806');
INSERT INTO `categories` VALUES (29, 6, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:02:00.486', '2023-08-04 16:02:00.486');
INSERT INTO `categories` VALUES (30, 7, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:02:19.083', '2023-08-04 16:02:19.083');
INSERT INTO `categories` VALUES (31, 1, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:02:36.706', '2023-08-04 16:02:36.706');
INSERT INTO `categories` VALUES (33, 3, 'tutors', 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, '2023-08-04 16:02:37.978', '2023-08-04 16:07:47.193');
INSERT INTO `categories` VALUES (34, 4, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:02:38.500', '2023-08-04 16:02:38.500');
INSERT INTO `categories` VALUES (35, 5, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, '2023-08-04 16:06:16.144', '2023-08-04 16:06:16.144');
INSERT INTO `categories` VALUES (36, 6, 'tutors', 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, '2023-08-04 16:06:23.125', '2023-08-04 16:06:23.125');
INSERT INTO `categories` VALUES (37, 4, 'projects', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-05 03:00:23.493', '2023-08-05 03:00:23.493');
INSERT INTO `categories` VALUES (38, 5, 'projects', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-05 03:02:18.844', '2023-08-05 03:02:18.844');
INSERT INTO `categories` VALUES (39, 6, 'projects', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, '2023-08-05 03:02:19.937', '2023-08-05 03:02:19.937');

-- ----------------------------
-- Table structure for dissertations
-- ----------------------------
DROP TABLE IF EXISTS `dissertations`;
CREATE TABLE `dissertations`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `tutor` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `province` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `city` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `university` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `college` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `date` datetime NULL DEFAULT NULL,
  `technique` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `key_word` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `digest` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram`
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dissertations
-- ----------------------------
INSERT INTO `dissertations` VALUES (1, 'title', 'author', 'tutor', 'province', 'city', 'university', 'collegeg', '2023-08-01 15:30:45', '123', '123,456', '123', 'abcdefg', 'text', 0, 0, '2023-08-04 15:52:46.097', '2023-08-04 15:52:46.097');

-- ----------------------------
-- Table structure for institutes
-- ----------------------------
DROP TABLE IF EXISTS `institutes`;
CREATE TABLE `institutes`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `university` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `college` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nation` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `province` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `city` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `classification` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of institutes
-- ----------------------------
INSERT INTO `institutes` VALUES (1, '123', '123', '123', 'china', '山东', '潍坊', '123', 0, '2023-08-04 15:51:28.920', '2023-08-04 15:51:28.920');
INSERT INTO `institutes` VALUES (3, '123', '123', '123', 'china', '山东', '潍坊', '123', 0, '2023-08-04 15:51:43.212', '2023-08-04 15:51:43.212');
INSERT INTO `institutes` VALUES (4, '123', '123', '123', 'china', '山东', '潍坊', '123', 0, '2023-08-04 15:51:43.579', '2023-08-04 15:51:43.579');

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `module` bigint UNSIGNED NULL DEFAULT NULL,
  `department` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint UNSIGNED NULL DEFAULT 0,
  `date` datetime NULL DEFAULT NULL,
  `region` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `category` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `text`(`text`) WITH PARSER `ngram`,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram`,
  FULLTEXT INDEX `title_2`(`title`, `text`) WITH PARSER `ngram`
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of news
-- ----------------------------

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `classification` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `sponsor` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `approval_number` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `create_date` datetime NULL DEFAULT NULL,
  `superintendent` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `organization` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `region` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `data` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram`
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of projects
-- ----------------------------
INSERT INTO `projects` VALUES (1, '123', '123', 'trump', '123456', '2023-08-01 15:30:45', 'trump', 'china', 'p', 0, 0, 0, '2023-08-04 15:49:12.871', '2023-08-04 15:49:12.871');
INSERT INTO `projects` VALUES (3, '123', '123', 'trump', '123456', '2023-08-01 15:30:45', 'trump', 'china', 'p', 0, 0, 0, '2023-08-04 15:49:13.988', '2023-08-04 15:49:13.988');
INSERT INTO `projects` VALUES (4, '123', '123', 'trump', '123456', '2023-08-01 15:30:45', 'trump', 'china', 'p', 0, 0, 0, '2023-08-05 03:00:23.482', '2023-08-05 03:00:23.482');
INSERT INTO `projects` VALUES (5, '123', '123', 'trump', '123456', '2023-08-01 15:30:45', 'trump', 'china', 'p', 0, 0, 0, '2023-08-05 03:02:18.839', '2023-08-05 03:02:18.839');
INSERT INTO `projects` VALUES (6, '123', '123', 'trump', '123456', '2023-08-01 15:30:45', 'trump', 'china', 'p', 0, 0, 0, '2023-08-05 03:02:19.924', '2023-08-05 03:02:19.924');

-- ----------------------------
-- Table structure for teamworks
-- ----------------------------
DROP TABLE IF EXISTS `teamworks`;
CREATE TABLE `teamworks`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `grade` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `direction` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `sponsor` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `number` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `principal` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `member` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `province` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `city` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `county` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `picture` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `text` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teamworks
-- ----------------------------
INSERT INTO `teamworks` VALUES (1, 'name', '3', 'horizontal', 'Trump', '114514', '1145.1.4', 'homo', '1,23,456', '37', '07', '04', 'http://', '', 0, 0, '2023-08-04 15:48:18.992', '2023-08-04 15:48:18.992');
INSERT INTO `teamworks` VALUES (3, 'name', '3', 'horizontal', 'Trump', '114514', '1145.1.4', 'homo', '1,23,456', '37', '07', '04', 'http://', '', 0, 0, '2023-08-04 15:48:20.267', '2023-08-04 15:48:20.267');
INSERT INTO `teamworks` VALUES (4, 'name', '3', 'horizontal', 'Trump', '114514', '1145.1.4', 'homo', '1,23,456', '37', '07', '04', 'http://', '', 0, 0, '2023-08-04 15:48:20.768', '2023-08-04 15:48:20.768');
INSERT INTO `teamworks` VALUES (5, 'name', '3', 'horizontal', 'Trump', '114514', '1145.1.4', 'homo', '1,23,456', '37', '07', '04', 'http://', '', 0, 0, '2023-08-04 15:48:21.208', '2023-08-04 15:48:21.208');

-- ----------------------------
-- Table structure for tutors
-- ----------------------------
DROP TABLE IF EXISTS `tutors`;
CREATE TABLE `tutors`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `university` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `college` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nation` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `province` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `city` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `direction` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `contact` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `classification` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tutors
-- ----------------------------
INSERT INTO `tutors` VALUES (1, 'Trump', 'xjtu', 'e a p', 'China', 'sd', 'wf', 'xxx', '教授', 'xxx@xjtu.edu.cn', '博导/硕导', '2023-08-04 16:02:36.700', '2023-08-04 16:02:36.700');
INSERT INTO `tutors` VALUES (3, 'trump', 'xjtu', 'e a p', 'China', 'sd', 'wf', 'xxx', '教授', 'xxx@xjtu.edu.cn', '博导/硕导', '2023-08-04 16:02:37.972', '2023-08-04 16:07:47.189');
INSERT INTO `tutors` VALUES (4, 'Trump', 'xjtu', 'e a p', 'China', 'sd', 'wf', 'xxx', '教授', 'xxx@xjtu.edu.cn', '博导/硕导', '2023-08-04 16:02:38.496', '2023-08-04 16:02:38.496');
INSERT INTO `tutors` VALUES (5, 'Trump', 'xjtu', 'e a p', 'China', 'sd', 'wf', 'xxx', '教授', 'xxx@xjtu.edu.cn', '博导/硕导', '2023-08-04 16:06:16.139', '2023-08-04 16:06:16.139');
INSERT INTO `tutors` VALUES (6, 'Trump', 'xjtu', 'e a p', 'China', 'sd', 'wf', 'xxx', '教授', 'xxx@xjtu.edu.cn', '博导/硕导', '2023-08-04 16:06:23.112', '2023-08-04 16:06:23.112');

-- ----------------------------
-- Table structure for unique_databases
-- ----------------------------
DROP TABLE IF EXISTS `unique_databases`;
CREATE TABLE `unique_databases`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `trimmer` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `key_word` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `introduction` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `attachment` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `click` bigint NULL DEFAULT NULL,
  `download` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of unique_databases
-- ----------------------------
INSERT INTO `unique_databases` VALUES (1, 'fauihfaw', 'agwreua', '2131', 'fahifbnwjnf', '123', 0, 0, '2023-08-04 14:45:05.207', '2023-08-04 14:45:05.207');
INSERT INTO `unique_databases` VALUES (3, 'fauihfaw', 'agwreua', '2131', 'fahifbnwjnf', '123', 0, 0, '2023-08-04 14:45:06.266', '2023-08-04 14:45:06.266');
INSERT INTO `unique_databases` VALUES (4, 'fauihfaw', 'agwreua', '2131', 'fahifbnwjnf', '123', 0, 0, '2023-08-04 14:45:06.732', '2023-08-04 14:45:06.732');
INSERT INTO `unique_databases` VALUES (5, 'fauihfaw', 'agwreua', '2131', 'fahifbnwjnf', '123', 0, 0, '2023-08-04 15:21:07.139', '2023-08-04 15:21:07.139');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `usertype` bigint NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
