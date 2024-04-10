/*
SQLyog Community v12.2.5 (32 bit)
MySQL - 10.4.32-MariaDB : Database - core
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`core` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;

USE `core`;

/*Table structure for table `auto_res` */

DROP TABLE IF EXISTS `auto_res`;

CREATE TABLE `auto_res` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nama_database` varchar(191) NOT NULL,
  `nama_file_backup` longtext NOT NULL,
  PRIMARY KEY (`id`,`nama_database`),
  KEY `idx_auto_res_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `auto_res` */

insert  into `auto_res`(`created_at`,`updated_at`,`deleted_at`,`id`,`nama_database`,`nama_file_backup`) values 
('2024-04-06 07:44:18.847','2024-04-06 07:44:18.847',NULL,1,'Inventaris','mysql-2024-04-05-11-58-00-Inventaris-e7e999ca-9d08-40ca-9763-1427db049d2e.sql.zip'),
('2024-04-06 07:44:18.850','2024-04-06 07:44:18.850',NULL,2,'Keuangan','mysql-2024-04-05-11-21-00-Keuangan-6b8067ef-ec81-4c04-900d-e959fa6c33ae.sql.zip'),
('2024-04-06 07:44:18.851','2024-04-06 07:44:18.851',NULL,3,'PemesananMakanan','mysql-2024-04-05-14-15-00-PemesananMakanan-07c70105-3397-40ba-96b7-ba02b4958e0f.sql.zip'),
('2024-04-06 07:44:18.854','2024-04-06 07:44:18.854',NULL,4,'Pendidikan','mysql-2024-04-05-12-18-48-Pendidikan-49e597-30c1-4e52-92a1-cf8e918499be.sql.zip'),
('2024-04-06 07:44:18.856','2024-04-06 07:44:18.856',NULL,5,'TokoBuku','mysql-2024-04-05-13-10-TokoBuku-fd37150b-5f94-477a-bb2b-dba45ad66cc3.sql.zip'),
('2024-04-10 07:28:37.089','2024-04-10 07:28:37.089',NULL,6,'Inventaris','mysql-2024-04-06-11-58-00-Inventaris-7ef2777b-67d3-4019-b6d8-605109a8ca14.sql'),
('2024-04-10 07:28:37.127','2024-04-10 07:28:37.127',NULL,7,'Inventaris','mysql-2024-04-07-12-58-00-Inventaris-ebf75de9-3438-4eef-a743-eec57637e990.sql'),
('2024-04-10 07:43:06.497','2024-04-10 07:43:06.497',NULL,8,'TokoBuku','mysql-2024-04-05-14-30-TokoBuku-beae61de-dc98-4c67-8ac7-d26e2235caa5.sql'),
('2024-04-10 07:43:06.500','2024-04-10 07:43:06.500',NULL,9,'TokoBuku','mysql-2024-04-05-14-45-TokoBuku-5d3f8989-c6ed-460f-85f9-4a2eb84ae694.sql');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
