-- MySQL dump 10.13  Distrib 8.0.30, for Linux (x86_64)
--
-- Host: localhost    Database: ecommerce
-- ------------------------------------------------------
-- Server version	8.0.30-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `level` tinyint(1) NOT NULL DEFAULT '1',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_login_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','20d135f0f28185b84a4cf7aa51f29500',3,1,'2022-08-13 12:52:22','2022-08-17 15:02:24','2022-08-17 15:02:24');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `type` tinyint(1) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,0,'??i???n tho???i','dien-thoai',0,1,'2022-08-13 13:39:10','2022-08-13 13:39:10'),(2,0,'??m thanh','am-thanh',0,1,'2022-08-13 13:39:27','2022-08-17 15:05:44'),(3,0,'M??y t??nh b???ng','may-tinh-bang',0,1,'2022-08-13 13:39:40','2022-08-13 13:39:40'),(4,2,'Tai nghe','tai-nghe',1,1,'2022-08-17 15:02:47','2022-08-17 15:06:16'),(5,2,'Loa','loa',1,1,'2022-08-17 15:04:22','2022-08-17 15:06:25'),(6,0,'Ph??? ki???n','phu-kien',0,1,'2022-08-17 15:25:26','2022-08-17 15:25:26'),(7,6,'??? c???ng','o-cung',1,1,'2022-08-17 15:26:04','2022-08-17 15:26:04'),(8,6,'C?????ng l???c','cuong-luc',1,1,'2022-08-17 15:26:33','2022-08-17 15:26:33'),(9,6,'B??n ph??m - chu???t','ban-phim-chuot',1,1,'2022-08-17 15:27:10','2022-08-17 15:27:10'),(10,6,'Balo','balo',1,1,'2022-08-17 15:27:52','2022-08-17 15:27:52'),(11,6,'Ph??? ki???n m??y','phu-kien-may',1,1,'2022-08-17 15:28:10','2022-08-17 15:28:10');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `version` varchar(255) NOT NULL,
  `class` varchar(255) NOT NULL,
  `group` varchar(255) NOT NULL,
  `namespace` varchar(255) NOT NULL,
  `time` int NOT NULL,
  `batch` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (10,'2022-02-20-141414','App\\Database\\Migrations\\Admin','default','App',1660369574,1),(11,'2022-03-15-015657','App\\Database\\Migrations\\Menu','default','App',1660369574,1),(12,'2022-08-01-103328','App\\Database\\Migrations\\Category','default','App',1660369574,1),(13,'2022-08-02-061829','App\\Database\\Migrations\\Product','default','App',1660369574,1),(14,'2022-08-02-061956','App\\Database\\Migrations\\ProductItems','default','App',1660369574,1),(15,'2022-08-02-064557','App\\Database\\Migrations\\ProductItemColors','default','App',1660369574,1),(16,'2022-08-02-065047','App\\Database\\Migrations\\ProductItemImages','default','App',1660369574,1),(17,'2022-08-02-065113','App\\Database\\Migrations\\ProductAttributeValue','default','App',1660369574,1),(18,'2022-08-02-065254','App\\Database\\Migrations\\ProductAttribute','default','App',1660369862,2);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NOT NULL,
  `category_id` int NOT NULL,
  `name` varchar(512) NOT NULL,
  `slug` varchar(512) NOT NULL,
  `additional_information` varchar(2048) NOT NULL,
  `support_information` varchar(2048) NOT NULL,
  `description` text NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_admin_id_foreign` (`admin_id`),
  KEY `product_category_id_foreign` (`category_id`),
  CONSTRAINT `product_admin_id_foreign` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`),
  CONSTRAINT `product_category_id_foreign` FOREIGN KEY (`category_id`) REFERENCES `menu` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1,1,1,'Samsung Galaxy A23','samsung-galaxy-a23','<p>Kh&ocirc;ng</p>\r\n','<ul>\r\n	<li>T???ng cho kh&aacute;ch l???n ?????u mua h&agrave;ng online t???i web\r\n	<p>M&atilde; gi???m <strong>100.000?? &aacute;p d???ng ????n h&agrave;ng t??? 400.000??</strong></p>\r\n\r\n	<p>?????i si&ecirc;u th??? Online v???i <strong>15.000</strong> s???n ph???m ti&ecirc;u d&ugrave;ng, th???t, c&aacute;, rau&hellip;</p>\r\n\r\n	<p><strong>FREEship</strong> ????n h&agrave;ng t??? 300.000?? ho???c th&agrave;nh vi&ecirc;n V&Agrave;NG</p>\r\n\r\n	<p>Giao nhanh trong <strong>2 gi???</strong></p>\r\n	&Aacute;p d???ng t???i Tp.HCM v&agrave; 22 t???nh th&agrave;nh</li>\r\n	<li>\r\n	<p>T???ng su???t mua xe ?????p gi???m ?????n 30%(kh&ocirc;ng k&egrave;m khuy???n m&atilde;i kh&aacute;c)</p>\r\n	</li>\r\n</ul>\r\n','<h3>Samsung Galaxy A23 4GB s??? h???u b??? th&ocirc;ng s??? k??? thu???t kh&aacute; ???n t?????ng trong ph&acirc;n kh&uacute;c, m&aacute;y c&oacute; m???t hi???u n??ng ???n ?????nh, c???m 4 camera th&ocirc;ng minh c&ugrave;ng m???t di???n m???o tr??? trung ph&ugrave; h???p cho m???i ?????i t?????ng ng?????i d&ugrave;ng.</h3>\r\n\r\n<h3>X??? l&yacute; m?????t m&agrave; nh??? chipset ?????n t??? Qualcomm</h3>\r\n\r\n<p>????? m&aacute;y v???n h&agrave;nh m???t c&aacute;ch ???n ?????nh h??n&nbsp;Samsung trang b??? cho Galaxy A23 con chip qu???c d&acirc;n d&agrave;nh cho th??? tr?????ng di ?????ng t???m trung hi???n nay (04/2022) mang t&ecirc;n Snapdragon 680 8 nh&acirc;n v???i hi???u su???t t???i ??a ?????t ???????c l&agrave; 2.4 GHz.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113212.jpg\" onclick=\"return false;\"><img alt=\"Hi???u n??ng m???nh m???  - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113212.jpg\" /></a></p>\r\n\r\n<p>??&aacute;nh gi&aacute; s???c m???nh c???a thi???t b??? qua hai ???ng d???ng th?????ng ???????c m???i ng?????i d&ugrave;ng ????? so s&aacute;nh hi???u n??ng v???i k???t qu??? ?????t ???????c nh?? sau: 283 (????n nh&acirc;n), 1515 (??a nh&acirc;n) tr&ecirc;n Benchmark v&agrave; 6830 cho ???ng d???ng PCMark.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113214.jpg\" onclick=\"return false;\"><img alt=\"??i???m ????nh gi?? hi???u n??ng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113214.jpg\" /></a></p>\r\n\r\n<p>Kh???i ?????ng v???i m???t t???a game chi???n thu???t hi???n ??ang r???t th???nh h&agrave;nh l&agrave; Li&ecirc;n Qu&acirc;n Mobile ??? m???c c???u h&igrave;nh max setting, Galaxy A23 cho ra qu&aacute; tr&igrave;nh ch??i t????ng ?????i l&agrave; m?????t m&agrave;, h&igrave;nh ???nh ?????p m???t v&agrave; k??? n??ng ???????c t&aacute;i hi???n s???ng ?????ng. T&igrave;nh tr???ng gi???t, lag v???n c&ograve;n xu???t hi???n nh??ng kh&ocirc;ng qu&aacute; ??&aacute;ng k???, t???c ????? khung h&igrave;nh dao ?????ng loanh quanh ??? m???c 55 - 60 FPS.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113216.jpg\" onclick=\"return false;\"><img alt=\"C???u h??nh Li??n Qu??n Mobile - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113216.jpg\" /></a></p>\r\n\r\n<p>Tr&ograve; ch??i th??? 2 m&agrave; m&igrave;nh th??? qua l&agrave; PUBG Mobile v???i c???u h&igrave;nh ???????c m&igrave;nh thi???t l???p l&agrave; m???c ????? h???a m?????t v&agrave; t???c ????? khung h&igrave;nh c???c cao ????? ?????m b???o m&aacute;y ho???t ?????ng m???t c&aacute;ch ???n ?????nh nh???t c&oacute; th???, nh??ng b&ugrave; l???i h&igrave;nh ???nh th??? hi???n kh&ocirc;ng ???????c xu???t s???c.</p>\r\n\r\n<p>Tr???i nghi???m ???n ?????nh, nh&acirc;n v???t di chuy???n t????ng ?????i l&agrave; m?????t m&agrave;, qu&aacute; tr&igrave;nh ??i nh???t ????? di???n ra kh&aacute; m?????t, t???c ????? khung h&igrave;nh dao ?????ng loanh quanh ??? m???c 25 - 29 FPS. Hi???n t?????ng kh???ng khung h&igrave;nh khi m&igrave;nh tham chi???n ??? nh???ng v??? tr&iacute; ??&ocirc;ng k??? ?????ch v???n c&ograve;n xu???t hi???n nh??ng kh&ocirc;ng ?????n m???c kh&oacute; ch???u.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113218.jpg\" onclick=\"return false;\"><img alt=\"C???u h??nh PUBG Mobile - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113218.jpg\" /></a></p>\r\n\r\n<p>M&aacute;y trang b??? 4 GB RAM c&ugrave;ng 128 GB b??? nh??? trong mang ?????n kh??? n??ng ??a nhi???m m???t c&aacute;ch m?????t m&agrave; c&ugrave;ng kh&ocirc;ng gian l??u tr??? ??&aacute;p ???ng v???a ????? cho ng?????i d&ugrave;ng c?? b???n ????? t???i xu???ng m???t l?????ng ???ng d???ng v&agrave; h&igrave;nh ???nh kh&aacute; l???n.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113220.jpg\" onclick=\"return false;\"><img alt=\"??a nhi???m m?????t m?? - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113220.jpg\" /></a></p>\r\n\r\n<h3>Ch???p ???nh s???c n&eacute;t v???i c???m camera th&ocirc;ng minh</h3>\r\n\r\n<p>M&aacute;y s??? h???u 4 camera v???i camera ch&iacute;nh c&oacute; ????? ph&acirc;n gi???i l&ecirc;n ?????n 50 MP, camera g&oacute;c r???ng 5 MP,&nbsp;c???m bi???n x&oacute;a ph&ocirc;ng&nbsp;v&agrave;&nbsp;macro c&oacute; c&ugrave;ng ????? ph&acirc;n gi???i 2 MP, k&egrave;m v???i ??&oacute; l&agrave; nhi???u t&iacute;nh n??ng ch???p ???nh m???i l??? gi&uacute;p m&igrave;nh th???a s???c kh&aacute;m ph&aacute;.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113223.jpg\" onclick=\"return false;\"><img alt=\"Trang b??? 4 camera - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113223.jpg\" /></a></p>\r\n\r\n<p>Kh&aacute; ???n t?????ng v??? k???t qu??? thu l???i tr&ecirc;n b???c h&igrave;nh m&agrave; m&igrave;nh c&oacute; ch???p t??? ??i???n tho???i khi ??ang di chuy???n ngo&agrave;i ???????ng, m&agrave;u s???c ???nh c&oacute; ????? t????ng ph???n cao, c&aacute;c chi ti???t nh??? ?????u ???????c m&aacute;y thu l???i r&otilde; n&eacute;t, ???nh kh&ocirc;ng qu&aacute; &ldquo;b???&rdquo; khi zoom hay h???u k??? - ch???nh s???a.&nbsp;</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113221.jpg\" onclick=\"return false;\"><img alt=\"???nh ch???p ??? m??i tr?????ng ????? s??ng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113221.jpg\" /></a></p>\r\n\r\n<p>V??? ph???n ch???p ??&ecirc;m th&igrave; Galaxy A23 ch??a mang l???i k???t qu??? t???t, t???ng th??? b???c ???nh ch??? d???ng ??? m???c ch???p nh???n ???????c, hi???n t?????ng &ldquo;nh&ograve;e s&aacute;ng&rdquo; ??? nh???ng v??? tr&iacute; ngu???n s&aacute;ng cao nh?? b&oacute;ng ??&egrave;n v???n xu???t hi???n.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113225.jpg\" onclick=\"return false;\"><img alt=\"???nh ch???p v??o bu???i t???i - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113225.jpg\" /></a></p>\r\n\r\n<p>??? ch??? ????? ch???p ???nh selfie b???ng camera 8 MP cho ra b???c h&igrave;nh s???c n&eacute;t, n?????c da kh&ocirc;ng qu&aacute; b???t, s??? d???ng t&iacute;nh n??ng l&agrave;m ?????p ??i k&egrave;m ????? che ??i nh???ng khuy???t ??i???m li ti nh?? m???n, n???t ru???i nh??? gi&uacute;p m&igrave;nh t??? tin h??n tr&ecirc;n c&aacute;c b???c ???nh t??? ch???p.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113226.jpg\" onclick=\"return false;\"><img alt=\"???nh ch???p selfie - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113226.jpg\" /></a></p>\r\n\r\n<h3>Thi???t k??? ?????c tr??ng t??? d&ograve;ng Galaxy A</h3>\r\n\r\n<p>Galaxy A23 c&oacute; khung v&agrave; m???t l??ng ???????c l&agrave;m t??? nh???a mang l???i c???m gi&aacute;c c???m n???m nh??? nh&agrave;ng, c&ugrave;ng v???i ??&oacute; l&agrave; c???nh vi???n bo cong gi&uacute;p m&igrave;nh s??? d???ng l&acirc;u d&agrave;i m&agrave; kh&ocirc;ng c???m th???y b??? c???n tay nh?? tr&ecirc;n m???t s??? d&ograve;ng s???n ph???m c&oacute; thi???t k??? vu&ocirc;ng v???c.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113228.jpg\" onclick=\"return false;\"><img alt=\"Thi???t k??? ho??n to??n t??? nh???a - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113228.jpg\" /></a></p>\r\n\r\n<p>S??? h???u m???t m???t l??ng s&aacute;ng b&oacute;ng c&ugrave;ng m&agrave;u s???c tr??? trung gi&uacute;p m&igrave;nh tr??? n&ecirc;n n???i b???t h??n khi c???m chi???c m&aacute;y tr&ecirc;n tay, tuy nhi&ecirc;n theo m&igrave;nh ??&acirc;y c??ng l&agrave; m???t ??i???m h???n ch??? b???i n&oacute; v???n xu???t hi???n t&igrave;nh tr???ng b&aacute;m d???u v&acirc;n tay.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113230.jpg\" onclick=\"return false;\"><img alt=\"M???t l??ng s??ng b??ng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113230.jpg\" /></a></p>\r\n\r\n<p>Tr&ecirc;n c???nh vi???n ???????c b??? tr&iacute; ph&iacute;m ngu???n t&iacute;ch h???p c???m bi???n v&acirc;n tay v???i t???c ????? ph???n h???i kh&aacute; nhanh, v??? tr&iacute; ?????t c???a ph&iacute;m c??ng kh&ocirc;ng qu&aacute; cao gi&uacute;p m&igrave;nh d??? d&agrave;ng k&iacute;ch ho???t thi???t b??? m???t c&aacute;ch nhanh ch&oacute;ng ch??? v???i m???t tay.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113232.jpg\" onclick=\"return false;\"><img alt=\"C???m bi???n v??n tay c???nh vi???n - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113232.jpg\" /></a></p>\r\n\r\n<p>M???t tr?????c l&agrave; m&agrave;n h&igrave;nh gi???t n?????c, trang b??? t???m n???n PLS TFT LCD v???i k&iacute;ch th?????c 6.6 inch c&oacute; ????? ph&acirc;n gi???i Full HD+ (1080 x 2408 Pixels) cho ra m&agrave;u s???c c&oacute; ????? t????ng ph???n cao, h&igrave;nh ???nh t&aacute;i hi???n chi ti???t c&ugrave;ng m???t kh&ocirc;ng gian hi???n th??? r???ng l???n h??? tr??? cho c&aacute;c t&aacute;c v??? h???c t???p, l&agrave;m vi???c t???t h??n.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113233.jpg\" onclick=\"return false;\"><img alt=\"M??n h??nh k??ch th?????c l???n - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113233.jpg\" /></a></p>\r\n\r\n<p>M&agrave;n h&igrave;nh 90 Hz c??ng l&agrave; m???t ??i???m s&aacute;ng tr&ecirc;n Galaxy A23 v&igrave; n&oacute; gi&uacute;p m&igrave;nh thao t&aacute;c c&ocirc;ng vi???c m???t c&aacute;ch m?????t m&agrave; v&agrave; ??&atilde; m???t h??n, c&ugrave;ng v???i ??&oacute; l&agrave; t&iacute;nh n??ng t??? ?????ng ??i???u ch???nh t???n s??? qu&eacute;t xu???ng m???c m???c ?????nh (60 Hz) nh???m ti???t ki???m ??i???n n??ng cho nh???ng t&aacute;c v??? kh&ocirc;ng c???n thi???t nh?? ?????c v??n b???n, ??&agrave;m tho???i r???t h???u &iacute;ch.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113235.jpg\" onclick=\"return false;\"><img alt=\"M??n h??nh  t???n s??? 90 Hz - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113235.jpg\" /></a></p>\r\n\r\n<h3>Pin tr&acirc;u d&ugrave;ng l&acirc;u c??? ng&agrave;y</h3>\r\n\r\n<p>B&ecirc;n trong thi???t b??? l&agrave; vi&ecirc;n pin c&oacute; dung l?????ng 5000 mAh th???a s???c ??&aacute;p ???ng nhu c???u s??? d???ng li&ecirc;n t???c trong nhi???u gi???, m&igrave;nh c&oacute; d&ugrave;ng m&aacute;y ????? ph???c v??? cho vi???c l?????t web, xem phim, ch??i game cho th???y Galaxy A23 ??&aacute;p ???ng th???i l?????ng l&ecirc;n ?????n h??n 8 ti???ng*.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113237.jpg\" onclick=\"return false;\"><img alt=\"Th???i l?????ng s??? d???ng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113237.jpg\" /></a></p>\r\n\r\n<p>Samsung trang b??? cho m&aacute;y c&ocirc;ng ngh??? s???c pin nhanh 25 W nh???m t???i ??u th???i gian s???c gi&uacute;p r&uacute;t ng???n th???i gian l???p ?????y vi&ecirc;n pin xu???ng c&ograve;n 1 gi??? 35 ph&uacute;t*, ??&acirc;y l&agrave; m???t kho???ng th???i gian kh&aacute; h???p l&yacute;, kh&ocirc;ng qu&aacute; l&acirc;u.</p>\r\n',1,'2022-08-13 14:35:42','2022-08-13 17:10:59');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_attribute_values`
--

DROP TABLE IF EXISTS `product_attribute_values`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_attribute_values` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(512) NOT NULL,
  `key` varchar(255) NOT NULL,
  `value` varchar(2048) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_attribute_values`
--

LOCK TABLES `product_attribute_values` WRITE;
/*!40000 ALTER TABLE `product_attribute_values` DISABLE KEYS */;
INSERT INTO `product_attribute_values` VALUES (1,'C??ng ngh??? m??n h??nh OLED LPTO','C??ng ngh??? m??n h??nh','OLED LPTO',1,'2022-08-13 13:59:52','2022-08-13 13:59:52'),(2,' ????? ph??n gi???i  FHD+ (2640 x 1080 Pixels)',' ????? ph??n gi???i','FHD+ (2640 x 1080 Pixels)',1,'2022-08-13 14:01:43','2022-08-13 14:01:43'),(3,' ????? s??ng t???i ??a  1200 nits',' ????? s??ng t???i ??a','1200 nits',1,'2022-08-13 14:02:20','2022-08-13 14:02:20'),(4,'M??n h??nh PLS TFT LCD 6.6\" Full HD+','M??n h??nh','PLS TFT LCD 6.6\" Full HD+',1,'2022-08-13 14:06:09','2022-08-13 14:06:09'),(5,'H??? ??i???u h??nh Android 12','H??? ??i???u h??nh','Android 12',0,'2022-08-13 14:06:36','2022-08-13 14:06:36'),(6,'Camera sau Ch??nh 50 MP & Ph??? 5 MP, 2 MP, 2 MP','Camera sau','Ch??nh 50 MP & Ph??? 5 MP, 2 MP, 2 MP',0,'2022-08-13 14:06:56','2022-08-13 14:06:56'),(7,'Camera tr?????c 8 MP','Camera tr?????c','8 MP',1,'2022-08-13 14:31:29','2022-08-13 14:31:29'),(8,'  Chip Snapdragon 680 8 nh??n','Chip','Snapdragon 680 8 nh??n',1,'2022-08-13 14:31:49','2022-08-13 14:31:49'),(9,'RAM ','RAM','4 GB',1,'2022-08-13 14:32:12','2022-08-13 14:32:12'),(10,'B??? nh??? trong 128 GB','B??? nh??? trong','128 GB',0,'2022-08-13 14:32:37','2022-08-13 14:32:37'),(11,'SIM 2 Nano SIM H??? tr??? 4G','SIM','2 Nano SIM H??? tr??? 4G',1,'2022-08-13 14:33:02','2022-08-13 14:33:02'),(12,'Pin, S???c 5000 mAh 25 W','Pin, S???c','5000 mAh 25 W',1,'2022-08-13 14:33:24','2022-08-13 14:33:24');
/*!40000 ALTER TABLE `product_attribute_values` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_attributes`
--

