-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `name` varchar(20) NOT NULL COMMENT `用户名称`,
    `age` bigint NOT NULL COMMENT `用户年龄`,
    `passwrd` bigint NOT NULL COMMENT `账号密码`,
    `confirmpasswrd` varchar NOT NULL COMMENT `确认密码`，
) ENGINE InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE 