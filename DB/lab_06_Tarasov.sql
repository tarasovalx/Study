USE master;
GO

DROP DATABASE IF EXISTS LAB6;
GO

CREATE DATABASE LAB6 ON(
    NAME=LAB6_dat,
    FILENAME='C:\data\lab6bd.mdf',
    SIZE=10,
    MAXSIZE=UNLIMITED,
    FILEGROWTH=5%
)
LOG ON(
    NAME=LAB6_log,
    FILENAME='C:\data\lab6log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)
GO

USE LAB6;

DROP TABLE IF EXISTS [User]
CREATE TABLE [User](
	UserID INT IDENTITY(1,1) PRIMARY KEY,
	UserName NVARCHAR(60) NOT NULL CHECK (LEN(UserName) > 6),
	DateOfBirth Date CHECK (DateOfBirth < DATEADD(year, -12, GETDATE())) DEFAULT DATEADD(year, -12, GETDATE())
)

INSERT INTO [User] (UserName) VALUES
--(N'test'),
(N'GoodName')

INSERT INTO [User] (UserName, DateOfBirth) VALUES
(N'Alexey Tarasov', '1999-12-06'),
(N'Plaunov Stepan', '2000-10-23')



DROP TABLE IF EXISTS Artist

CREATE TABLE Artist(
    ArtistId UNIQUEIDENTIFIER PRIMARY KEY DEFAULT(NEWID()),
    ArtistName NVARCHAR(60) NOT NULL,
	Country NVARCHAR(50),
)

INSERT INTO Artist(ArtistName, Country) VALUES
('DavidBowie', 'UK'), ('Queen', 'UK'), ('DeepPurple', 'UK')

DROP SEQUENCE IF EXISTS IdSequence 

CREATE SEQUENCE IdSequence 
    START WITH 1
    INCREMENT BY 1
    MAXVALUE 50


DROP TABLE IF EXISTS Song

CREATE TABLE Song(
	SongId INT PRIMARY KEY,
	SongName NVARCHAR(60) NOT NULL,
)

INSERT INTO Song(SongId, SongName) VALUES
(NEXT VALUE FOR IdSequence, 'Life on Mars?'),
(NEXT VALUE FOR IdSequence, 'Bohemian rapsody'),
(NEXT VALUE FOR IdSequence, 'Hush')



DROP TABLE IF EXISTS Artist

CREATE TABLE Artist(
    ArtistId INT IDENTITY(1,1) PRIMARY KEY,
    ArtistName NVARCHAR(60) NOT NULL,
	Country NVARCHAR(50),
)

INSERT INTO Artist(ArtistName, Country) VALUES
('DavidBowie', 'UK'), ('Queen', 'UK'), ('DeepPurple', 'UK'), ('Ray Charles', 'US')

DROP TABLE IF EXISTS Song

CREATE TABLE Song(
	SongId INT IDENTITY(1,1) PRIMARY KEY,
	SongName NVARCHAR(60) NOT NULL,
	AritstId INT,
	CONSTRAINT ArtistFK FOREIGN KEY (AritstId) REFERENCES Artist(ArtistId)
	--ON UPDATE SET NULL
    --ON UPDATE NO ACTION
    --ON UPDATE SET DEFAULT
	--ON UPDATE CASCADE

    --ON DELETE SET NULL
    --ON DELETE NO ACTION
    --ON DELETE SET DEFAULT
    --ON DELETE CASCADE
)	

INSERT INTO Song(SongName, AritstId) VALUES
('Space Oddity', 1), ('Life on Mars?', 1), ('Bohemian rapsody', 2), ('Do not stop me now', 2), ('Hit the road Jack', 4)




SELECT * FROM Artist
SELECT * FROM Song 

DELETE FROM Artist WHERE ArtistId=1

SELECT * FROM Artist
SELECT * FROM Song 

--SCOPE_IDENTITY()
--IDENT_CURRENT('dbo.Song')
--@@IDENTITY
