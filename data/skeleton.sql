/*
 Navicat Premium Data Transfer

 Source Server         : spots
 Source Server Type    : MySQL
 Source Server Version : 50649
 Source Host           : 8.210.163.66:3306
 Source Schema         : spots

 Target Server Type    : MySQL
 Target Server Version : 50649
 File Encoding         : 65001

 Date: 09/02/2021 16:39:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `login_name` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录密码',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态，0-禁用，1-启用',
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `gmt_modified` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
) ENGINE = InnoDB AUTO_INCREMENT = 92 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;

-- insert admin
INSERT INTO `spots`.`user`(`id`, `username`, `login_name`, `password`, `status`, `gmt_create`, `gmt_modified`) VALUES (1, '管理员', 'admin', '$2a$10$1TjD6BV97RAjgzbt4nhwQekvvwRR/rn/O5XbYCVFi1KhKJz2d/GR6', 1, '2020-11-13 17:47:53', '2021-02-07 01:00:00');
