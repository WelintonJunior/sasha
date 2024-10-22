CREATE DATABASE  IF NOT EXISTS `sacha` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `sacha`;
-- MySQL dump 10.13  Distrib 8.0.38, for Win64 (x86_64)
--
-- Host: localhost    Database: sacha
-- ------------------------------------------------------
-- Server version	8.0.39

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cha_chamada`
--

DROP TABLE IF EXISTS `cha_chamada`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cha_chamada` (
  `cha_id` int NOT NULL AUTO_INCREMENT,
  `cha_dis_disciplina` int NOT NULL,
  `cha_id_professor` varchar(8) NOT NULL,
  PRIMARY KEY (`cha_id`),
  KEY `cha_dis_disciplina` (`cha_dis_disciplina`),
  KEY `cha_id_professor` (`cha_id_professor`),
  CONSTRAINT `cha_chamada_ibfk_1` FOREIGN KEY (`cha_dis_disciplina`) REFERENCES `dis_disciplina` (`dis_id`),
  CONSTRAINT `cha_chamada_ibfk_2` FOREIGN KEY (`cha_id_professor`) REFERENCES `usu_users` (`usu_ra`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cha_chamada`
--

LOCK TABLES `cha_chamada` WRITE;
/*!40000 ALTER TABLE `cha_chamada` DISABLE KEYS */;
INSERT INTO `cha_chamada` VALUES (9,1,''),(12,1,'');
/*!40000 ALTER TABLE `cha_chamada` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `det_detalhesdachamada`
--

DROP TABLE IF EXISTS `det_detalhesdachamada`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `det_detalhesdachamada` (
  `det_id` int NOT NULL AUTO_INCREMENT,
  `det_id_chamada` int DEFAULT NULL,
  `det_presente` tinyint(1) DEFAULT NULL,
  `det_ra_aluno` varchar(8) NOT NULL,
  PRIMARY KEY (`det_id`),
  KEY `det_id_chamada` (`det_id_chamada`),
  KEY `det_ra_aluno` (`det_ra_aluno`),
  CONSTRAINT `det_detalhesdachamada_ibfk_1` FOREIGN KEY (`det_id_chamada`) REFERENCES `cha_chamada` (`cha_id`),
  CONSTRAINT `det_detalhesdachamada_ibfk_2` FOREIGN KEY (`det_ra_aluno`) REFERENCES `usu_users` (`usu_ra`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `det_detalhesdachamada`
--

LOCK TABLES `det_detalhesdachamada` WRITE;
/*!40000 ALTER TABLE `det_detalhesdachamada` DISABLE KEYS */;
INSERT INTO `det_detalhesdachamada` VALUES (3,9,1,'');
/*!40000 ALTER TABLE `det_detalhesdachamada` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dis_disciplina`
--

DROP TABLE IF EXISTS `dis_disciplina`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dis_disciplina` (
  `dis_id` int NOT NULL AUTO_INCREMENT,
  `dis_nome` varchar(50) NOT NULL,
  PRIMARY KEY (`dis_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dis_disciplina`
--

LOCK TABLES `dis_disciplina` WRITE;
/*!40000 ALTER TABLE `dis_disciplina` DISABLE KEYS */;
INSERT INTO `dis_disciplina` VALUES (1,'GGTI');
/*!40000 ALTER TABLE `dis_disciplina` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usu_users`
--

DROP TABLE IF EXISTS `usu_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usu_users` (
  `password` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `type` char(1) NOT NULL,
  `usu_ra` varchar(8) NOT NULL,
  PRIMARY KEY (`usu_ra`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usu_users`
--

LOCK TABLES `usu_users` WRITE;
/*!40000 ALTER TABLE `usu_users` DISABLE KEYS */;
INSERT INTO `usu_users` VALUES ('123123','Welinton','1','');
/*!40000 ALTER TABLE `usu_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-22 19:41:30
