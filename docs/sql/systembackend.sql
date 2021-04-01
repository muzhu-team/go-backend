/*
 Navicat Premium Data Transfer

 Source Server         : go-gin
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3308
 Source Schema         : systembackend

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 01/04/2021 15:03:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int UNSIGNED NULL DEFAULT 0 COMMENT '标签ID',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '简述',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `cover_image_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '封面图片地址',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '新建时间',
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '标签名称',
  `created_on` time(6) NOT NULL DEFAULT '00:00:00.000000' COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章标签管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
