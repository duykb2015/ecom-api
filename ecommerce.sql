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
INSERT INTO `menu` VALUES (1,0,'Điện thoại','dien-thoai',0,1,'2022-08-13 13:39:10','2022-08-13 13:39:10'),(2,0,'Âm thanh','am-thanh',0,1,'2022-08-13 13:39:27','2022-08-17 15:05:44'),(3,0,'Máy tính bảng','may-tinh-bang',0,1,'2022-08-13 13:39:40','2022-08-13 13:39:40'),(4,2,'Tai nghe','tai-nghe',1,1,'2022-08-17 15:02:47','2022-08-17 15:06:16'),(5,2,'Loa','loa',1,1,'2022-08-17 15:04:22','2022-08-17 15:06:25'),(6,0,'Phụ kiện','phu-kien',0,1,'2022-08-17 15:25:26','2022-08-17 15:25:26'),(7,6,'Ổ cứng','o-cung',1,1,'2022-08-17 15:26:04','2022-08-17 15:26:04'),(8,6,'Cường lực','cuong-luc',1,1,'2022-08-17 15:26:33','2022-08-17 15:26:33'),(9,6,'Bàn phím - chuột','ban-phim-chuot',1,1,'2022-08-17 15:27:10','2022-08-17 15:27:10'),(10,6,'Balo','balo',1,1,'2022-08-17 15:27:52','2022-08-17 15:27:52'),(11,6,'Phụ kiện máy','phu-kien-may',1,1,'2022-08-17 15:28:10','2022-08-17 15:28:10');
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
INSERT INTO `product` VALUES (1,1,1,'Samsung Galaxy A23','samsung-galaxy-a23','<p>Kh&ocirc;ng</p>\r\n','<ul>\r\n	<li>Tặng cho kh&aacute;ch lần đầu mua h&agrave;ng online tại web\r\n	<p>M&atilde; giảm <strong>100.000đ &aacute;p dụng đơn h&agrave;ng từ 400.000đ</strong></p>\r\n\r\n	<p>Đại si&ecirc;u thị Online với <strong>15.000</strong> sản phẩm ti&ecirc;u d&ugrave;ng, thịt, c&aacute;, rau&hellip;</p>\r\n\r\n	<p><strong>FREEship</strong> đơn h&agrave;ng từ 300.000đ hoặc th&agrave;nh vi&ecirc;n V&Agrave;NG</p>\r\n\r\n	<p>Giao nhanh trong <strong>2 giờ</strong></p>\r\n	&Aacute;p dụng tại Tp.HCM v&agrave; 22 tỉnh th&agrave;nh</li>\r\n	<li>\r\n	<p>Tặng suất mua xe đạp giảm đến 30%(kh&ocirc;ng k&egrave;m khuyến m&atilde;i kh&aacute;c)</p>\r\n	</li>\r\n</ul>\r\n','<h3>Samsung Galaxy A23 4GB sở hữu bộ th&ocirc;ng số kỹ thuật kh&aacute; ấn tượng trong ph&acirc;n kh&uacute;c, m&aacute;y c&oacute; một hiệu năng ổn định, cụm 4 camera th&ocirc;ng minh c&ugrave;ng một diện mạo trẻ trung ph&ugrave; hợp cho mọi đối tượng người d&ugrave;ng.</h3>\r\n\r\n<h3>Xử l&yacute; mượt m&agrave; nhờ chipset đến từ Qualcomm</h3>\r\n\r\n<p>Để m&aacute;y vận h&agrave;nh một c&aacute;ch ổn định hơn&nbsp;Samsung trang bị cho Galaxy A23 con chip quốc d&acirc;n d&agrave;nh cho thị trường di động tầm trung hiện nay (04/2022) mang t&ecirc;n Snapdragon 680 8 nh&acirc;n với hiệu suất tối đa đạt được l&agrave; 2.4 GHz.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113212.jpg\" onclick=\"return false;\"><img alt=\"Hiệu năng mạnh mẽ  - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113212.jpg\" /></a></p>\r\n\r\n<p>Đ&aacute;nh gi&aacute; sức mạnh của thiết bị qua hai ứng dụng thường được mọi người d&ugrave;ng để so s&aacute;nh hiệu năng với kết quả đạt được như sau: 283 (đơn nh&acirc;n), 1515 (đa nh&acirc;n) tr&ecirc;n Benchmark v&agrave; 6830 cho ứng dụng PCMark.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113214.jpg\" onclick=\"return false;\"><img alt=\"Điểm đánh giá hiệu năng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113214.jpg\" /></a></p>\r\n\r\n<p>Khởi động với một tựa game chiến thuật hiện đang rất thịnh h&agrave;nh l&agrave; Li&ecirc;n Qu&acirc;n Mobile ở mức cấu h&igrave;nh max setting, Galaxy A23 cho ra qu&aacute; tr&igrave;nh chơi tương đối l&agrave; mượt m&agrave;, h&igrave;nh ảnh đẹp mắt v&agrave; kỹ năng được t&aacute;i hiện sống động. T&igrave;nh trạng giật, lag vẫn c&ograve;n xuất hiện nhưng kh&ocirc;ng qu&aacute; đ&aacute;ng kể, tốc độ khung h&igrave;nh dao động loanh quanh ở mức 55 - 60 FPS.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113216.jpg\" onclick=\"return false;\"><img alt=\"Cấu hình Liên Quân Mobile - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113216.jpg\" /></a></p>\r\n\r\n<p>Tr&ograve; chơi thứ 2 m&agrave; m&igrave;nh thử qua l&agrave; PUBG Mobile với cấu h&igrave;nh được m&igrave;nh thiết lập l&agrave; mức đồ họa mượt v&agrave; tốc độ khung h&igrave;nh cực cao để đảm bảo m&aacute;y hoạt động một c&aacute;ch ổn định nhất c&oacute; thể, nhưng b&ugrave; lại h&igrave;nh ảnh thể hiện kh&ocirc;ng được xuất sắc.</p>\r\n\r\n<p>Trải nghiệm ổn định, nh&acirc;n vật di chuyển tương đối l&agrave; mượt m&agrave;, qu&aacute; tr&igrave;nh đi nhặt đồ diễn ra kh&aacute; mượt, tốc độ khung h&igrave;nh dao động loanh quanh ở mức 25 - 29 FPS. Hiện tượng khựng khung h&igrave;nh khi m&igrave;nh tham chiến ở những vị tr&iacute; đ&ocirc;ng kẻ địch vẫn c&ograve;n xuất hiện nhưng kh&ocirc;ng đến mức kh&oacute; chịu.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113218.jpg\" onclick=\"return false;\"><img alt=\"Cấu hình PUBG Mobile - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113218.jpg\" /></a></p>\r\n\r\n<p>M&aacute;y trang bị 4 GB RAM c&ugrave;ng 128 GB bộ nhớ trong mang đến khả năng đa nhiệm một c&aacute;ch mượt m&agrave; c&ugrave;ng kh&ocirc;ng gian lưu trữ đ&aacute;p ứng vừa đủ cho người d&ugrave;ng cơ bản để tải xuống một lượng ứng dụng v&agrave; h&igrave;nh ảnh kh&aacute; lớn.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113220.jpg\" onclick=\"return false;\"><img alt=\"Đa nhiệm mượt mà - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113220.jpg\" /></a></p>\r\n\r\n<h3>Chụp ảnh sắc n&eacute;t với cụm camera th&ocirc;ng minh</h3>\r\n\r\n<p>M&aacute;y sở hữu 4 camera với camera ch&iacute;nh c&oacute; độ ph&acirc;n giải l&ecirc;n đến 50 MP, camera g&oacute;c rộng 5 MP,&nbsp;cảm biến x&oacute;a ph&ocirc;ng&nbsp;v&agrave;&nbsp;macro c&oacute; c&ugrave;ng độ ph&acirc;n giải 2 MP, k&egrave;m với đ&oacute; l&agrave; nhiều t&iacute;nh năng chụp ảnh mới lạ gi&uacute;p m&igrave;nh thỏa sức kh&aacute;m ph&aacute;.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113223.jpg\" onclick=\"return false;\"><img alt=\"Trang bị 4 camera - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113223.jpg\" /></a></p>\r\n\r\n<p>Kh&aacute; ấn tượng về kết quả thu lại tr&ecirc;n bức h&igrave;nh m&agrave; m&igrave;nh c&oacute; chụp từ điện thoại khi đang di chuyển ngo&agrave;i đường, m&agrave;u sắc ảnh c&oacute; độ tương phản cao, c&aacute;c chi tiết nhỏ đều được m&aacute;y thu lại r&otilde; n&eacute;t, ảnh kh&ocirc;ng qu&aacute; &ldquo;bể&rdquo; khi zoom hay hậu kỳ - chỉnh sửa.&nbsp;</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113221.jpg\" onclick=\"return false;\"><img alt=\"Ảnh chụp ở môi trường đủ sáng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113221.jpg\" /></a></p>\r\n\r\n<p>Về phần chụp đ&ecirc;m th&igrave; Galaxy A23 chưa mang lại kết quả tốt, tổng thể bức ảnh chỉ dừng ở mức chấp nhận được, hiện tượng &ldquo;nh&ograve;e s&aacute;ng&rdquo; ở những vị tr&iacute; nguồn s&aacute;ng cao như b&oacute;ng đ&egrave;n vẫn xuất hiện.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113225.jpg\" onclick=\"return false;\"><img alt=\"Ảnh chụp vào buổi tối - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113225.jpg\" /></a></p>\r\n\r\n<p>Ở chế độ chụp ảnh selfie bằng camera 8 MP cho ra bức h&igrave;nh sắc n&eacute;t, nước da kh&ocirc;ng qu&aacute; bệt, sử dụng t&iacute;nh năng l&agrave;m đẹp đi k&egrave;m để che đi những khuyết điểm li ti như mụn, nốt ruồi nhỏ gi&uacute;p m&igrave;nh tự tin hơn tr&ecirc;n c&aacute;c bức ảnh tự chụp.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113226.jpg\" onclick=\"return false;\"><img alt=\"Ảnh chụp selfie - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113226.jpg\" /></a></p>\r\n\r\n<h3>Thiết kế đặc trưng từ d&ograve;ng Galaxy A</h3>\r\n\r\n<p>Galaxy A23 c&oacute; khung v&agrave; mặt lưng được l&agrave;m từ nhựa mang lại cảm gi&aacute;c cầm nắm nhẹ nh&agrave;ng, c&ugrave;ng với đ&oacute; l&agrave; cạnh viền bo cong gi&uacute;p m&igrave;nh sử dụng l&acirc;u d&agrave;i m&agrave; kh&ocirc;ng cảm thấy bị cấn tay như tr&ecirc;n một số d&ograve;ng sản phẩm c&oacute; thiết kế vu&ocirc;ng vức.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113228.jpg\" onclick=\"return false;\"><img alt=\"Thiết kế hoàn toàn từ nhựa - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113228.jpg\" /></a></p>\r\n\r\n<p>Sở hữu một mặt lưng s&aacute;ng b&oacute;ng c&ugrave;ng m&agrave;u sắc trẻ trung gi&uacute;p m&igrave;nh trở n&ecirc;n nổi bật hơn khi cầm chiếc m&aacute;y tr&ecirc;n tay, tuy nhi&ecirc;n theo m&igrave;nh đ&acirc;y cũng l&agrave; một điểm hạn chế bởi n&oacute; vẫn xuất hiện t&igrave;nh trạng b&aacute;m dấu v&acirc;n tay.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113230.jpg\" onclick=\"return false;\"><img alt=\"Mặt lưng sáng bóng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113230.jpg\" /></a></p>\r\n\r\n<p>Tr&ecirc;n cạnh viền được bố tr&iacute; ph&iacute;m nguồn t&iacute;ch hợp cảm biến v&acirc;n tay với tốc độ phản hồi kh&aacute; nhanh, vị tr&iacute; đặt của ph&iacute;m cũng kh&ocirc;ng qu&aacute; cao gi&uacute;p m&igrave;nh dễ d&agrave;ng k&iacute;ch hoạt thiết bị một c&aacute;ch nhanh ch&oacute;ng chỉ với một tay.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113232.jpg\" onclick=\"return false;\"><img alt=\"Cảm biến vân tay cạnh viền - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113232.jpg\" /></a></p>\r\n\r\n<p>Mặt trước l&agrave; m&agrave;n h&igrave;nh giọt nước, trang bị tấm nền PLS TFT LCD với k&iacute;ch thước 6.6 inch c&oacute; độ ph&acirc;n giải Full HD+ (1080 x 2408 Pixels) cho ra m&agrave;u sắc c&oacute; độ tương phản cao, h&igrave;nh ảnh t&aacute;i hiện chi tiết c&ugrave;ng một kh&ocirc;ng gian hiển thị rộng lớn hỗ trợ cho c&aacute;c t&aacute;c vụ học tập, l&agrave;m việc tốt hơn.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113233.jpg\" onclick=\"return false;\"><img alt=\"Màn hình kích thước lớn - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113233.jpg\" /></a></p>\r\n\r\n<p>M&agrave;n h&igrave;nh 90 Hz cũng l&agrave; một điểm s&aacute;ng tr&ecirc;n Galaxy A23 v&igrave; n&oacute; gi&uacute;p m&igrave;nh thao t&aacute;c c&ocirc;ng việc một c&aacute;ch mượt m&agrave; v&agrave; đ&atilde; mắt hơn, c&ugrave;ng với đ&oacute; l&agrave; t&iacute;nh năng tự động điều chỉnh tần số qu&eacute;t xuống mức mặc định (60 Hz) nhằm tiết kiệm điện năng cho những t&aacute;c vụ kh&ocirc;ng cần thiết như đọc văn bản, đ&agrave;m thoại rất hữu &iacute;ch.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113235.jpg\" onclick=\"return false;\"><img alt=\"Màn hình  tần số 90 Hz - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113235.jpg\" /></a></p>\r\n\r\n<h3>Pin tr&acirc;u d&ugrave;ng l&acirc;u cả ng&agrave;y</h3>\r\n\r\n<p>B&ecirc;n trong thiết bị l&agrave; vi&ecirc;n pin c&oacute; dung lượng 5000 mAh thừa sức đ&aacute;p ứng nhu cầu sử dụng li&ecirc;n tục trong nhiều giờ, m&igrave;nh c&oacute; d&ugrave;ng m&aacute;y để phục vụ cho việc lướt web, xem phim, chơi game cho thấy Galaxy A23 đ&aacute;p ứng thời lượng l&ecirc;n đến hơn 8 tiếng*.</p>\r\n\r\n<p><a href=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113237.jpg\" onclick=\"return false;\"><img alt=\"Thời lượng sử dụng - Samsung Galaxy A23 4GB\" src=\"https://cdn.tgdd.vn/Products/Images/42/262650/samsung-galaxy-a23-160422-113237.jpg\" /></a></p>\r\n\r\n<p>Samsung trang bị cho m&aacute;y c&ocirc;ng nghệ sạc pin nhanh 25 W nhằm tối ưu thời gian sạc gi&uacute;p r&uacute;t ngắn thời gian lấp đầy vi&ecirc;n pin xuống c&ograve;n 1 giờ 35 ph&uacute;t*, đ&acirc;y l&agrave; một khoảng thời gian kh&aacute; hợp l&yacute;, kh&ocirc;ng qu&aacute; l&acirc;u.</p>\r\n',1,'2022-08-13 14:35:42','2022-08-13 17:10:59');
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
INSERT INTO `product_attribute_values` VALUES (1,'Công nghệ màn hình OLED LPTO','Công nghệ màn hình','OLED LPTO',1,'2022-08-13 13:59:52','2022-08-13 13:59:52'),(2,' Độ phân giải  FHD+ (2640 x 1080 Pixels)',' Độ phân giải','FHD+ (2640 x 1080 Pixels)',1,'2022-08-13 14:01:43','2022-08-13 14:01:43'),(3,' Độ sáng tối đa  1200 nits',' Độ sáng tối đa','1200 nits',1,'2022-08-13 14:02:20','2022-08-13 14:02:20'),(4,'Màn hình PLS TFT LCD 6.6\" Full HD+','Màn hình','PLS TFT LCD 6.6\" Full HD+',1,'2022-08-13 14:06:09','2022-08-13 14:06:09'),(5,'Hệ điều hành Android 12','Hệ điều hành','Android 12',0,'2022-08-13 14:06:36','2022-08-13 14:06:36'),(6,'Camera sau Chính 50 MP & Phụ 5 MP, 2 MP, 2 MP','Camera sau','Chính 50 MP & Phụ 5 MP, 2 MP, 2 MP',0,'2022-08-13 14:06:56','2022-08-13 14:06:56'),(7,'Camera trước 8 MP','Camera trước','8 MP',1,'2022-08-13 14:31:29','2022-08-13 14:31:29'),(8,'  Chip Snapdragon 680 8 nhân','Chip','Snapdragon 680 8 nhân',1,'2022-08-13 14:31:49','2022-08-13 14:31:49'),(9,'RAM ','RAM','4 GB',1,'2022-08-13 14:32:12','2022-08-13 14:32:12'),(10,'Bộ nhớ trong 128 GB','Bộ nhớ trong','128 GB',0,'2022-08-13 14:32:37','2022-08-13 14:32:37'),(11,'SIM 2 Nano SIM Hỗ trợ 4G','SIM','2 Nano SIM Hỗ trợ 4G',1,'2022-08-13 14:33:02','2022-08-13 14:33:02'),(12,'Pin, Sạc 5000 mAh 25 W','Pin, Sạc','5000 mAh 25 W',1,'2022-08-13 14:33:24','2022-08-13 14:33:24');
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
INSERT INTO `product_item_colors` VALUES (1,1,'Đen','#000000','5690000','0',5690000,1,'2022-08-13 19:06:14','2022-08-13 19:06:14'),(3,1,'Cam','#ff8040','5690000','9',569,1,'2022-08-13 19:14:02','2022-08-13 19:14:02');
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
