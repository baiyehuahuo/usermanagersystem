/*
 Navicat Premium Data Transfer

 Source Server         : xiaoyan
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 26/11/2021 21:20:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `account` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `nick_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `avatar_ext` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                          PRIMARY KEY (`account`) USING BTREE,
                          UNIQUE INDEX `email_unique_index`(`email`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
