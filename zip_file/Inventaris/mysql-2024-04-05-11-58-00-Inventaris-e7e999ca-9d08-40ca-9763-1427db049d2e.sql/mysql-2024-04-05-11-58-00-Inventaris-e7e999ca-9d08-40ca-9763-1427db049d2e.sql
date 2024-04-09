/*
SQLyog Community v12.2.5 (32 bit)
MySQL - 10.4.32-MariaDB : Database - inventaris
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`inventaris` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;

USE `inventaris`;

/*Table structure for table `barang` */

DROP TABLE IF EXISTS `barang`;

CREATE TABLE `barang` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nama` varchar(255) DEFAULT NULL,
  `Jumlah` int(11) DEFAULT NULL,
  `Harga` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `barang` */

insert  into `barang`(`ID`,`Nama`,`Jumlah`,`Harga`) values 

(1,'Buku',39,'27.65'),

(2,'Pensil',95,'701.54'),

(3,'Baju',63,'80.38'),

(4,'Celana',49,'213.33'),

(5,'Sepatu',59,'329.92'),

(6,'Topi',86,'348.47'),

(7,'Kacamata',13,'651.72'),

(8,'Jam Tangan',84,'246.50'),

(9,'Tas',71,'817.93'),

(10,'Botol Minum',95,'322.25'),

(11,'Piring',74,'761.42'),

(12,'Gelas',57,'566.73'),

(13,'Sendok',12,'915.11'),

(14,'Garpu',20,'286.55'),

(15,'Kompor Gas',81,'208.45'),

(16,'Panci',60,'381.56'),

(17,'Spatula',10,'373.66'),

(18,'Meja',55,'665.45'),

(19,'Kursi',65,'280.54'),

(20,'Lemari',43,'341.25'),

(21,'Lampu',39,'958.97'),

(22,'Karpet',60,'149.10'),

(23,'Bantal',92,'202.24'),

(24,'Guling',22,'500.56'),

(25,'Speaker',83,'690.51'),

(26,'Headphone',93,'614.60'),

(27,'Mouse',26,'462.89'),

(28,'Keyboard',53,'263.52'),

(29,'Monitor',72,'840.89'),

(30,'CPU',2,'605.15'),

(31,'Laptop',94,'929.88'),

(32,'Smartphone',80,'221.39'),

(33,'Tablet',70,'837.40'),

(34,'Printer',8,'918.21'),

(35,'Scanner',33,'909.43'),

(36,'Speaker Bluetooth',54,'15.32'),

(37,'Power Bank',43,'106.14'),

(38,'Charger',23,'878.88'),

(39,'USB Flash Drive',67,'743.49'),

(40,'Hard Disk Eksternal',68,'216.34'),

(41,'Memory Card',1,'419.04'),

(42,'Tas Laptop',5,'14.78'),

(43,'Stiker Laptop',91,'509.45'),

(44,'Kabel HDMI',81,'546.48'),

(45,'Kabel USB',28,'798.08'),

(46,'Baterai',12,'245.26'),

(47,'Mousepad',84,'477.22'),

(48,'Webcam',85,'859.75'),

(49,'Microphone',72,'40.35');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
