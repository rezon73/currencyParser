-- MySQL dump 10.17  Distrib 10.3.18-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: quote_collector
-- ------------------------------------------------------
-- Server version	10.3.18-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actual_quote`
--

DROP TABLE IF EXISTS `actual_quote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `actual_quote` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `symbol_id` int(11) NOT NULL,
  `exchange_id` int(11) NOT NULL,
  `price` decimal(16,6) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `symbol_id_2` (`symbol_id`,`exchange_id`),
  KEY `symbol_id` (`symbol_id`),
  KEY `exchange_id` (`exchange_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actual_quote`
--

LOCK TABLES `actual_quote` WRITE;
/*!40000 ALTER TABLE `actual_quote` DISABLE KEYS */;
/*!40000 ALTER TABLE `actual_quote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `symbol`
--

DROP TABLE IF EXISTS `symbol`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `symbol` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) NOT NULL,
  `parent_id` int(11) NOT NULL DEFAULT 0,
  `exchange_id` int(11) DEFAULT 0,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `parent_id` (`parent_id`,`exchange_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `symbol`
--

LOCK TABLES `symbol` WRITE;
/*!40000 ALTER TABLE `symbol` DISABLE KEYS */;
INSERT INTO `symbol` VALUES (1,'BTCUSD',0,0,0),(2,'BTCUSDT',1,1,0),(6,'BTC_USD',1,2,0),(7,'ETHBTC',0,0,0),(8,'ETH_BTC',7,2,0);
/*!40000 ALTER TABLE `symbol` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-12-02 19:17:25
