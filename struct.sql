-- --------------------------------------------------------
-- Хост:                         127.0.0.1
-- Версия сервера:               10.5.9-MariaDB - mariadb.org binary distribution
-- Операционная система:         Win64
-- HeidiSQL Версия:              11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Дамп структуры для таблица phones.addr
DROP TABLE IF EXISTS `addr`;
CREATE TABLE IF NOT EXISTS `addr` (
  `idaddr` int(3) NOT NULL AUTO_INCREMENT,
  `addr` varchar(70) NOT NULL,
  `lat` float(10,6) NOT NULL,
  `lng` float(10,6) NOT NULL,
  `postcode` int(6) NOT NULL,
  PRIMARY KEY (`idaddr`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.certs
DROP TABLE IF EXISTS `certs`;
CREATE TABLE IF NOT EXISTS `certs` (
  `idcert` int(11) NOT NULL AUTO_INCREMENT,
  `filename` varchar(50) DEFAULT NULL,
  `startdate` date NOT NULL DEFAULT '2021-04-27',
  `enddate` date NOT NULL DEFAULT '2021-04-27',
  PRIMARY KEY (`idcert`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.depart
DROP TABLE IF EXISTS `depart`;
CREATE TABLE IF NOT EXISTS `depart` (
  `iddep` int(3) NOT NULL AUTO_INCREMENT,
  `depart` varchar(200) NOT NULL,
  `sdep` varchar(60) NOT NULL,
  `email` varchar(50) NOT NULL,
  `idaddr` int(3) NOT NULL,
  `idparent` int(3) NOT NULL DEFAULT 0,
  `abbr` varchar(3) NOT NULL,
  PRIMARY KEY (`iddep`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.docs
DROP TABLE IF EXISTS `docs`;
CREATE TABLE IF NOT EXISTS `docs` (
  `iddoc` int(11) NOT NULL AUTO_INCREMENT,
  `filename` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`iddoc`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.persons
DROP TABLE IF EXISTS `persons`;
CREATE TABLE IF NOT EXISTS `persons` (
  `idperson` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(300) NOT NULL,
  `date` date NOT NULL,
  `file` varchar(50) NOT NULL,
  `cellular` varchar(50) NOT NULL,
  `business` varchar(50) NOT NULL,
  `iddep` int(3) NOT NULL,
  `idpos` int(3) NOT NULL,
  `idrank` int(3) NOT NULL,
  `idrole` int(3) NOT NULL DEFAULT 0,
  `passwd` varchar(45) NOT NULL DEFAULT 'e10adc3949ba59abbe56e057f20f883e',
  PRIMARY KEY (`idperson`)
) ENGINE=InnoDB AUTO_INCREMENT=654 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.places
DROP TABLE IF EXISTS `places`;
CREATE TABLE IF NOT EXISTS `places` (
  `idplace` int(3) NOT NULL AUTO_INCREMENT,
  `place` varchar(10) NOT NULL,
  `work` varchar(18) NOT NULL,
  `internal` varchar(3) NOT NULL,
  `ipphone` varchar(18) NOT NULL,
  `arm` varchar(13) NOT NULL,
  `idperson` int(3) DEFAULT NULL,
  `idaddr` int(3) NOT NULL,
  PRIMARY KEY (`idplace`)
) ENGINE=InnoDB AUTO_INCREMENT=395 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.pos
DROP TABLE IF EXISTS `pos`;
CREATE TABLE IF NOT EXISTS `pos` (
  `idpos` int(3) NOT NULL AUTO_INCREMENT,
  `pos` varchar(300) NOT NULL,
  PRIMARY KEY (`idpos`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.ranks
DROP TABLE IF EXISTS `ranks`;
CREATE TABLE IF NOT EXISTS `ranks` (
  `idrank` int(3) NOT NULL AUTO_INCREMENT,
  `rank` varchar(300) NOT NULL,
  PRIMARY KEY (`idrank`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.role
DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
  `idrole` int(3) NOT NULL,
  `role` varchar(6) NOT NULL,
  PRIMARY KEY (`idrole`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

-- Дамп структуры для таблица phones.tokens
DROP TABLE IF EXISTS `tokens`;
CREATE TABLE IF NOT EXISTS `tokens` (
  `idtoken` int(11) NOT NULL AUTO_INCREMENT,
  `number` varchar(50) DEFAULT NULL,
  `idowner` int(11) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`idtoken`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Экспортируемые данные не выделены.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
