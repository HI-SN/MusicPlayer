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

 Date: 15/01/2025 04:35:09
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
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of album_info
-- ----------------------------
INSERT INTO `album_info` VALUES (1, '男女情歌对唱冠军全记录', '男女情歌对唱冠军全记录\r\n最强的对唱阵容，让你一次拥有……关于暗恋的，失恋的，热恋的，最感动的男女情歌对唱', '2007-02-05', 'http://p1.music.126.net/81BsxxhomJ4aJZYvEbyPkw==/109951165671182684.jpg?param=177y177');
INSERT INTO `album_info` VALUES (2, '我想念', '原创泷式情歌 乐迷心中的「宝藏」', '2023-07-22', 'http://p1.music.126.net/MIWMnBEedpJ8IpOau5D7PA==/109951168829804653.jpg?param=177y177');
INSERT INTO `album_info` VALUES (3, '好安静', '这张是汪苏泷未收录单曲合辑，集结《好安静》、《我也不知道》、《累不累》、《不分手的恋爱》、《苦笑》、《三国杀》、《某人》等七首原创作品，皆由汪苏泷一人包办词曲。', '2011-09-09', 'http://p1.music.126.net/0WUciQeFa8l6LtNgfU-Mqg==/18537766044710327.jpg?param=177y177');
INSERT INTO `album_info` VALUES (4, '十万伏特', '暗夜深处涌动的火光，最能叫人目眩神迷\r\n冰冷机器迸发的电流，全新世界一触即发\r\n钢铁和心脏被闪电一同击穿\r\n汪苏泷第11张创作专辑《十万伏特》\r\n11首贯彻全身的音乐能量\r\n你是巨大的，你是自由的', '2024-12-08', 'http://p1.music.126.net/89vcG85t_XaaHj5oc0iaag==/109951170223583930.jpg?param=177y177');
INSERT INTO `album_info` VALUES (5, 'One Step Closer', '周柏豪新专辑《One Step Closer》于4月3日正式发行。收录作品包括《有生一天》、《近在千里》、《嚣》、十周年演唱会点题作《终于我们 (One Step Closer)》在内的7首单曲。', '2017-04-03', 'http://p1.music.126.net/O4E8E34ocUId-TKLrXH1UA==/18293674463464066.jpg?param=177y177');
INSERT INTO `album_info` VALUES (6, '少年般绚丽', '少年如盛开的花，如绚丽的光！不仅耀眼夺目，更是充满无限美的可能！', '2022-12-31', 'http://p2.music.126.net/LljI1W-pKVh5tS-sapFxeA==/109951168182630884.jpg?param=177y177');
INSERT INTO `album_info` VALUES (7, '21世纪罗曼史', '时代呼啸着、翻滚着、飞驰着\r\n我们沉默着、热恋着、舞蹈着\r\n汪苏泷第10张个人创作专辑《21世纪罗曼史》\r\n写给生活在这个世纪里的每一个人\r\n记录我们的命运与爱', '2022-10-28', 'http://p2.music.126.net/-0uzd-3roA8rZP7NPlLr5w==/109951168311222481.jpg?param=177y177');
INSERT INTO `album_info` VALUES (8, '时间', '你所不知道的周慧敏 从眼里到心底从表演到创作\r\n一九九六年全新国语大碟 首推力荐周慧敏倾心创作曲', '1996-10-08', 'http://p1.music.126.net/ypvp8edSHz2rUsQVnmcOJg==/109951170166456452.jpg?param=177y177');
INSERT INTO `album_info` VALUES (9, '飞行器的执行周期', '如果你有一个即将返程的飞行器，你期待它会带回什么故事？\r\n\r\n比如你想去外面，看世界。飞行器是你的载体，穿越星云，安静地，飞速前进⋯…\r\n对“未来\"的假想、“人工智能”的臆测一点点累积，最终命名为《飞行器的执行周期》。', '2016-11-25', 'http://p1.music.126.net/wSMfGvFzOAYRU_yVIfquAA==/2946691248081599.jpg?param=177y177');
INSERT INTO `album_info` VALUES (10, 'V', '【V】是罗马数字的「五」，也是代表胜利Victory的字母“V”，这张名为【V】的全新大碟，不仅是Maroon “5”生涯的第「五」张专辑，更见证了他们在流行乐坛的空前胜利！', '2014-08-29', 'http://p1.music.126.net/RkbAVgGxk1Nr5fjuaR9dww==/19175482788569403.jpg?param=177y177');
INSERT INTO `album_info` VALUES (11, 'Overexposed (Deluxe)', '\r\n【Overexposed】是历年专辑中最像也是最不像魔力红的作品，因为当中有他们一路走来所留下的精彩轨迹，同时又蕴含了前所未有的全新创意。', '2012-06-20', 'http://p1.music.126.net/zhb4NhgP262N24X7RmQBGg==/3222668584137511.jpg?param=177y177');
INSERT INTO `album_info` VALUES (12, 'The Best Ever Piano Classics', 'The Best Ever Piano Classics', '1999-03-16', 'http://p1.music.126.net/73-aaCvFEMzYWhk8sXxVvQ==/1729531790501554.jpg?param=177y177');
INSERT INTO `album_info` VALUES (13, 'Alone', '《Alone》是美国DJ和电音制作人Marshmello的一首热单，同时也是猫厂Monstercat的一首获得RIAA美国唱片工业协会销量认证的单曲。', '2016-05-13', 'http://p2.music.126.net/4hbR27M-uaKZVAOUojERTA==/109951168003842438.jpg?param=177y177');
INSERT INTO `album_info` VALUES (14, 'WALK - The 6th Album', '全世界等待的“K-POP英雄”回归！', '2024-07-15', 'https://p2.music.126.net/a2LTtXceOOwh7flZqr-4Mw==/109951169866453987.jpg?param=177y177');
INSERT INTO `album_info` VALUES (15, '‘The ReVe Festival’ Finale', 'Red Velvet，2019年音乐庆典华丽最终章！', '2019-12-23', 'https://p2.music.126.net/p3m7nswR_S2VjHqu71kKxg==/109951167760346730.jpg?param=177y177');
INSERT INTO `album_info` VALUES (16, 'Miracles In December', 'EXO于12月9日发布冬季特别专辑《12月的奇迹》，并计划展开温暖的特别舞台。', '2013-12-09', 'https://p2.music.126.net/Glv3YRUt7wh2lc339Ykd-g==/18527870440277729.jpg?param=177y177');

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
INSERT INTO `artist_album_relation` VALUES (1, 1);
INSERT INTO `artist_album_relation` VALUES (4, 2);
INSERT INTO `artist_album_relation` VALUES (4, 3);
INSERT INTO `artist_album_relation` VALUES (4, 4);
INSERT INTO `artist_album_relation` VALUES (5, 5);
INSERT INTO `artist_album_relation` VALUES (4, 6);
INSERT INTO `artist_album_relation` VALUES (4, 7);
INSERT INTO `artist_album_relation` VALUES (6, 8);
INSERT INTO `artist_album_relation` VALUES (7, 9);
INSERT INTO `artist_album_relation` VALUES (8, 10);
INSERT INTO `artist_album_relation` VALUES (8, 11);
INSERT INTO `artist_album_relation` VALUES (2, 12);
INSERT INTO `artist_album_relation` VALUES (3, 13);
INSERT INTO `artist_album_relation` VALUES (9, 14);
INSERT INTO `artist_album_relation` VALUES (10, 15);
INSERT INTO `artist_album_relation` VALUES (11, 16);

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
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of artist_info
-- ----------------------------
INSERT INTO `artist_info` VALUES (1, '周杰伦', '亚洲天王', 'http://p1.music.126.net/NWv6PtSBkyWZzqbJVzBr7g==/109951169164936450.jpg?param=640y300', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (2, '贝多芬', '音乐之神', 'http://p1.music.126.net/iEBu3AGqh4C-3wW-_lT8JQ==/109951165867914165.jpg?param=640y300', '古典作曲家', '德国');
INSERT INTO `artist_info` VALUES (3, 'Marshmello', '百大DJ', 'http://p2.music.126.net/lhxiXf_sdSU7ogG7FNpQ7A==/109951164967534612.jpg?param=640y300', 'DJ', '欧美');
INSERT INTO `artist_info` VALUES (4, '汪苏泷', '1', 'http://p2.music.126.net/7G5HqyqcpZoP4cHL7-a-hQ==/109951170027064713.jpg?param=640y300', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (5, '周柏豪', '1', 'http://p1.music.126.net/ETv-9VNsaaW2xE5NDcUU3w==/109951169253764017.jpg?param=640y300', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (6, '周慧敏', '1', 'http://p1.music.126.net/AyMpC-0FrTVq--3ZN21HkQ==/109951168314010164.jpg?param=640y300', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (7, '郭顶', '1', 'http://p1.music.126.net/7OoAuH2Iqjr3Owmqf4pNFQ==/109951165912271970.jpg?param=640y300', '流行歌手', '中国');
INSERT INTO `artist_info` VALUES (8, 'Maroon 5', '1', 'http://p1.music.126.net/L6Vf-GYbpOroZbQo_yxFzg==/109951169421839676.jpg?param=640y300', '乐队', '欧美');
INSERT INTO `artist_info` VALUES (9, 'NCT127', '划数', '	https://p2.music.126.net/iSYZmktYfHWuoWbjJj56fw==/109951169913665206.jpg?param=130y130', '男团', '韩国');
INSERT INTO `artist_info` VALUES (10, 'Red Velvet', '莱德贝贝', 'https://p2.music.126.net/_orn5sSfaUVOXsYqkul1Ow==/109951170290688256.jpg?param=130y130', '女团', '韩国');
INSERT INTO `artist_info` VALUES (11, 'EXO', '地', 'https://p1.music.126.net/13uIhYsC21O2fF76gYwF_A==/109951168721802905.jpg?param=130y130', '男团', '韩国');

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
INSERT INTO `artist_song_relation` VALUES (1, 1);
INSERT INTO `artist_song_relation` VALUES (4, 2);
INSERT INTO `artist_song_relation` VALUES (4, 3);
INSERT INTO `artist_song_relation` VALUES (4, 4);
INSERT INTO `artist_song_relation` VALUES (4, 5);
INSERT INTO `artist_song_relation` VALUES (4, 6);
INSERT INTO `artist_song_relation` VALUES (4, 7);
INSERT INTO `artist_song_relation` VALUES (4, 8);
INSERT INTO `artist_song_relation` VALUES (5, 9);
INSERT INTO `artist_song_relation` VALUES (5, 10);
INSERT INTO `artist_song_relation` VALUES (4, 11);
INSERT INTO `artist_song_relation` VALUES (4, 12);
INSERT INTO `artist_song_relation` VALUES (4, 13);
INSERT INTO `artist_song_relation` VALUES (4, 14);
INSERT INTO `artist_song_relation` VALUES (6, 15);
INSERT INTO `artist_song_relation` VALUES (7, 16);
INSERT INTO `artist_song_relation` VALUES (7, 17);
INSERT INTO `artist_song_relation` VALUES (8, 18);
INSERT INTO `artist_song_relation` VALUES (8, 19);
INSERT INTO `artist_song_relation` VALUES (2, 20);
INSERT INTO `artist_song_relation` VALUES (3, 21);
INSERT INTO `artist_song_relation` VALUES (9, 22);
INSERT INTO `artist_song_relation` VALUES (10, 23);
INSERT INTO `artist_song_relation` VALUES (11, 24);

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
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment_info
-- ----------------------------
INSERT INTO `comment_info` VALUES (8, 'aaaa nice', '2024-11-24 18:27:34', 'CrMhKMGnQH6R-Vq', 'moment', 6);
INSERT INTO `comment_info` VALUES (9, 'asdfasdf', '2024-11-24 18:28:01', 'CrMhKMGnQH6R-Vq', 'moment', 11);
INSERT INTO `comment_info` VALUES (10, 'asdfasdf', '2024-11-24 18:28:13', 'CrMhKMGnQH6R-Vq', 'moment', 13);
INSERT INTO `comment_info` VALUES (11, 'zxcasd   sdd', '2024-11-24 18:28:42', 'e9nRUN7ZRB6pDw6', 'moment', 18);
INSERT INTO `comment_info` VALUES (12, 'zxcasd   sdd', '2024-11-24 18:28:45', 'e9nRUN7ZRB6pDw6', 'moment', 18);
INSERT INTO `comment_info` VALUES (13, 'zxcasd   sdd', '2024-11-24 18:33:45', 'e9nRUN7ZRB6pDw6', 'moment', 19);
INSERT INTO `comment_info` VALUES (14, 'zxcasd   sdd', '2024-11-24 18:34:00', 'e9nRUN7ZRB6pDw6', 'moment', 19);
INSERT INTO `comment_info` VALUES (15, 'zxcasd   sdd', '2024-11-24 18:34:25', 'e9nRUN7ZRB6pDw6', 'moment', 19);
INSERT INTO `comment_info` VALUES (18, 'cupidatat', '2024-12-07 08:31:05', 'e9nRUN7ZRB6pDw6', 'moment', 15);
INSERT INTO `comment_info` VALUES (19, 'cupidatat', '2024-12-07 08:32:32', 'e9nRUN7ZRB6pDw6', 'moment', 15);
INSERT INTO `comment_info` VALUES (20, 'enim velit', '2024-12-07 09:07:15', 'e9nRUN7ZRB6pDw6', 'moment', 15);
INSERT INTO `comment_info` VALUES (21, 'commodo', '2024-12-07 09:07:19', 'e9nRUN7ZRB6pDw6', 'moment', 8);
INSERT INTO `comment_info` VALUES (22, 'nulla qui reprehenderit', '2024-12-07 09:07:21', 'e9nRUN7ZRB6pDw6', 'moment', 9);
INSERT INTO `comment_info` VALUES (24, 'tempor voluptate minim pariatur laborum', '2024-12-07 12:55:54', 'e9nRUN7ZRB6pDw6', 'moment', 12);

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
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 2);
INSERT INTO `follow_artist` VALUES ('e9nRUN7ZRB6pDw6', 2);
INSERT INTO `follow_artist` VALUES ('ZRQ6M-UcS2yedwY', 2);
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 3);
INSERT INTO `follow_artist` VALUES ('OMgiLMEDTxCSSB1', 3);
INSERT INTO `follow_artist` VALUES ('ZRQ6M-UcS2yedwY', 3);
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 4);
INSERT INTO `follow_artist` VALUES ('OMgiLMEDTxCSSB1', 4);
INSERT INTO `follow_artist` VALUES ('ZRQ6M-UcS2yedwY', 4);
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 5);
INSERT INTO `follow_artist` VALUES ('OMgiLMEDTxCSSB1', 5);
INSERT INTO `follow_artist` VALUES ('tgy6oWzjSYCjgOA', 5);
INSERT INTO `follow_artist` VALUES ('CrMhKMGnQH6R-Vq', 6);
INSERT INTO `follow_artist` VALUES ('e9nRUN7ZRB6pDw6', 6);
INSERT INTO `follow_artist` VALUES ('gPF1ZjgCTJSjqhU', 6);
INSERT INTO `follow_artist` VALUES ('tgy6oWzjSYCjgOA', 6);
INSERT INTO `follow_artist` VALUES ('gPF1ZjgCTJSjqhU', 7);
INSERT INTO `follow_artist` VALUES ('gPF1ZjgCTJSjqhU', 8);
INSERT INTO `follow_artist` VALUES ('tgy6oWzjSYCjgOA', 8);

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
INSERT INTO `follow_user` VALUES ('e9nRUN7ZRB6pDw6', 'CrMhKMGnQH6R-Vq');
INSERT INTO `follow_user` VALUES ('ZRQ6M-UcS2yedwY', 'CrMhKMGnQH6R-Vq');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'e9nRUN7ZRB6pDw6');
INSERT INTO `follow_user` VALUES ('e9nRUN7ZRB6pDw6', 'gPF1ZjgCTJSjqhU');
INSERT INTO `follow_user` VALUES ('e9nRUN7ZRB6pDw6', 'OMgiLMEDTxCSSB1');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'tgy6oWzjSYCjgOA');
INSERT INTO `follow_user` VALUES ('CrMhKMGnQH6R-Vq', 'ZRQ6M-UcS2yedwY');
INSERT INTO `follow_user` VALUES ('gPF1ZjgCTJSjqhU', 'ZRQ6M-UcS2yedwY');

