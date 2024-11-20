/*
 Navicat Premium Data Transfer

 Source Server         : MYSQL
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : musicplayer

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 20/11/2024 18:49:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for album_info
-- ----------------------------
DROP TABLE IF EXISTS `album_info`;
CREATE TABLE `album_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `release_date` date NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of album_info
-- ----------------------------

-- ----------------------------
-- Table structure for artist_album_relation
-- ----------------------------
DROP TABLE IF EXISTS `artist_album_relation`;
CREATE TABLE `artist_album_relation`  (
  `artist_id` int(11) NOT NULL,
  `album_id` int(11) NOT NULL,
  PRIMARY KEY (`artist_id`, `album_id`) USING BTREE,
  INDEX `aa_fk2`(`album_id` ASC) USING BTREE,
  CONSTRAINT `aa_fk1` FOREIGN KEY (`artist_id`) REFERENCES `artist_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `aa_fk2` FOREIGN KEY (`album_id`) REFERENCES `album_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of artist_album_relation
-- ----------------------------

-- ----------------------------
-- Table structure for artist_info
-- ----------------------------
DROP TABLE IF EXISTS `artist_info`;
CREATE TABLE `artist_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `profile_pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nation` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of artist_info
-- ----------------------------
INSERT INTO `artist_info` VALUES (1, '周杰伦', '亚洲天王', '999', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (2, '贝多芬', '音乐之神', '2222', '古典作曲家', '德国');
INSERT INTO `artist_info` VALUES (4, 'Mashroom', '百大DJ', '4444', 'DJ', '欧美');

-- ----------------------------
-- Table structure for artist_song_relation
-- ----------------------------
DROP TABLE IF EXISTS `artist_song_relation`;
CREATE TABLE `artist_song_relation`  (
  `artist_id` int(11) NOT NULL,
  `song_id` int(11) NOT NULL,
  PRIMARY KEY (`artist_id`, `song_id`) USING BTREE,
  INDEX `as_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `as_fk1` FOREIGN KEY (`artist_id`) REFERENCES `artist_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `as_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of artist_song_relation
-- ----------------------------

-- ----------------------------
-- Table structure for comment_info
-- ----------------------------
DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `target_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `comment_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment_info
-- ----------------------------

-- ----------------------------
-- Table structure for download_info
-- ----------------------------
DROP TABLE IF EXISTS `download_info`;
CREATE TABLE `download_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `song_id` int(11) NULL DEFAULT NULL,
  `download_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `download_fk1`(`user_id` ASC) USING BTREE,
  INDEX `download_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `download_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `download_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of download_info
-- ----------------------------

-- ----------------------------
-- Table structure for follow_artist
-- ----------------------------
DROP TABLE IF EXISTS `follow_artist`;
CREATE TABLE `follow_artist`  (
  `follower_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `followed_id` int(11) NOT NULL,
  PRIMARY KEY (`follower_id`, `followed_id`) USING BTREE,
  INDEX `followA_fk2`(`followed_id` ASC) USING BTREE,
  CONSTRAINT `followA_fk1` FOREIGN KEY (`follower_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `followA_fk2` FOREIGN KEY (`followed_id`) REFERENCES `artist_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of follow_artist
-- ----------------------------
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 1);
INSERT INTO `follow_artist` VALUES ('e9nRUN7ZRB6pDw6', 1);

-- ----------------------------
-- Table structure for follow_user
-- ----------------------------
DROP TABLE IF EXISTS `follow_user`;
CREATE TABLE `follow_user`  (
  `follower_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `followed_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '被关注用户',
  PRIMARY KEY (`follower_id`, `followed_id`) USING BTREE,
  INDEX `follow_fk2`(`followed_id` ASC) USING BTREE,
  CONSTRAINT `follow_fk1` FOREIGN KEY (`follower_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `follow_fk2` FOREIGN KEY (`followed_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of follow_user
-- ----------------------------
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', '9EyQGH08RTmYc40');
INSERT INTO `follow_user` VALUES ('e9nRUN7ZRB6pDw6', '9EyQGH08RTmYc40');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'e9nRUN7ZRB6pDw6');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'ZRQ6M-UcS2yedwY');

-- ----------------------------
-- Table structure for like_info
-- ----------------------------
DROP TABLE IF EXISTS `like_info`;
CREATE TABLE `like_info`  (
  `moment_id` int(11) NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`moment_id`, `user_id`) USING BTREE,
  INDEX `like_fk2`(`user_id` ASC) USING BTREE,
  CONSTRAINT `like_fk1` FOREIGN KEY (`moment_id`) REFERENCES `moment_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `like_fk2` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of like_info
-- ----------------------------

-- ----------------------------
-- Table structure for local_songlist
-- ----------------------------
DROP TABLE IF EXISTS `local_songlist`;
CREATE TABLE `local_songlist`  (
  `song_id` int(11) NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `added_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`song_id`, `user_id`) USING BTREE,
  INDEX `local_fk2`(`user_id` ASC) USING BTREE,
  CONSTRAINT `local_fk1` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `local_fk2` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of local_songlist
-- ----------------------------

-- ----------------------------
-- Table structure for message_info
-- ----------------------------
DROP TABLE IF EXISTS `message_info`;
CREATE TABLE `message_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `sender_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `receiver_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `message_fk1`(`sender_id` ASC) USING BTREE,
  INDEX `message_fk2`(`receiver_id` ASC) USING BTREE,
  CONSTRAINT `message_fk1` FOREIGN KEY (`sender_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `message_fk2` FOREIGN KEY (`receiver_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of message_info
-- ----------------------------

-- ----------------------------
-- Table structure for moment_info
-- ----------------------------
DROP TABLE IF EXISTS `moment_info`;
CREATE TABLE `moment_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `moment_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `moment_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of moment_info
-- ----------------------------
INSERT INTO `moment_info` VALUES (2, '-123-ni hao a', '2024-11-13 08:44:31', '9EyQGH08RTmYc40', '0000');
INSERT INTO `moment_info` VALUES (4, 'ni hao a', '2024-11-13 08:55:29', 'e9nRUN7ZRB6pDw6', 'hahah');
INSERT INTO `moment_info` VALUES (5, '11111', '2024-11-13 09:28:17', 'ZRQ6M-UcS2yedwY', '0000');
INSERT INTO `moment_info` VALUES (6, '22222', '2024-11-13 09:28:42', 'ZRQ6M-UcS2yedwY', '0');
INSERT INTO `moment_info` VALUES (8, '44444', '2024-11-13 09:30:58', 'ZRQ6M-UcS2yedwY', '0');
INSERT INTO `moment_info` VALUES (9, '44444', '2024-11-13 10:48:25', 'ZRQ6M-UcS2yedwY', '0');

-- ----------------------------
-- Table structure for playlist_info
-- ----------------------------
DROP TABLE IF EXISTS `playlist_info`;
CREATE TABLE `playlist_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `hits` int(11) NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `playlist_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `playlist_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of playlist_info
-- ----------------------------

-- ----------------------------
-- Table structure for ranking_info
-- ----------------------------
DROP TABLE IF EXISTS `ranking_info`;
CREATE TABLE `ranking_info`  (
  `song_id` int(11) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `rank` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`song_id`, `name`) USING BTREE,
  CONSTRAINT `ranking_fk1` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ranking_info
-- ----------------------------

-- ----------------------------
-- Table structure for setting_info
-- ----------------------------
DROP TABLE IF EXISTS `setting_info`;
CREATE TABLE `setting_info`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  CONSTRAINT `setting_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of setting_info
-- ----------------------------

-- ----------------------------
-- Table structure for song_info
-- ----------------------------
DROP TABLE IF EXISTS `song_info`;
CREATE TABLE `song_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `duration` int(11) NULL DEFAULT NULL,
  `album_id` int(11) NULL DEFAULT NULL,
  `genre` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `release_date` date NULL DEFAULT NULL,
  `song_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `lyrics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `song_hit` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of song_info
-- ----------------------------

-- ----------------------------
-- Table structure for song_playlist_relation
-- ----------------------------
DROP TABLE IF EXISTS `song_playlist_relation`;
CREATE TABLE `song_playlist_relation`  (
  `playlist_id` int(11) NOT NULL,
  `song_id` int(11) NOT NULL,
  PRIMARY KEY (`playlist_id`, `song_id`) USING BTREE,
  INDEX `sp_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `sp_fk1` FOREIGN KEY (`playlist_id`) REFERENCES `playlist_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `sp_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of song_playlist_relation
-- ----------------------------

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号ID\r\n',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名（昵称）',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `country` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `region` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `gender` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `profile_pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES ('9EyQGH08RTmYc40', '劳国琴', '123456', 'k8qlqi.pv5@qq.com', '14272676013', '2024-11-13 07:47:59', 'repreh', 'esse mollit', '女', 'laborum voluptate irure cillum veniam', 'https://via.placeholder.com/400x400/671b7b/1d121f.gif', '2024-11-13 08:03:06');
INSERT INTO `user_info` VALUES ('CrMhKMGnQH6R-Vq', 'dj05', '$2a$10$xL8TTTC6EXuXyHytmRc6v.fjiF3StFDwRbFmOTNVqHzLsORpHsd2u', 'dj05@qq.com', '', '2024-11-13 07:48:03', '', '', '', '', '', '2024-11-13 07:48:03');
INSERT INTO `user_info` VALUES ('e9nRUN7ZRB6pDw6', 'dj02', '$2a$10$AEQaSl6DgBW131bxh5XPfek6QXGR2N7GlV8j0N.s7fnJNzgv2xb4e', 'dj02@qq.com', '', '2024-11-13 07:47:43', '', '', '', '', '', '2024-11-13 07:47:43');
INSERT INTO `user_info` VALUES ('tgy6oWzjSYCjgOA', 'dj03', '$2a$10$1JcBEF.dmGEe64G2orpFUeY4wO876EpxNtIrbAFEx/5KCGfbNezfy', 'dj03@qq.com', '', '2024-11-13 07:47:54', '', '', '', '', '', '2024-11-13 07:47:54');
INSERT INTO `user_info` VALUES ('ZRQ6M-UcS2yedwY', 'dj01', '$2a$10$JtynMIb9xQq/lax.kQMgOOKEPG6j423wgIb9DmDoUYoAK7ZNZX3Ta', 'dj01@qq.com', '', '2024-11-13 07:47:12', '', '', '', '', '', '2024-11-13 07:47:12');

-- ----------------------------
-- Table structure for user_like_playlist
-- ----------------------------
DROP TABLE IF EXISTS `user_like_playlist`;
CREATE TABLE `user_like_playlist`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `playlist_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`, `playlist_id`) USING BTREE,
  INDEX `up_fk2`(`playlist_id` ASC) USING BTREE,
  CONSTRAINT `up_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `up_fk2` FOREIGN KEY (`playlist_id`) REFERENCES `playlist_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_like_playlist
-- ----------------------------

-- ----------------------------
-- Table structure for user_like_song
-- ----------------------------
DROP TABLE IF EXISTS `user_like_song`;
CREATE TABLE `user_like_song`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `song_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`, `song_id`) USING BTREE,
  INDEX `us_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `us_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `us_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_like_song
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
