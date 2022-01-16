USE master;
GO

DROP DATABASE IF EXISTS LAB5
GO

CREATE DATABASE LAB5 ON(
    NAME=LAB5_dat,
    FILENAME='C:\data\lab5bd.mdf',
    SIZE=10,
    MAXSIZE=100MB,
    FILEGROWTH=5%
)

LOG ON(
    NAME=LAB5_log,
    FILENAME='C:\data\lab5log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)
GO

USE LAB5;
GO

DROP TABLE IF EXISTS [User]

CREATE TABLE [User](
    UserName varchar(20),
    DateOfBirth date,
)

INSERT INTO [User](UserName, DateOfBirth) VALUES ('Alexey','1999-12-06')
GO


ALTER DATABASE LAB5 ADD FILEGROUP LAB5_filegroup
GO

ALTER DATABASE LAB5 ADD FILE(
    NAME=test,
    FILENAME='C:\data\test.ndf',
    SIZE=10MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
) TO FILEGROUP LAB5_filegroup

GO
ALTER DATABASE LAB5 MODIFY FILEGROUP LAB5_filegroup DEFAULT;
GO

USE LAB5;
GO


DROP TABLE IF EXISTS Playlist
GO

CREATE TABLE Playlist(
    PlaylistID int PRIMARY KEY,
    PlaylistName VARCHAR(20)
)

INSERT INTO Playlist(PlaylistID, PlaylistName) VALUES (1, 'Classic')
GO

ALTER DATABASE LAB5
	MODIFY FILEGROUP [PRIMARY] DEFAULT;
GO;

DROP TABLE Playlist;
ALTER DATABASE LAB5 REMOVE FILE test
ALTER DATABASE LAB5 REMOVE FILEGROUP LAB5_filegroup
GO

USE LAB5;
GO
CREATE SCHEMA TestSchema;
GO

ALTER SCHEMA TestSchema TRANSFER dbo.[User];
DROP TABLE TestSchema.[User]

DROP SCHEMA TestSchema
GO
