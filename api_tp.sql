/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : api_tp

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-03-21 16:57:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tp_admin
-- ----------------------------
DROP TABLE IF EXISTS `tp_admin`;
CREATE TABLE `tp_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of tp_admin
-- ----------------------------
INSERT INTO `tp_admin` VALUES ('1', 'admin', '超级管理员', '2dc4a8162eb8ad6918603d37f1dd2c93', '0', '13888888889', 'admin2019@163.com', 'kmcB', '1553146786', '[::1]', '1', '0', '0', '1553146786', '1553146786');
INSERT INTO `tp_admin` VALUES ('2', 'ljj', '林俊杰', '2dc4a8162eb8ad6918603d37f1dd2c93', '1,2', '13811558899', 'ljj@163.com', 'ONNy', '1553146786', '127.0.0.1', '0', '0', '0', '1553146786', '1553146786');
INSERT INTO `tp_admin` VALUES ('3', 'sxf', '宋晓峰', '2dc4a8162eb8ad6918603d37f1dd2c93', '2,1', '13811559988', 'sxf@163.com', '6fWE', '1553146786', '127.0.0.1', '1', '1', '0', '1553146786', '1553146786');
INSERT INTO `tp_admin` VALUES ('4', 'zs', '赵四', '2dc4a8162eb8ad6918603d37f1dd2c93', '2', '13988009988', 'zs@163.com', 'i8Nf', '0', '', '1', '1', '0', '1553146786', '1553146786');

-- ----------------------------
-- Table structure for tp_group
-- ----------------------------
DROP TABLE IF EXISTS `tp_group`;
CREATE TABLE `tp_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) NOT NULL DEFAULT '' COMMENT '组名',
  `detail` varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态：1-正常，0-删除',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_id` (`create_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='一级分类表';

-- ----------------------------
-- Records of tp_group
-- ----------------------------
INSERT INTO `tp_group` VALUES ('1', '一级分类A', '独坐幽篁里，弹琴复长啸。', '1', '0', '0', '1553149848', '1553149848');
INSERT INTO `tp_group` VALUES ('2', '一级分类B', '深林人不知，明月来相照。', '1', '0', '0', '1553063448', '1553063448');
INSERT INTO `tp_group` VALUES ('3', '一级分类C', '银烛秋光冷画屏，轻罗小扇扑流萤。', '1', '0', '0', '1552977048', '1552977048');
INSERT INTO `tp_group` VALUES ('4', '一级分类D', '天阶夜色凉如水，卧看牵牛织女星。', '1', '0', '0', '1552890648', '1552890648');

-- ----------------------------
-- Table structure for tp_node
-- ----------------------------
DROP TABLE IF EXISTS `tp_node`;
CREATE TABLE `tp_node` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sub_group_id` int(11) NOT NULL DEFAULT '0' COMMENT '主表ID',
  `node_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '方法名称：1-元素 2-属性 3-文本 4-注释',
  `node_name` varchar(100) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `detail` text COMMENT '说明',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：0-已删除，1-禁用，2-启用',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`sub_group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='节点表';

-- ----------------------------
-- Records of tp_node
-- ----------------------------
INSERT INTO `tp_node` VALUES ('1', '1', '1', '咏柳', '参数配置', '1', '0', '0', '1553150547', '1553150547');
INSERT INTO `tp_node` VALUES ('2', '2', '2', '山行', '路由设置', '1', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('3', '3', '3', '终南望余雪', '控制器函数', '1', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('4', '4', '4', '望洞庭', 'XSRF 过滤', '1', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('5', '5', '1', '暮江吟', 'XSRF 过滤', '2', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('6', '6', '2', '夜雨寄北', 'session 控制', '2', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('7', '7', '3', '浣溪沙', '过滤器', '2', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('8', '8', '4', '登泰山记', 'flash 数据', '2', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('9', '1', '1', '终南山', 'URL构建', '3', '0', '0', '1553150555', '1553150555');
INSERT INTO `tp_node` VALUES ('10', '2', '2', '题西林壁', '多种格式数据输出', '3', '0', '0', '1555137048', '1555137048');
INSERT INTO `tp_node` VALUES ('11', '3', '3', '巫山高', '表单数据验证', '3', '0', '0', '1555137048', '1555137048');
INSERT INTO `tp_node` VALUES ('12', '4', '4', '南山诗', '错误处理', '3', '0', '0', '1555137048', '1555137048');
INSERT INTO `tp_node` VALUES ('13', '5', '1', '咏孤石', '日志处理', '4', '0', '0', '1555741848', '1555741848');
INSERT INTO `tp_node` VALUES ('14', '6', '2', '题君山', 'ORM 使用', '4', '0', '0', '1555741848', '1555741848');
INSERT INTO `tp_node` VALUES ('15', '7', '3', '沁园春·雪', 'CRUD 操作', '4', '0', '0', '1555741848', '1555741848');

-- ----------------------------
-- Table structure for tp_sub_group
-- ----------------------------
DROP TABLE IF EXISTS `tp_sub_group`;
CREATE TABLE `tp_sub_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '组ID',
  `sub_group_name` varchar(50) NOT NULL DEFAULT '0' COMMENT '二级分类名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-删除',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='二级分类表';

-- ----------------------------
-- Records of tp_sub_group
-- ----------------------------
INSERT INTO `tp_sub_group` VALUES ('1', '1', '二级分类A', '1', '0', '0', '1552890648', '1552890648');
INSERT INTO `tp_sub_group` VALUES ('2', '1', '二级分类B', '1', '0', '0', '1552372248', '1552372248');
INSERT INTO `tp_sub_group` VALUES ('3', '2', '二级分类C', '1', '0', '0', '1552372248', '1552372248');
INSERT INTO `tp_sub_group` VALUES ('4', '2', '二级分类D', '1', '0', '0', '1552372248', '1552372248');
INSERT INTO `tp_sub_group` VALUES ('5', '3', '二级分类E', '1', '0', '0', '1557729048', '1557729048');
INSERT INTO `tp_sub_group` VALUES ('6', '3', '二级分类F', '1', '0', '0', '1557729048', '1557729048');
INSERT INTO `tp_sub_group` VALUES ('7', '4', '二级分类G', '1', '0', '0', '1562999448', '1562999448');
INSERT INTO `tp_sub_group` VALUES ('8', '4', '二级分类H', '1', '0', '0', '1565677848', '1565677848');
