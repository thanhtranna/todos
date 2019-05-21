-- MySQL dump 10.13  Distrib 5.7.20, for Linux (x86_64)
--
-- Host: 0.0.0.0    Database: todos
-- ------------------------------------------------------
-- Server version	5.7.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `todos`
--

DROP TABLE IF EXISTS `todos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `todos` (
  `id` varchar(60) NOT NULL,
  `user_id` varchar(60) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `content` mediumtext,
  PRIMARY KEY (`id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `todos`
--

LOCK TABLES `todos` WRITE;
/*!40000 ALTER TABLE `todos` DISABLE KEYS */;
INSERT INTO `todos` VALUES ('21bdd2b3-24b0-460c-a9e8-fecb4885a4f8','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:02','2019-05-22 01:07:02','this is a test content1'),('29ccc4eb-26a0-4fd3-8a13-bf9a36229397','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:04','2019-05-22 01:07:04','this is a test content2'),('2c98d10c-716f-4007-8d20-63c418041c23','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:04','2019-05-22 01:07:04','this is a test content3'),('2e784ad1-ac5b-4b81-856f-6047c6eafcad','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:02','2019-05-22 01:07:02','this is a test content4'),('57713d91-c3c1-4484-afb7-2466b1ae7628','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 00:59:27','2019-05-22 00:59:27','this is a test content5'),('68673224-8be1-4f96-9014-e127393732aa','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:01','2019-05-22 01:07:01','this is a test content6'),('69720063-6dba-4632-bd4f-8b1d6694b3bb','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:04','2019-05-22 01:07:04','this is a test content7'),('743eb190-1cd8-4a53-8ed3-9a1511fe1725','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:05','2019-05-22 01:07:05','this is a test content8'),('793640ca-eaff-4bc1-b940-6a710de08c78','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content9'),('7d86f808-9dea-4f78-bbcb-8b89cd8b72ac','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:04','2019-05-22 01:07:04','this is a test content10'),('7f70fc5d-13df-481d-a1b7-78ed6aa5e345','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:02','2019-05-22 01:07:02','this is a test content11'),('85a9cdb5-a8d2-4e23-8413-ba7e2e3cb816','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:02','2019-05-22 01:07:02','this is a test content12'),('a47415d0-2d8c-486a-971a-6976e9af8220','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content13'),('a9e7ee22-ec69-4f3d-8a1d-c83f463920e6','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content14'),('aa99080c-3090-4c33-837a-39c74843b519','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 00:39:42','2019-05-22 00:39:42','this is a test content15'),('ab8bca47-a057-42c0-86d0-740277c59717','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:05','2019-05-22 01:07:05','this is a test content16'),('ce1ef3ae-cd3e-43e0-bd78-a403a603efb6','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:01','2019-05-22 01:07:01','this is a test content17'),('de871861-36c1-4b84-a591-e626cd6805b2','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content18'),('e236a6be-7f62-45f4-ba08-ec096f9f4f62','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:01','2019-05-22 01:07:01','this is a test content19'),('e809fe9d-db49-492a-81a4-b93aa63281cd','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 00:59:27','2019-05-22 00:59:27','this is a test content20'),('e9c43807-a8a1-4712-8a45-9c352e626025','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content21'),('ea416ce1-7cc9-4680-ab38-05b559adba02','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 00:59:26','2019-05-22 00:59:26','this is a test content22'),('ee418b8d-ec9c-4941-9cf7-e543f4f328d4','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:03','2019-05-22 01:07:03','this is a test content23'),('f33c2504-fda6-418a-be2f-f3c8ef9b8233','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:02','2019-05-22 01:07:02','this is a test content24'),('f4fb3f11-908a-4feb-be1f-22b8e64614c1','49efb134-30f4-45f4-81c1-9ba51caa4e7b','2019-05-22 01:07:04','2019-05-22 01:07:04','this is a test content25');
/*!40000 ALTER TABLE `todos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-22  1:18:26