-- ----------------------------
-- Table structure for like_comment
-- ----------------------------
DROP TABLE IF EXISTS `like_comment`;
CREATE TABLE `like_comment`  (
  `comment_id` int(11) NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`comment_id`, `user_id`) USING BTREE,
  INDEX `like_c_fk2`(`user_id` ASC) USING BTREE,
  CONSTRAINT `like_c_fk1` FOREIGN KEY (`comment_id`) REFERENCES `comment_info` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `like_c_fk2` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of like_comment
-- ----------------------------
INSERT INTO `like_comment` VALUES (8, 'CrMhKMGnQH6R-Vq');
INSERT INTO `like_comment` VALUES (9, 'CrMhKMGnQH6R-Vq');
INSERT INTO `like_comment` VALUES (18, 'CrMhKMGnQH6R-Vq');
INSERT INTO `like_comment` VALUES (20, 'e9nRUN7ZRB6pDw6');
INSERT INTO `like_comment` VALUES (19, 'gPF1ZjgCTJSjqhU');
INSERT INTO `like_comment` VALUES (22, 'OMgiLMEDTxCSSB1');
INSERT INTO `like_comment` VALUES (8, 'ZRQ6M-UcS2yedwY');
INSERT INTO `like_comment` VALUES (9, 'ZRQ6M-UcS2yedwY');
INSERT INTO `like_comment` VALUES (10, 'ZRQ6M-UcS2yedwY');

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
INSERT INTO `like_info` VALUES (11, 'CrMhKMGnQH6R-Vq');
INSERT INTO `like_info` VALUES (23, 'CrMhKMGnQH6R-Vq');
INSERT INTO `like_info` VALUES (20, 'e9nRUN7ZRB6pDw6');
INSERT INTO `like_info` VALUES (21, 'e9nRUN7ZRB6pDw6');
INSERT INTO `like_info` VALUES (15, 'gPF1ZjgCTJSjqhU');
INSERT INTO `like_info` VALUES (20, 'OMgiLMEDTxCSSB1');
INSERT INTO `like_info` VALUES (15, 'tgy6oWzjSYCjgOA');
INSERT INTO `like_info` VALUES (18, 'tgy6oWzjSYCjgOA');
INSERT INTO `like_info` VALUES (15, 'ZRQ6M-UcS2yedwY');
INSERT INTO `like_info` VALUES (19, 'ZRQ6M-UcS2yedwY');
INSERT INTO `like_info` VALUES (20, 'ZRQ6M-UcS2yedwY');

