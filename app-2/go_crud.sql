/*
 Navicat Premium Data Transfer

 Source Server         : go_crud
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : go_crud

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 18/01/2024 20:23:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES (11, 'Clothes', '2024-01-15 15:32:39', '2024-01-16 12:42:40');
INSERT INTO `categories` VALUES (13, 'Shoes', '2024-01-15 16:34:23', '2024-01-16 12:42:46');
INSERT INTO `categories` VALUES (14, 'Electronic', '2024-01-16 12:42:54', '2024-01-16 16:37:53');

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `category_id` int NOT NULL,
  `stock` int NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (7, 'sabrina', 7, 1, 'asdfasdfa', '2024-01-12 14:05:46', '2024-01-12 14:05:46');
INSERT INTO `products` VALUES (8, 'Product1', 1, 50, 'Description for Product1', '2024-01-16 00:59:22', '2024-01-16 00:59:22');
INSERT INTO `products` VALUES (10, 'Coat', 11, 50, 'Winter Clothes', '2024-01-16 12:53:13', '2024-01-16 16:38:18');
INSERT INTO `products` VALUES (11, 'SAMSUNG', 14, 2, 'Handphone', '2024-01-16 12:54:54', '2024-01-16 16:38:43');
INSERT INTO `products` VALUES (12, 'ADIDAS', 13, 10, 'Sepatu Mesi', '2024-01-16 13:00:49', '2024-01-16 15:02:23');

SET FOREIGN_KEY_CHECKS = 1;