DROP TABLE IF EXISTS `product_attributes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_attributes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `product_item_id` int DEFAULT NULL,
  `product_attribute_value_id` int NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_attributes_product_id_foreign` (`product_id`),
  KEY `product_attributes_product_item_id_foreign` (`product_item_id`),
  KEY `product_attributes_product_attribute_value_id_foreign` (`product_attribute_value_id`),
  CONSTRAINT `product_attributes_product_attribute_value_id_foreign` FOREIGN KEY (`product_attribute_value_id`) REFERENCES `product_attribute_values` (`id`),
  CONSTRAINT `product_attributes_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `product_attributes_product_item_id_foreign` FOREIGN KEY (`product_item_id`) REFERENCES `product_items` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_attributes`
--

LOCK TABLES `product_attributes` WRITE;
/*!40000 ALTER TABLE `product_attributes` DISABLE KEYS */;
INSERT INTO `product_attributes` VALUES (1,1,NULL,4,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(2,1,NULL,5,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(3,1,NULL,6,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(4,1,NULL,7,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(5,1,NULL,8,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(6,1,NULL,11,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(7,1,NULL,12,1,'2022-08-13 14:35:42','2022-08-13 14:35:42'),(8,1,1,9,1,'2022-08-13 19:07:59','2022-08-13 19:07:59'),(9,1,1,10,1,'2022-08-13 19:07:59','2022-08-13 19:07:59');
/*!40000 ALTER TABLE `product_attributes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_category`
--

DROP TABLE IF EXISTS `product_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_category_menu_id_foreign` (`menu_id`),
  CONSTRAINT `product_category_menu_id_foreign` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_category`
--

LOCK TABLES `product_category` WRITE;
/*!40000 ALTER TABLE `product_category` DISABLE KEYS */;
INSERT INTO `product_category` VALUES (1,1,'Samsung','samsung',1,'2022-08-13 13:43:42','2022-08-13 13:43:42'),(2,1,'Xiaomi','xiaomi',1,'2022-08-13 13:43:55','2022-08-13 13:43:55'),(3,1,'iPhone','iphone',0,'2022-08-13 13:44:12','2022-08-13 13:44:12'),(4,2,'iPod','ipod',0,'2022-08-13 13:44:24','2022-08-13 13:44:24');
/*!40000 ALTER TABLE `product_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_item_colors`
--

DROP TABLE IF EXISTS `product_item_colors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_item_colors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_item_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `hexcode` varchar(255) NOT NULL,
  `price` varchar(255) NOT NULL,
  `discount` varchar(255) NOT NULL,
  `quantity` int DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_item_colors_product_item_id_foreign` (`product_item_id`),
  CONSTRAINT `product_item_colors_product_item_id_foreign` FOREIGN KEY (`product_item_id`) REFERENCES `product_items` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_item_colors`
--

LOCK TABLES `product_item_colors` WRITE;
/*!40000 ALTER TABLE `product_item_colors` DISABLE KEYS */;
INSERT INTO `product_item_colors` VALUES (1,1,'??en','#000000','5690000','0',5690000,1,'2022-08-13 19:06:14','2022-08-13 19:06:14'),(3,1,'Cam','#ff8040','5690000','9',569,1,'2022-08-13 19:14:02','2022-08-13 19:14:02');
/*!40000 ALTER TABLE `product_item_colors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_item_images`
--

DROP TABLE IF EXISTS `product_item_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_item_images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_item_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_item_images_product_item_id_foreign` (`product_item_id`),
  CONSTRAINT `product_item_images_product_item_id_foreign` FOREIGN KEY (`product_item_id`) REFERENCES `product_items` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_item_images`
--

LOCK TABLES `product_item_images` WRITE;
/*!40000 ALTER TABLE `product_item_images` DISABLE KEYS */;
INSERT INTO `product_item_images` VALUES (1,1,'1660388288_37f7a3808b979d727997.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08'),(2,1,'1660388288_fe8d8a411617e085b4fd.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08'),(3,1,'1660388288_582c80935d7d1172d3a9.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08'),(4,1,'1660388288_d3763772e06032555e77.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08'),(5,1,'1660388288_f97cec3e98daaee16158.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08'),(6,1,'1660388288_373ac74ba76cd45cc4a4.jpg',1,'2022-08-13 17:58:08','2022-08-13 17:58:08');
/*!40000 ALTER TABLE `product_item_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_items`
--

DROP TABLE IF EXISTS `product_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `admin_id` int NOT NULL,
  `name` varchar(2048) NOT NULL,
  `slug` varchar(2048) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `product_items_product_id_foreign` (`product_id`),
  KEY `product_items_admin_id_foreign` (`admin_id`),
  CONSTRAINT `product_items_admin_id_foreign` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`),
  CONSTRAINT `product_items_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_items`
--

LOCK TABLES `product_items` WRITE;
/*!40000 ALTER TABLE `product_items` DISABLE KEYS */;
INSERT INTO `product_items` VALUES (1,1,1,'Samsung Galaxy A23 4GB','samsung-galaxy-a23-4gb',1,'2022-08-13 17:58:08','2022-08-13 17:58:08');
/*!40000 ALTER TABLE `product_items` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-18 13:14:07