-- ----------------------------
-- Table structure for message_info
-- ----------------------------
DROP TABLE IF EXISTS `message_info`;
CREATE TABLE `message_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `sender_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `receiver_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `message_fk1`(`sender_id` ASC) USING BTREE,
  INDEX `message_fk2`(`receiver_id` ASC) USING BTREE,
  CONSTRAINT `message_fk1` FOREIGN KEY (`sender_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `message_fk2` FOREIGN KEY (`receiver_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of message_info
-- ----------------------------
INSERT INTO `message_info` VALUES (1, '2025-01-14 18:52:58', 'CrMhKMGnQH6R-Vq', 'e9nRUN7ZRB6pDw6', '你好！', 1);
INSERT INTO `message_info` VALUES (2, '2025-01-14 21:06:36', 'e9nRUN7ZRB6pDw6', 'CrMhKMGnQH6R-Vq', '你好呀！', 1);
INSERT INTO `message_info` VALUES (3, '2025-01-14 01:02:38', 'CrMhKMGnQH6R-Vq', 'gPF1ZjgCTJSjqhU', 'Hi', 0);
INSERT INTO `message_info` VALUES (4, '2025-01-08 01:03:58', 'CrMhKMGnQH6R-Vq', 'ZRQ6M-UcS2yedwY', '你喜欢什么类型的音乐？', 1);
INSERT INTO `message_info` VALUES (5, '2025-01-15 01:05:39', 'ZRQ6M-UcS2yedwY', 'CrMhKMGnQH6R-Vq', '我喜欢古典乐', 0);

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
INSERT INTO `moment_info` VALUES (18, 'aliquip amet', '2024-12-07 05:44:43', 'CrMhKMGnQH6R-Vq', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (19, 'aliquip amet', '2024-12-07 05:46:06', 'CrMhKMGnQH6R-Vq', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (20, 'aliquip amet', '2024-12-07 05:46:37', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (21, 'aliquip amet', '2024-12-07 06:01:28', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (22, 'aliquip amet', '2024-12-07 06:05:10', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');
INSERT INTO `moment_info` VALUES (23, 'aliquip amet', '2024-12-07 12:59:22', 'e9nRUN7ZRB6pDw6', 'https://loremflickr.com/400/400?lock=7745439452696942');

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
  `cover_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `playlist_fk1`(`user_id` ASC) USING BTREE,
  CONSTRAINT `playlist_fk1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of playlist_info
-- ----------------------------
INSERT INTO `playlist_info` VALUES (1, '那些被单曲循环无数次的歌', 'CrMhKMGnQH6R-Vq', '2023-12-01 19:51:45', '爱听就多多循环吧！每日更新，感谢支持！', '流行', 1998, 'https://p1.music.126.net/VHR1M8wfyhbNfwBfjqAiTw==/109951165473515928.jpg?param=200y200');
INSERT INTO `playlist_info` VALUES (2, '悲情古风：他朝若是同淋雪，此生也算共白头', 'e9nRUN7ZRB6pDw6', '2021-06-01 20:42:49', '这句话充满了悲情与古风之美，它描述了一种无奈中的深情，一种即便无法长久相守，也希望在某个瞬间能与你共同经历的情感。“他朝若是同淋雪”，这里的“他朝”指的是未来的某一天，而“同淋雪”则是一种情景的描绘，意味着在未来的某个时刻，两人能够一同经历风雪。这里的“雪”不仅仅是一种自然现象，更是象征着生活中的艰辛与困境。“此生也算共白头”，这里的“白头”通常指的是夫妻共度一生，直至白头偕老。但在这里，它有着更深层次的含义。即便两人无法长久相守，无法真正走到白头，但只要曾经有过那么一刻，两人能够一同面对风雪，一同经历生活的艰辛，那么此生也算是共同度过了某些重要的时刻，也算是一种心灵上的共鸣与陪伴。', '古风', 8900, 'https://p1.music.126.net/mCpqMfFxAc6YuqstAAXOEg==/109951169614143438.jpg?param=200y200');
INSERT INTO `playlist_info` VALUES (3, '夏日乡村 | 朗朗上口的夏日旋律', 'gPF1ZjgCTJSjqhU', '2021-02-01 20:46:22', '这份歌单将带你领略夏日乡村的美好 带给你轻松愉悦的听觉体验 伴随你度过一个充满阳光和温馨的夏日', '乡村', 8876, 'https://p1.music.126.net/JDQjfYUAWwdsQO5Mlk5KsA==/109951168648977448.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzMwNTkxNjcwNjU0LzY4YTkvMjAyMzgyNzExMjE0L3g0MzExNjk1NzgzNzM0MDAwLnBuZw==&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (4, '熬夜轻音乐 | 消除疲惫的旋律咖啡因', 'OMgiLMEDTxCSSB1', '2024-04-01 20:47:18', '睡不着的夜晚 用一首曼妙的钢琴曲 冲刷掉一整天的疲倦吧', '轻音乐', 2987, 'https://p1.music.126.net/Sff3IUB-pPUdJXr2euPviw==/109951168980789846.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NzQ5Mjc4MzIyL2ZhMzkvMjAyNDEwMjcxMzU4MjIveDgyNDE3MzI2ODcxMDIwMTMucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (5, '青春点唱机', 'tgy6oWzjSYCjgOA', '2020-03-01 20:48:30', '上课时 塞在袖子里的MP3 和同桌 一人一只的耳机 现在想想 那时的歌那么好听 只因我们 都曾听得入了神', '流行', 1999, 'https://p1.music.126.net/3lhdwKz7jHhoc3wl2WevEw==/109951170141161987.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NTA0NjYzMzIwLzdkYmMvMjAyNDEwMTUxNDU0MTcveDQ2MDE3MzE2NTM2NTc4NTMucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (6, '孤独患者 | 总有一首歌能包容你的心事', 'ZRQ6M-UcS2yedwY', '2024-09-01 20:49:45', '平淡告别，独自清醒，习惯孤独，成年人世界也许没有那么多可以把寂寞发泄的窗口，还好有音乐，它讲述了一切也包容了一切。', '流行', 5698, 'https://p1.music.126.net/LWDdeUXnSjD_YguIWUjEow==/109951168607138631.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzQ0NzgwNTIxMjg0LzhlMWUvMjAyNDYyNDExMjMwL3g3OTAxNzIxNzkwMTUwMDk4LnBuZw==&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (7, '去大自然旅行 | 来点乡村乐旋律', 'CrMhKMGnQH6R-Vq', '2023-05-11 20:51:03', '一起享受麦田味的旅行，有乡村、公路还有美好的阳光。', '乡村', 10089, 'https://p1.music.126.net/q_0TRnKnFuqnmAsWeSCsVQ==/109951168607336160.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzI4NzMyMTM4ODkxLzVkY2YvMjAyMzUxNDExMjcyNC94NjM3MTY4NjcxMzI0NDc5NC5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (8, '吉他弦下的田园风光', 'OMgiLMEDTxCSSB1', '2022-05-01 20:54:10', '像许多人一样你是否也向往着无忧无虑，隔绝世俗的世外桃源', '乡村', 899, 'https://p1.music.126.net/D8QDCiGkmBKuoQnL2DTwtA==/109951165005633820.jpg?param=200y200');
INSERT INTO `playlist_info` VALUES (9, '华语流行轻音乐', 'ZRQ6M-UcS2yedwY', '2019-10-01 20:55:32', '纯净的音符在空气中流淌 无需歌词 却能触动心灵的琴弦 华语流行纯音乐携着悸动的节奏 诉说着情感的故事 讲述着无尽的回忆', '轻音乐', 555, 'https://p1.music.126.net/lH4MAfT9Rm2bTswyT5hN0g==/109951169443307092.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NzQ3Njg0NjEwLzc2OTgvMjAyNDEwMjcxMTUxNTgveDMyMjE3MzI2Nzk1MTg1MzEucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (10, '新潮国风馆', 'tgy6oWzjSYCjgOA', '2014-01-14 20:57:14', '一起通过国风旋律 感受中华文化之美 领略那些充满中国味道的绝美创作', '古风', 6669, 'https://p1.music.126.net/_TpfX5uw7K1skadmUSWqJw==/109951168634398285.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzM2NjEwOTY0NDIxLzVjMGIvMjAyNDUxMjExMTk1Mi94Mzg2MTcxODE2MjM5MjY2Mi5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (11, '微醺K-Pop | 感受音乐与美酒的共鸣', 'CrMhKMGnQH6R-Vq', '2024-05-01 03:55:57', '沦陷于慵懒氛围 感受韩系音乐的温柔', '流行', 20008, 'https://p1.music.126.net/48Hmq5IvsOyd8wQVHmnS5A==/109951169361148873.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU1Njc3MTM2NTAxLzEyZDEvMjAyNDkxMDE0NTQ1NS94NzM4MTcyODU0MzI5NTU1MC5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (12, '爱在圣诞 | 温暖冬夜与音乐共舞', 'CrMhKMGnQH6R-Vq', '2023-12-01 04:01:25', '圣诞心愿伴着音符旋转跳跃～', '流行', 5567, 'https://p1.music.126.net/C8HUczIvg9CA0JDWpGSDvQ==/109951169107261020.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU3MTU2NTUzMDc0LzdlYzYvMjAyNDExMTkxNjQ2NTIveDYzOTE3MzQ1OTgwMTIyMzkucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (13, '国风电音大赏 | 电子水墨 鸾飘凤泊', 'e9nRUN7ZRB6pDw6', '2022-06-01 04:04:03', ' 当中国风遇上电子，犹如剑客出击，一剑碎困意！如果你喝腻了欧美浓咖般的提神音乐，不如试试这杯直追烈酒的国风浓茶，让你在秋冬迷蒙的清晨即刻清醒！', '古风', 2235, 'https://p1.music.126.net/bBn0Rg5YSBNSqt5lH9Ov6g==/109951168605350904.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzM2NjExMDIwMjk5L2I4M2YvMjAyNDUxMjExMjMyOS94NjI2MTcxODE2MjYwODk4Mi5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (14, '梦幻西游 | 梦回大唐盛世', 'e9nRUN7ZRB6pDw6', '2024-08-01 04:08:42', '梦幻西游以中国神话和文化为背景 拥有精美的画面和丰富的剧情 让玩家在游戏中体验到全面的社交和游戏乐趣 让玩家沉浸在梦幻的西游世界中 创造属于自己的传奇故事', '古风', 337, 'https://p1.music.126.net/3Bine2m64X1KH0LG-lP5lQ==/109951169102929135.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzMxNzg2NzMwNzk3LzZjN2QvMjAyMzEwMjgxNzE5MTgveDgzMDE3MDExNjMxNTg2NTYucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (15, '学习听国风 | 把君书卷灯前读', 'gPF1ZjgCTJSjqhU', '2023-09-01 04:09:53', '岁月悠长 白云苍狗 红尘中的少年捧起书卷 学向勤中得 萤窗万卷书', '古风', 1118, 'https://p1.music.126.net/j23rjySyyuVQusTbYthd5A==/109951170100622857.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NDgxMjQ4ODU3LzI3MTUvMjAyNDEwMTQxMTM5NC94MzYxMTczMTU1NTU0NDU3MC5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (16, 'Country Lover | 进入你的浪漫乌托邦', 'gPF1ZjgCTJSjqhU', '2024-10-01 04:10:53', ' 来往车辆喧嚣 你是否也步履匆匆 趁着阳光正好 跟随乡村音乐 感受老派浪漫吧', '乡村', 4532, 'https://p1.music.126.net/OAOBFu6i8d63EOWlRD_rmg==/109951168719796068.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzI5ODcwNzMyNzM3LzYyNDIvMjAyMzcxMTE1NTIxL3g5MjMxNjkxNzQwMzIxMDMyLnBuZw==&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (17, '溫溫乡音：步履田野寻觅逸趣安然', 'OMgiLMEDTxCSSB1', '2024-01-01 04:12:36', '迎着微风\r\n\r\n驱车路上\r\n\r\n阳光和煦\r\n\r\n岁月静好', '乡村', 367, 'https://p1.music.126.net/aqf2Rl9YG9tD6b-jXRqIOw==/109951168665695099.jpg?param=200y200');
INSERT INTO `playlist_info` VALUES (18, '散步轻音乐 | 和慵懒风光撞个满怀', 'CrMhKMGnQH6R-Vq', '2024-10-01 04:13:50', ' 就这样行走着 在温柔的风里 柔软的云朵里 明快的节奏里 还有橘子味的拥抱里', '轻音乐', 1235, 'https://p1.music.126.net/gyuT0kaTwRYmFLtr0ieeSw==/109951169009360894.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NzQ5MjQzMjMzL2IzMzkvMjAyNDEwMjcxMzU1MjQveDQzODE3MzI2ODY5MjQxMTIucG5n&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (19, '学习轻音乐 | 沉浸式学习 安静纯音旋律', 'tgy6oWzjSYCjgOA', '2025-01-03 04:14:51', '拿起耳机 让舒缓的旋律伴你沉浸在知识的海洋', '轻音乐', 3556, 'https://p1.music.126.net/bvplfpdgG8KGSnwAlKDXHQ==/109951170037636686.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NzUwMjAzNjA4LzIzMmQvMjAyNDEwMjcxNTk1Ny94MzE5MTczMjY5MTM5NzE4OC5wbmc=&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');
INSERT INTO `playlist_info` VALUES (20, '冬日轻音乐 | 柔软冬日里的温和轻音旋律', 'ZRQ6M-UcS2yedwY', '2020-05-01 04:14:58', ' 冬日悄悄到来 飞雪落满枝桠 让我们静静地温一壶茶 围炉夜话', '轻音乐', 5435, 'https://p1.music.126.net/2TMeYomgVVZPgx_pUGZtMQ==/109951169660249007.jpg?imageView=1&thumbnail=800y800&enlarge=1%7CimageView=1&watermark&type=1&image=b2JqL3c1bkRrTUtRd3JMRGpEekNtOE9tLzU2NzQ3NzczODQ4LzE5NzcvMjAyNDEwMjcxMjA0L3gxOTYxNzMyNjgwMDA0ODQ2LnBuZw==&dx=0&dy=0%7Cwatermark&type=1&image=b2JqL3dvbkRsc0tVd3JMQ2xHakNtOEt4LzI3NjEwNDk3MDYyL2VlOTMvOTIxYS82NjE4LzdhMDc5ZDg0NTYyMDAwZmVkZWJmMjVjYjE4NjhkOWEzLnBuZw==&dx=0&dy=0%7CimageView=1?param=200y200');

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
INSERT INTO `ranking_info` VALUES (1, '热歌', 1);
INSERT INTO `ranking_info` VALUES (2, '流行指数', 3);
INSERT INTO `ranking_info` VALUES (2, '热歌', 2);
INSERT INTO `ranking_info` VALUES (6, '新歌', 1);
INSERT INTO `ranking_info` VALUES (7, '新歌', 2);
INSERT INTO `ranking_info` VALUES (8, '新歌', 3);
INSERT INTO `ranking_info` VALUES (14, '热歌', 3);
INSERT INTO `ranking_info` VALUES (16, '流行指数', 1);
INSERT INTO `ranking_info` VALUES (18, '欧美', 1);
INSERT INTO `ranking_info` VALUES (18, '流行指数', 2);
INSERT INTO `ranking_info` VALUES (19, '欧美', 2);
INSERT INTO `ranking_info` VALUES (21, '欧美', 3);
INSERT INTO `ranking_info` VALUES (22, '韩国', 1);
INSERT INTO `ranking_info` VALUES (23, '韩国', 2);
INSERT INTO `ranking_info` VALUES (24, '韩国', 3);

-- ----------------------------
-- Table structure for setting_info
-- ----------------------------
DROP TABLE IF EXISTS `setting_info`;
CREATE TABLE `setting_info`  (
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `msg` int(11) NULL DEFAULT 0 COMMENT '0：所有人；1：关注的人',
  `see_rank` int(11) NULL DEFAULT 0 COMMENT '0：所有人可见；1：仅自己可见',
  `info_comment` int(11) NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_like` int(11) NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_msg` int(11) NULL DEFAULT 0 COMMENT '0：假；1：真',
  `info_sys` int(11) NULL DEFAULT 0 COMMENT '0：假；1：真',
  `service` int(11) NULL DEFAULT 0 COMMENT '0：假；1：真',
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
  `song_hit` int(11) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of song_info
-- ----------------------------
INSERT INTO `song_info` VALUES (1, '屋顶', 319, 1, '1', '2007-02-05', '1.mp3', '1.lrc', '2007-02-05 00:00:00', '2025-01-14 23:50:53', 1088);
INSERT INTO `song_info` VALUES (2, '我想念', 224, 2, '1', '2023-07-22', '2.mp3', '2.lrc', '2025-01-14 11:44:15', '2025-01-14 23:50:54', 2000);
INSERT INTO `song_info` VALUES (3, '不分手的恋爱', 205, 3, '1', '2011-09-09', '3.mp3', '3.lrc', '2025-01-14 13:28:05', '2025-01-14 23:50:54', 3567);
INSERT INTO `song_info` VALUES (4, '好安静', 265, 3, '1', '2011-09-09', '4.mp3', '4.lrc', '2025-01-14 14:15:12', '2025-01-14 23:50:54', 588);
INSERT INTO `song_info` VALUES (5, '我也不知道', 228, 3, '1', '2011-09-09', '5.mp3', '5.lrc', '2025-01-14 14:15:16', '2025-01-14 23:50:54', 999);
INSERT INTO `song_info` VALUES (6, '流浪是合理的需求', 275, 4, '1', '2024-12-08', '6.mp3', '6.lrc', '2025-01-14 14:41:46', '2025-01-14 23:50:55', 1111);
INSERT INTO `song_info` VALUES (7, '十万伏特', 279, 4, '1', '2024-12-08', '7.mp3', '7.lrc', '2025-01-08 15:24:45', '2025-01-14 23:50:55', 324);
INSERT INTO `song_info` VALUES (8, '想到我们', 305, 4, '1', '2024-12-08', '8.mp3', '8.lrc', '2024-10-10 15:25:46', '2025-01-14 23:50:55', 7689);
INSERT INTO `song_info` VALUES (9, '嚣', 252, 5, '1', '2017-04-03', '9.mp3', '9.lrc', '2024-04-17 15:41:14', '2025-01-14 23:50:55', 766);
INSERT INTO `song_info` VALUES (10, '怒花', 189, 5, '1', '2017-04-03', '10.mp3', '10.lrc', '2024-04-17 15:41:14', '2025-01-14 23:50:55', 897);
INSERT INTO `song_info` VALUES (11, '少年般绚丽', 190, 6, '1', '2022-12-31', '11.mp3', '11.lrc', '2024-04-17 15:41:14', '2025-01-14 23:50:59', 768);
INSERT INTO `song_info` VALUES (12, '讲话是闭嘴的时候', 291, 7, '1', '2022-10-28', '12.mp3', '12.lrc', '2023-04-01 15:54:28', '2025-01-14 23:50:58', 5639);
INSERT INTO `song_info` VALUES (13, '全城热恋', 254, 7, '1', '2022-10-28', '13.mp3', '13.lrc', '2023-06-01 15:54:35', '2025-01-14 23:50:56', 234);
INSERT INTO `song_info` VALUES (14, '大我年代', 297, 7, '1', '2022-10-28', '14.mp3', '14.lrc', '2023-10-01 15:54:38', '2025-01-14 23:51:00', 437);
INSERT INTO `song_info` VALUES (15, '时间', 273, 8, '1', '1996-10-08', '15.mp3', '15.lrc', '2020-01-01 15:58:25', '2025-01-14 23:51:01', 8973);
INSERT INTO `song_info` VALUES (16, '水星记', 325, 9, '1', '2016-11-25', '16.mp3', '16.lrc', '2023-04-01 16:01:05', '2025-01-14 23:51:01', 10008);
INSERT INTO `song_info` VALUES (17, '凄美地', 250, 9, '1', '2016-11-25', '17.mp3', '17.lrc', '2022-10-01 16:01:46', '2025-01-14 23:51:02', 15890);
INSERT INTO `song_info` VALUES (18, 'Maps', 190, 10, '1', '2014-08-29', '18.mp3', '18.lrc', '2020-05-01 16:05:18', '2025-01-14 23:51:02', 15783);
INSERT INTO `song_info` VALUES (19, 'One More Night', 219, 11, '1', '2012-06-20', '19.mp3', '19.lrc', '2020-05-01 16:05:18', '2025-01-14 23:51:03', 589);
INSERT INTO `song_info` VALUES (20, 'Minuet in G', 172, 12, '1', '1999-03-16', '20.mp3', '20.lrc', '2020-03-01 16:12:38', '2025-01-15 01:45:29', 167);
INSERT INTO `song_info` VALUES (21, 'Alone', 273, 13, '1', '2016-05-13', '21.mp3', '21.lrc', '2021-12-01 16:18:36', '2025-01-15 01:45:35', 12367);
INSERT INTO `song_info` VALUES (22, 'Walk', 191, 14, '1', '2024-07-15', '22.mp3', '22.lrc', '2024-08-01 01:45:36', '2025-01-15 01:52:18', 71367);
INSERT INTO `song_info` VALUES (23, 'Psycho', 210, 15, '1', '2019-12-23', '23.mp3', '23.lrc', '2021-03-01 01:47:21', '2025-01-15 01:52:23', 632675);
INSERT INTO `song_info` VALUES (24, '첫 눈', 207, 16, '1', '2013-12-09', '24.mp3', '24.lrc', '2018-06-22 01:51:45', '2025-01-15 02:40:53', 54367);

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
INSERT INTO `song_playlist_relation` VALUES (1, 1);
INSERT INTO `song_playlist_relation` VALUES (3, 1);
INSERT INTO `song_playlist_relation` VALUES (5, 1);
INSERT INTO `song_playlist_relation` VALUES (7, 1);
INSERT INTO `song_playlist_relation` VALUES (1, 2);
INSERT INTO `song_playlist_relation` VALUES (5, 2);
INSERT INTO `song_playlist_relation` VALUES (8, 2);
INSERT INTO `song_playlist_relation` VALUES (1, 3);
INSERT INTO `song_playlist_relation` VALUES (9, 3);
INSERT INTO `song_playlist_relation` VALUES (1, 4);
INSERT INTO `song_playlist_relation` VALUES (3, 4);
INSERT INTO `song_playlist_relation` VALUES (6, 4);
INSERT INTO `song_playlist_relation` VALUES (9, 4);
INSERT INTO `song_playlist_relation` VALUES (1, 5);
INSERT INTO `song_playlist_relation` VALUES (3, 5);
INSERT INTO `song_playlist_relation` VALUES (6, 5);
INSERT INTO `song_playlist_relation` VALUES (7, 5);
INSERT INTO `song_playlist_relation` VALUES (10, 5);
INSERT INTO `song_playlist_relation` VALUES (1, 6);
INSERT INTO `song_playlist_relation` VALUES (3, 6);
INSERT INTO `song_playlist_relation` VALUES (10, 6);
INSERT INTO `song_playlist_relation` VALUES (1, 7);
INSERT INTO `song_playlist_relation` VALUES (5, 7);
INSERT INTO `song_playlist_relation` VALUES (10, 7);
INSERT INTO `song_playlist_relation` VALUES (4, 8);
INSERT INTO `song_playlist_relation` VALUES (6, 8);
INSERT INTO `song_playlist_relation` VALUES (10, 9);
INSERT INTO `song_playlist_relation` VALUES (4, 10);
INSERT INTO `song_playlist_relation` VALUES (10, 10);
INSERT INTO `song_playlist_relation` VALUES (5, 11);
INSERT INTO `song_playlist_relation` VALUES (5, 12);
INSERT INTO `song_playlist_relation` VALUES (6, 13);
INSERT INTO `song_playlist_relation` VALUES (8, 15);
INSERT INTO `song_playlist_relation` VALUES (8, 16);
INSERT INTO `song_playlist_relation` VALUES (9, 16);
INSERT INTO `song_playlist_relation` VALUES (2, 17);
INSERT INTO `song_playlist_relation` VALUES (3, 17);
INSERT INTO `song_playlist_relation` VALUES (8, 17);
INSERT INTO `song_playlist_relation` VALUES (9, 17);
INSERT INTO `song_playlist_relation` VALUES (2, 18);
INSERT INTO `song_playlist_relation` VALUES (6, 18);
INSERT INTO `song_playlist_relation` VALUES (2, 19);
INSERT INTO `song_playlist_relation` VALUES (7, 19);
INSERT INTO `song_playlist_relation` VALUES (2, 20);
INSERT INTO `song_playlist_relation` VALUES (9, 20);
INSERT INTO `song_playlist_relation` VALUES (2, 21);
INSERT INTO `song_playlist_relation` VALUES (5, 21);
INSERT INTO `song_playlist_relation` VALUES (4, 22);
INSERT INTO `song_playlist_relation` VALUES (7, 22);
INSERT INTO `song_playlist_relation` VALUES (4, 23);
INSERT INTO `song_playlist_relation` VALUES (8, 23);
INSERT INTO `song_playlist_relation` VALUES (4, 24);
INSERT INTO `song_playlist_relation` VALUES (7, 24);

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
INSERT INTO `user_info` VALUES ('CrMhKMGnQH6R-Vq', 'sse_user', '$2a$10$xL8TTTC6EXuXyHytmRc6v.fjiF3StFDwRbFmOTNVqHzLsORpHsd2u', 'dj05@qq.com', '', '2024-11-13 07:48:03', '中国', '广东', '女', 'yeah', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2025-01-15 02:59:17');
INSERT INTO `user_info` VALUES ('e9nRUN7ZRB6pDw6', 'hhh', '$2a$10$vhsy/vbE/uA.3celgfTu7urxMjAcs1BorkTL8Lrs70frnMkPt8MnO', 'dj02@qq.com', '', '2024-11-13 07:47:43', '韩国', '首尔', '男', 'moa', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2025-01-15 02:59:49');
INSERT INTO `user_info` VALUES ('gPF1ZjgCTJSjqhU', '禄霞', 'Rphttk_gAEq92Xr', 'ppm_lfn86@vip.qq.com', '42794543813', '2024-11-25 14:10:13', '中国', '东北', '男', '战斗粉丝', 'https://loremflickr.com/400/400?lock=1075646362105748', '2025-01-14 23:51:54');
INSERT INTO `user_info` VALUES ('OMgiLMEDTxCSSB1', 'dj01', '$2a$10$YDzhOkP4/mtrS6yWgjEvROZYMia6RSx2jF95jEJCgTjvQCAdaD5OO', '1796654305@qq.com', '', '2024-12-07 13:59:12', '中国', '北京', '女', '啦啦啦', 'https://bpic.588ku.com/element_origin_min_pic/19/09/11/c38b3015813868a38d3067722e57d5ba.jpg', '2025-01-14 23:57:50');
INSERT INTO `user_info` VALUES ('tgy6oWzjSYCjgOA', 'dj03', '$2a$10$1JcBEF.dmGEe64G2orpFUeY4wO876EpxNtIrbAFEx/5KCGfbNezfy', 'dj03@qq.com', '', '2024-11-13 07:47:54', '中国', '浙江', '女', 'czenni', 'https://fastly.picsum.photos/id/793/200/200.jpg?hmac=3DeE830wjdSShKq_h_iFtV_jAxf43FO4xx-sivW0Q_Y', '2025-01-14 23:58:00');
INSERT INTO `user_info` VALUES ('ZRQ6M-UcS2yedwY', '磨文昊', '$2a$10$imLcf8YOIQ9/rt8aA.mXLO1eIqao2.kfcd4atB802b.6f2cx7BQNu', 'mawnuy_i4n52@foxmail.com', '99855657488', '2024-11-13 07:47:12', '韩国', '清潭洞', '女', '航空邮件倡导者，梦想家', 'https://loremflickr.com/400/400?lock=1157922648024234', '2025-01-14 23:56:29');

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
INSERT INTO `user_like_playlist` VALUES ('e9nRUN7ZRB6pDw6', 1);
INSERT INTO `user_like_playlist` VALUES ('tgy6oWzjSYCjgOA', 2);
INSERT INTO `user_like_playlist` VALUES ('CrMhKMGnQH6R-Vq', 3);
INSERT INTO `user_like_playlist` VALUES ('ZRQ6M-UcS2yedwY', 3);
INSERT INTO `user_like_playlist` VALUES ('tgy6oWzjSYCjgOA', 4);
INSERT INTO `user_like_playlist` VALUES ('e9nRUN7ZRB6pDw6', 5);
INSERT INTO `user_like_playlist` VALUES ('OMgiLMEDTxCSSB1', 6);
INSERT INTO `user_like_playlist` VALUES ('OMgiLMEDTxCSSB1', 7);
INSERT INTO `user_like_playlist` VALUES ('gPF1ZjgCTJSjqhU', 8);
INSERT INTO `user_like_playlist` VALUES ('CrMhKMGnQH6R-Vq', 9);
INSERT INTO `user_like_playlist` VALUES ('gPF1ZjgCTJSjqhU', 9);
INSERT INTO `user_like_playlist` VALUES ('CrMhKMGnQH6R-Vq', 10);
INSERT INTO `user_like_playlist` VALUES ('ZRQ6M-UcS2yedwY', 10);

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
INSERT INTO `user_like_song` VALUES ('e9nRUN7ZRB6pDw6', 2);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 3);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 4);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 5);
INSERT INTO `user_like_song` VALUES ('e9nRUN7ZRB6pDw6', 6);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 7);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 8);
INSERT INTO `user_like_song` VALUES ('OMgiLMEDTxCSSB1', 8);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 9);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 10);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 11);
INSERT INTO `user_like_song` VALUES ('e9nRUN7ZRB6pDw6', 17);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 22);
INSERT INTO `user_like_song` VALUES ('CrMhKMGnQH6R-Vq', 24);

SET FOREIGN_KEY_CHECKS = 1;
