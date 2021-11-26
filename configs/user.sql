SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `account` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `password` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `email`	varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `avatar_ext` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci,
                          PRIMARY KEY (`account`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
