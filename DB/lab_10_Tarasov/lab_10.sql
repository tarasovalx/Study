USE master
GO

DROP DATABASE IF EXISTS LAB10;
GO
CREATE DATABASE LAB10 ON(
    NAME=LAB10_dat,
    FILENAME='C:\data\lab10bd.mdf',
    SIZE=10,
    MAXSIZE=UNLIMITED,
    FILEGROWTH=5%
)
LOG ON(
    NAME=LAB10_log,
    FILENAME='C:\data\lab10log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)
GO

USE LAB10
GO

DROP TABLE IF EXISTS TestTable
GO
CREATE TABLE TestTable(
    id INT IDENTITY(1,1) PRIMARY KEY,
    ColChar NVARCHAR(12) NOT NULL UNIQUE,
    ColInt INT DEFAULT 0,
)

INSERT INTO TestTable(ColChar, ColInt) VALUES
('Name_A', 100),
('Name_B', 200)

GO
