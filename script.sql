CREATE DATABASE `my_db` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_spanish_ci */;

CREATE TABLE `odontologos` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `Apellido` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `Matricula` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

CREATE TABLE `paciente` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `Apellido` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `Domicilio` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `DNI` int(11) NOT NULL,
  `Fecha_Alta` varchar(50) COLLATE utf8_spanish_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

CREATE TABLE `turno` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `DNI_paciente` int(11) NOT NULL,
  `Matricula_odontologo` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `Fecha_Hora` varchar(50) COLLATE utf8_spanish_ci DEFAULT NULL,
  `Descripcion` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `Matricula_odontologo` (`Matricula_odontologo`),
  KEY `DNI_paciente` (`DNI_paciente`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

INSERT INTO odontologos (Nombre, Apellido, Matricula)
VALUES ('Juan', 'Pérez', '12345');
INSERT INTO odontologos (Nombre, Apellido, Matricula)
VALUES ('Vero', 'Aguirre', '6789');
INSERT INTO paciente (Nombre, Apellido, Domicilio, DNI, Fecha_Alta)
VALUES ('Jose', 'Gómez', 'Calle 123', 12345678, '2023-09-19');
INSERT INTO paciente (Nombre, Apellido, Domicilio, DNI, Fecha_Alta)
VALUES ('Ana', 'Gómez', 'Calle 123', 12345678, '2023-09-19');
INSERT INTO turno (DNI_paciente, Matricula_odontologo, Fecha_Hora, Descripcion)
VALUES (12345678, '12345', '2023-09-19 10:00 AM', 'Limpieza dental');
INSERT INTO turno (DNI_paciente, Matricula_odontologo, Fecha_Hora, Descripcion)
VALUES (12345678, '6789', '2023-09-19 10:00 AM', 'Consulta general');