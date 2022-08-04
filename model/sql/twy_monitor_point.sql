/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云测试平台
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 47.111.90.234:33069
 Source Schema         : twy_db

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 14/07/2022 10:33:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for twy_monitor_point
-- 监测点表
-- ----------------------------
DROP TABLE IF EXISTS `twy_monitor_point`;
CREATE TABLE `twy_monitor_point`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `serial_number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '编号',
  `monitor_point_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '监测点名称',
  `abbreviation` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称缩写',
  `point_type` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型：综保/局放/测温/微水/油色谱/机器人/其他',
  `point_category` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类别：遥测/遥信',
  `state` tinyint NOT NULL DEFAULT 1 COMMENT '状态(1=公开,2=私人,3=禁用)',
  `unit` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '单位',
  `coefficient` double(20, 2) NULL DEFAULT 1.00 COMMENT '系数',
  `type_purpose` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型用途：1 一次系统图',
  `type_screen` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型用途：1 大屏数据  ',
  `type_other` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型用途：1 其他',
  `rules_state` tinyint NOT NULL DEFAULT 0 COMMENT '是否开启规则：0关闭，1开启',
  `upper_threshold` bigint NOT NULL DEFAULT 0 COMMENT '上限阀值',
  `lower_threshold` bigint NOT NULL DEFAULT 0 COMMENT '下限阀值',
  `source_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源名称',
  `source_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源ID',
  `asset_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资产ID',
  `asset_group_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资产组ID',
  `tenant_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '租户ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11310 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
