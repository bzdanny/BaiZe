/*
 Navicat Premium Data Transfer

 Source Server         : 华为
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : 116.63.188.235:3306
 Source Schema         : baize

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 07/08/2022 15:54:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '部门名称',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10974733240102914 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` VALUES (100, 0, '0', '白泽科技', 0, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 1, '2021-09-28 11:12:15');
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '深圳总公司', 1, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 1, '2021-09-28 11:12:15');
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', '长沙分公司', 2, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 1, '2021-09-28 11:11:58');
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '研发部门', 1, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', '市场部门', 2, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', '测试部门', 3, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', '财务部门', 4, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', '运维部门', 5, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', '市场部门', 1, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 0, NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', '财务部门', 2, '白泽', '15888888888', 'ry@qq.com', '0', '0', 1, '2021-08-15 12:02:04', 1, '2021-09-28 11:11:58');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64208701686288384 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '性别男');
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '性别女');
INSERT INTO `sys_dict_data` VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '性别未知');
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '显示菜单');
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '隐藏菜单');
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '默认分组');
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '系统分组');
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '系统默认是');
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '系统默认否');
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '通知');
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '公告');
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 1, '2021-08-15 12:02:09', 0, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '关闭状态');
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '新增操作');
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '修改操作');
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '删除操作');
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '授权操作');
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '导出操作');
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '导入操作');
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '强退操作');
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '生成操作');
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '清空操作');
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 1, '2021-08-15 12:02:09', 0, NULL, '停用状态');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE KEY `dict_type` (`dict_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64208610808303616 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 1, '2021-08-15 12:02:09', 0, NULL, '用户性别列表');
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', 1, '2021-08-15 12:02:09', 0, NULL, '菜单状态列表');
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 1, '2021-08-15 12:02:09', 0, NULL, '系统开关列表');
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 1, '2021-08-15 12:02:09', 0, NULL, '任务状态列表');
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 1, '2021-08-15 12:02:09', 0, NULL, '任务分组列表');
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 1, '2021-08-15 12:02:09', 0, NULL, '系统是否列表');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 1, '2021-08-15 12:02:09', 0, NULL, '通知类型列表');
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', 1, '2021-08-15 12:02:09', 0, NULL, '通知状态列表');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 1, '2021-08-15 12:02:09', 0, NULL, '操作类型列表');
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 1, '2021-08-15 12:02:09', 0, NULL, '登录状态列表');
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `job_params` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '参数',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'cron执行表达式',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` VALUES (1, '系统默认（无参）', 'DEFAULT', NULL, 'NoParams', '0/1 * * * * *', '1', 1, '2021-08-15 12:02:11', 0, '2021-12-05 12:16:58', '');
INSERT INTO `sys_job` VALUES (2, '系统默认（有参）', 'DEFAULT', '{\"name\":\"baize\",\"age\":1}', 'Params', '0/15 * * * * ?', '1', 1, '2021-08-15 12:02:11', 0, '2021-12-05 14:38:01', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor` (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '用户账号',
  `ipaddr` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=75594893786877952 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
  `permission_id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `permission_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint DEFAULT '0' COMMENT '父菜单ID',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `perms` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '权限标识',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64195037163950080 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
BEGIN;
INSERT INTO `sys_permission` VALUES (1, '系统管理', 0, '0', 'system', 1, '2021-08-15 12:02:06', 0, NULL, '系统管理目录');
INSERT INTO `sys_permission` VALUES (2, '系统监控', 0, '0', 'monitor', 1, '2021-08-15 12:02:06', 0, NULL, '系统监控目录');
INSERT INTO `sys_permission` VALUES (3, '系统工具', 0, '0', 'tool', 1, '2021-08-15 12:02:06', 0, NULL, '系统工具目录');
INSERT INTO `sys_permission` VALUES (4, '白泽官网', 0, '0', 'baize', 1, '2021-08-15 12:02:06', 1, '2022-08-07 13:04:04', '若依官网地址');
INSERT INTO `sys_permission` VALUES (100, '用户管理', 1, '0', 'system:user:list', 1, '2021-08-15 12:02:06', 0, NULL, '用户管理菜单');
INSERT INTO `sys_permission` VALUES (101, '角色管理', 1, '0', 'system:role:list', 1, '2021-08-15 12:02:06', 0, NULL, '角色管理菜单');
INSERT INTO `sys_permission` VALUES (102, '菜单管理', 1, '0', 'system:menu:list', 1, '2021-08-15 12:02:06', 0, NULL, '菜单管理菜单');
INSERT INTO `sys_permission` VALUES (103, '部门管理', 1, '0', 'system:dept:list', 1, '2021-08-15 12:02:06', 0, NULL, '部门管理菜单');
INSERT INTO `sys_permission` VALUES (104, '岗位管理', 1, '0', 'system:post:list', 1, '2021-08-15 12:02:06', 0, NULL, '岗位管理菜单');
INSERT INTO `sys_permission` VALUES (105, '字典管理', 1, '0', 'system:dict:list', 1, '2021-08-15 12:02:06', 0, NULL, '字典管理菜单');
INSERT INTO `sys_permission` VALUES (108, '日志管理', 1, '0', 'logininfor', 1, '2021-08-15 12:02:06', 0, NULL, '日志管理菜单');
INSERT INTO `sys_permission` VALUES (109, '在线用户', 2, '0', 'monitor:online:list', 1, '2021-08-15 12:02:06', 0, NULL, '在线用户菜单');
INSERT INTO `sys_permission` VALUES (110, '定时任务', 2, '0', 'monitor:job:list', 1, '2021-08-15 12:02:06', 0, NULL, '定时任务菜单');
INSERT INTO `sys_permission` VALUES (112, '服务监控', 2, '0', 'monitor:server:list', 1, '2021-08-15 12:02:06', 0, NULL, '服务监控菜单');
INSERT INTO `sys_permission` VALUES (114, '表单构建', 3, '0', 'tool:build:list', 1, '2021-08-15 12:02:06', 0, NULL, '表单构建菜单');
INSERT INTO `sys_permission` VALUES (115, '代码生成', 3, '0', 'tool:gen:list', 1, '2021-08-15 12:02:06', 0, NULL, '代码生成菜单');
INSERT INTO `sys_permission` VALUES (116, '系统接口', 3, '0', 'tool:swagger:list', 1, '2021-08-15 12:02:06', 0, NULL, '系统接口菜单');
INSERT INTO `sys_permission` VALUES (501, '登录日志', 108, '0', 'monitor:logininfor:list', 1, '2021-08-15 12:02:06', 0, NULL, '登录日志菜单');
INSERT INTO `sys_permission` VALUES (1001, '用户查询', 100, '0', 'system:user:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1002, '用户新增', 100, '0', 'system:user:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1003, '用户修改', 100, '0', 'system:user:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1004, '用户删除', 100, '0', 'system:user:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1005, '用户导出', 100, '0', 'system:user:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1006, '用户导入', 100, '0', 'system:user:import', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1007, '重置密码', 100, '0', 'system:user:resetPwd', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1008, '角色查询', 101, '0', 'system:role:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1009, '角色新增', 101, '0', 'system:role:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1010, '角色修改', 101, '0', 'system:role:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1011, '角色删除', 101, '0', 'system:role:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1012, '角色导出', 101, '0', 'system:role:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1013, '菜单查询', 102, '0', 'system:menu:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1014, '菜单新增', 102, '0', 'system:menu:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1015, '菜单修改', 102, '0', 'system:menu:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1016, '菜单删除', 102, '0', 'system:menu:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1017, '部门查询', 103, '0', 'system:dept:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1018, '部门新增', 103, '0', 'system:dept:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1019, '部门修改', 103, '0', 'system:dept:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1020, '部门删除', 103, '0', 'system:dept:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1021, '岗位查询', 104, '0', 'system:post:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1022, '岗位新增', 104, '0', 'system:post:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1023, '岗位修改', 104, '0', 'system:post:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1024, '岗位删除', 104, '0', 'system:post:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1025, '岗位导出', 104, '0', 'system:post:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1026, '字典查询', 105, '0', 'system:dict:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1027, '字典新增', 105, '0', 'system:dict:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1028, '字典修改', 105, '0', 'system:dict:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1029, '字典删除', 105, '0', 'system:dict:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1030, '字典导出', 105, '0', 'system:dict:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1043, '登录查询', 501, '0', 'monitor:logininfor:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1044, '登录删除', 501, '0', 'monitor:logininfor:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1045, '日志导出', 501, '0', 'monitor:logininfor:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1046, '在线查询', 109, '0', 'monitor:online:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1047, '批量强退', 109, '0', 'monitor:online:batchLogout', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1048, '单条强退', 109, '0', 'monitor:online:forceLogout', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1049, '任务查询', 110, '0', 'monitor:job:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1050, '任务新增', 110, '0', 'monitor:job:add', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1051, '任务修改', 110, '0', 'monitor:job:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1052, '任务删除', 110, '0', 'monitor:job:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1053, '状态修改', 110, '0', 'monitor:job:changeStatus', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1054, '任务导出', 110, '0', 'monitor:job:export', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1055, '生成查询', 115, '0', 'tool:gen:query', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1056, '生成修改', 115, '0', 'tool:gen:edit', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1057, '生成删除', 115, '0', 'tool:gen:remove', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1058, '导入代码', 115, '0', 'tool:gen:import', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1059, '预览代码', 115, '0', 'tool:gen:preview', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (1060, '生成代码', 115, '0', 'tool:gen:code', 1, '2021-08-15 12:02:06', 0, NULL, '');
INSERT INTO `sys_permission` VALUES (13501682332667904, 'a', 2, '0', '', 1, '2022-08-07 14:10:51', 1, '2022-08-07 14:12:23', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64200976684290048 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, '0', 1, '2021-08-15 12:02:05', 0, NULL, '');
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理', 2, '0', 1, '2021-08-15 12:02:05', 0, NULL, '');
INSERT INTO `sys_post` VALUES (3, 'hr', '人力资源', 3, '0', 1, '2021-08-15 12:02:05', 0, NULL, '');
INSERT INTO `sys_post` VALUES (4, 'user', '普通员工', 4, '0', 1, '2021-08-15 12:02:05', 0, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `permission_check_strictly` tinyint(1) DEFAULT '1' COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) DEFAULT '1' COMMENT '部门树选择项是否关联显示',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '0' COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12176427318128641 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'admin', 1, '1', 1, 1, '0', '0', 1, '2021-08-15 12:02:06', 1, '2021-10-30 09:44:44', '超级管理员');
INSERT INTO `sys_role` VALUES (2, '普通角色', 'common', 0, '2', 1, 1, '0', '0', 1, '2021-08-15 12:02:06', 1, '2022-08-03 22:24:35', '普通角色');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` VALUES (2, 100);
INSERT INTO `sys_role_dept` VALUES (2, 101);
INSERT INTO `sys_role_dept` VALUES (2, 105);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission` (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `permission_id` bigint NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint DEFAULT NULL COMMENT '部门ID',
  `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '密码',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `user_name` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63926280189382656 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 103, 'admin', '白泽', '00', 'ry@163.com', '15888888887', '0', '', '$2a$14$41kuso/yYISP68qmqTsqIOkITY1.wbS2QpuR.BUhyN6BvxRoquGw.', '0', '0', '127.0.0.1', '2022-08-07 15:33:19', 1, '2021-08-15 12:02:05', 1, '2022-07-28 22:53:53', '管理员');
INSERT INTO `sys_user` VALUES (2, 105, 'bz', '白泽', '00', 'ry@qq.com', '15666666666', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '0', '0', '127.0.0.1', '2022-08-03 22:25:15', 1, '2021-08-15 12:02:05', 1, '2021-11-14 15:55:36', '测试员');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` VALUES (1, 2);
INSERT INTO `sys_user_post` VALUES (2, 2);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (2, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
