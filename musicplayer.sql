ALTER USER `root`@`localhost` IDENTIFIED BY 'Hxy1234]';
/*
 Navicat Premium Dump SQL

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : localhost:3306
 Source Schema         : musicplayer

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 07/12/2024 22:42:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for album_info
-- ----------------------------
DROP TABLE IF EXISTS `album_info`;
CREATE TABLE `album_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
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
  `artist_id` int NOT NULL,
  `album_id` int NOT NULL,
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
  `id` int NOT NULL AUTO_INCREMENT,
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
  `artist_id` int NOT NULL,
  `song_id` int NOT NULL,
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
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `target_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `comment_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment_info
-- ----------------------------
INSERT INTO `comment_info` VALUES (8, 'aaaa nice', '2024-11-24 18:27:34', 'CrMhKMGnQH6R-Vq', 'song', 3);
INSERT INTO `comment_info` VALUES (9, 'asdfasdf', '2024-11-24 18:28:01', 'CrMhKMGnQH6R-Vq', 'song', 3);
INSERT INTO `comment_info` VALUES (10, 'asdfasdf', '2024-11-24 18:28:13', 'CrMhKMGnQH6R-Vq', 'song', 3);
INSERT INTO `comment_info` VALUES (11, 'zxcasd   sdd', '2024-11-24 18:28:42', 'e9nRUN7ZRB6pDw6', 'song', 3);
INSERT INTO `comment_info` VALUES (12, 'zxcasd   sdd', '2024-11-24 18:28:45', 'e9nRUN7ZRB6pDw6', 'song', 3);
INSERT INTO `comment_info` VALUES (13, 'zxcasd   sdd', '2024-11-24 18:33:45', 'e9nRUN7ZRB6pDw6', 'song', 3);
INSERT INTO `comment_info` VALUES (14, 'zxcasd   sdd', '2024-11-24 18:34:00', 'e9nRUN7ZRB6pDw6', 'song', 3);
INSERT INTO `comment_info` VALUES (15, 'zxcasd   sdd', '2024-11-24 18:34:25', 'e9nRUN7ZRB6pDw6', 'song', 3);
INSERT INTO `comment_info` VALUES (18, 'cupidatat', '2024-12-07 08:31:05', 'e9nRUN7ZRB6pDw6', 'moment', 20);
INSERT INTO `comment_info` VALUES (19, 'cupidatat', '2024-12-07 08:32:32', 'e9nRUN7ZRB6pDw6', 'moment', 20);
INSERT INTO `comment_info` VALUES (20, 'enim velit', '2024-12-07 09:07:15', 'e9nRUN7ZRB6pDw6', 'moment', 20);
INSERT INTO `comment_info` VALUES (21, 'commodo', '2024-12-07 09:07:19', 'e9nRUN7ZRB6pDw6', 'moment', 20);
INSERT INTO `comment_info` VALUES (22, 'nulla qui reprehenderit', '2024-12-07 09:07:21', 'e9nRUN7ZRB6pDw6', 'moment', 20);
INSERT INTO `comment_info` VALUES (24, 'tempor voluptate minim pariatur laborum', '2024-12-07 12:55:54', 'e9nRUN7ZRB6pDw6', 'moment', 20);

-- ----------------------------
-- Table structure for download_info
-- ----------------------------
DROP TABLE IF EXISTS `download_info`;
CREATE TABLE `download_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `song_id` int NULL DEFAULT NULL,
  `download_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `download_fk1`(`user_id` ASC) USING BTREE,
  INDEX `download_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `download_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `download_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of download_info
-- ----------------------------

-- ----------------------------
-- Table structure for follow_artist
-- ----------------------------
DROP TABLE IF EXISTS `follow_artist`;
CREATE TABLE `follow_artist`  (
  `follower_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `followed_id` int NOT NULL,
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
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'e9nRUN7ZRB6pDw6');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'ZRQ6M-UcS2yedwY');

-- ----------------------------
-- Table structure for like_comment
-- ----------------------------
DROP TABLE IF EXISTS `like_comment`;
CREATE TABLE `like_comment`  (
  `comment_id` int NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`comment_id`, `user_id`) USING BTREE,
  INDEX `like_c_fk2`(`user_id` ASC) USING BTREE,
  CONSTRAINT `like_c_fk1` FOREIGN KEY (`comment_id`) REFERENCES `comment_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `like_c_fk2` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of like_comment
-- ----------------------------
INSERT INTO `like_comment` VALUES (20, 'e9nRUN7ZRB6pDw6');

-- ----------------------------
-- Table structure for like_info
-- ----------------------------
DROP TABLE IF EXISTS `like_info`;
CREATE TABLE `like_info`  (
  `moment_id` int NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`moment_id`, `user_id`) USING BTREE,
  INDEX `like_fk2`(`user_id` ASC) USING BTREE,
  CONSTRAINT `like_fk1` FOREIGN KEY (`moment_id`) REFERENCES `moment_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `like_fk2` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of like_info
-- ----------------------------
INSERT INTO `like_info` VALUES (20, 'e9nRUN7ZRB6pDw6');
INSERT INTO `like_info` VALUES (21, 'e9nRUN7ZRB6pDw6');

-- ----------------------------
-- Table structure for local_songlist
-- ----------------------------
DROP TABLE IF EXISTS `local_songlist`;
CREATE TABLE `local_songlist`  (
  `song_id` int NOT NULL,
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
  `id` int NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of message_info
-- ----------------------------

-- ----------------------------
-- Table structure for moment_info
-- ----------------------------
DROP TABLE IF EXISTS `moment_info`;
CREATE TABLE `moment_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `moment_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `moment_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of moment_info
-- ----------------------------
INSERT INTO `moment_info` VALUES (6, '22222', '2024-11-13 09:28:42', 'ZRQ6M-UcS2yedwY', '0');
INSERT INTO `moment_info` VALUES (8, '44444', '2024-11-13 09:30:58', 'ZRQ6M-UcS2yedwY', '0');
INSERT INTO `moment_info` VALUES (9, '44444', '2024-11-13 10:48:25', 'ZRQ6M-UcS2yedwY', '0');
INSERT INTO `moment_info` VALUES (11, 'anim culpa voluptate Duis reprehenderit', '2024-11-25 14:49:14', 'gPF1ZjgCTJSjqhU', 'https://loremflickr.com/400/400?lock=541072719010564');
INSERT INTO `moment_info` VALUES (12, 'cillum ut sint nulla sunt', '2024-11-25 14:49:40', 'gPF1ZjgCTJSjqhU', 'https://loremflickr.com/400/400?lock=1630037935104163');
INSERT INTO `moment_info` VALUES (14, 'veniam eu Lorem', '2024-11-25 15:14:05', 'gPF1ZjgCTJSjqhU', 'https://loremflickr.com/400/400?lock=8933374674082770');
INSERT INTO `moment_info` VALUES (15, 'Lorem amet consectetur eu dolor', '2024-11-25 15:14:46', 'CrMhKMGnQH6R-Vq', 'https://loremflickr.com/400/400?lock=3910359776945596');
INSERT INTO `moment_info` VALUES (18, 'aliquip amet', '2024-12-07 05:44:43', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (19, 'aliquip amet', '2024-12-07 05:46:06', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (20, 'aliquip amet', '2024-12-07 05:46:37', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (21, 'aliquip amet', '2024-12-07 06:01:28', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (22, 'aliquip amet', '2024-12-07 06:05:10', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (23, 'aliquip amet', '2024-12-07 12:59:22', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');

-- ----------------------------
-- Table structure for playlist_info
-- ----------------------------
DROP TABLE IF EXISTS `playlist_info`;
CREATE TABLE `playlist_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `hits` int NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `playlist_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `playlist_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of playlist_info
-- ----------------------------

-- ----------------------------
-- Table structure for ranking_info
-- ----------------------------
DROP TABLE IF EXISTS `ranking_info`;
CREATE TABLE `ranking_info`  (
  `song_id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `rank` int NULL DEFAULT NULL,
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
  `msg` int NULL DEFAULT 0 COMMENT '0：所有人；1：关注的人',
  `see_rank` int NULL DEFAULT 0 COMMENT '0：所有人可见；1：仅自己可见',
  `info_comment` int NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_like` int NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_msg` int NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_sys` int NULL DEFAULT 0 COMMENT '0：假；1：真',
  `service` int NULL DEFAULT 0 COMMENT '0：假；1：真',
  PRIMARY KEY (`user_id`) USING BTREE,
  CONSTRAINT `setting_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of setting_info
-- ----------------------------
INSERT INTO `setting_info` VALUES ('CrMhKMGnQH6R-Vq', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `setting_info` VALUES ('e9nRUN7ZRB6pDw6', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `setting_info` VALUES ('gPF1ZjgCTJSjqhU', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `setting_info` VALUES ('OMgiLMEDTxCSSB1', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `setting_info` VALUES ('tgy6oWzjSYCjgOA', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `setting_info` VALUES ('ZRQ6M-UcS2yedwY', 0, 0, 1, 0, 0, 1, 0);

-- ----------------------------
-- Table structure for song_info
-- ----------------------------
DROP TABLE IF EXISTS `song_info`;
CREATE TABLE `song_info`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `duration` int NULL DEFAULT NULL,
  `album_id` int NULL DEFAULT NULL,
  `genre` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `release_date` date NULL DEFAULT NULL,
  `song_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `lyrics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `song_hit` int NULL DEFAULT 0,
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
  `playlist_id` int NOT NULL,
  `song_id` int NOT NULL,
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
INSERT INTO `user_info` VALUES ('CrMhKMGnQH6R-Vq', 'dj05', '$2a$10$xL8TTTC6EXuXyHytmRc6v.fjiF3StFDwRbFmOTNVqHzLsORpHsd2u', 'dj05@qq.com', '', '2024-11-13 07:48:03', '', '', '', '', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2024-12-06 17:30:17');
INSERT INTO `user_info` VALUES ('e9nRUN7ZRB6pDw6', 'dj02', '$2a$10$vhsy/vbE/uA.3celgfTu7urxMjAcs1BorkTL8Lrs70frnMkPt8MnO', 'dj02@qq.com', '', '2024-11-13 07:47:43', '', '', '', '', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2024-12-07 14:14:06');
INSERT INTO `user_info` VALUES ('gPF1ZjgCTJSjqhU', '禄霞', 'Rphttk_gAEq92Xr', 'ppm_lfn86@vip.qq.com', '42794543813', '2024-11-25 14:10:13', 'cu', '东北', '男', '战斗粉丝', 'https://loremflickr.com/400/400?lock=1075646362105748', '2024-11-25 14:18:42');
INSERT INTO `user_info` VALUES ('OMgiLMEDTxCSSB1', 'dj01', '$2a$10$YDzhOkP4/mtrS6yWgjEvROZYMia6RSx2jF95jEJCgTjvQCAdaD5OO', '1796654305@qq.com', '', '2024-12-07 13:59:12', '', '', '', '', '', '2024-12-07 13:59:12');
INSERT INTO `user_info` VALUES ('tgy6oWzjSYCjgOA', 'dj03', '$2a$10$1JcBEF.dmGEe64G2orpFUeY4wO876EpxNtIrbAFEx/5KCGfbNezfy', 'dj03@qq.com', '', '2024-11-13 07:47:54', '', '', '', '', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2024-12-06 17:30:19');
INSERT INTO `user_info` VALUES ('ZRQ6M-UcS2yedwY', '磨文昊', '$2a$10$imLcf8YOIQ9/rt8aA.mXLO1eIqao2.kfcd4atB802b.6f2cx7BQNu', 'mawnuy_i4n52@foxmail.com', '99855657488', '2024-11-13 07:47:12', 'non', '东北', '女', '航空邮件倡导者，梦想家', 'https://loremflickr.com/400/400?lock=1157922648024234', '2024-12-07 13:04:28');

-- ----------------------------
-- Table structure for user_like_playlist
-- ----------------------------
DROP TABLE IF EXISTS `user_like_playlist`;
CREATE TABLE `user_like_playlist`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `playlist_id` int NOT NULL,
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
  `song_id` int NOT NULL,
  PRIMARY KEY (`user_id`, `song_id`) USING BTREE,
  INDEX `us_fk2`(`song_id` ASC) USING BTREE,
  CONSTRAINT `us_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `us_fk2` FOREIGN KEY (`song_id`) REFERENCES `song_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_like_song
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
