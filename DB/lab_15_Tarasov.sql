USE master;

DROP DATABASE IF EXISTS LAB15_1
CREATE DATABASE LAB15_1 ON(
    NAME=LAB15_1DB,
    FILENAME='C:\data\lab15_1bd.mdf',
    SIZE=10,
    MAXSIZE=UNLIMITED,
    FILEGROWTH=5%
)

LOG ON(
    NAME=LAB15_1LOG,
    FILENAME='C:\data\lab15_1log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)

USE LAB15_1;
DROP TABLE IF EXISTS Song
CREATE TABLE Song(
	SongID INT PRIMARY KEY,
	SongName NVARCHAR(60) NOT NULL,
	ArtistId INT,
	Duration Float,
	AgeLimitation bit DEFAULT 0,
	FilePath NVARCHAR(200)
)



DROP DATABASE IF EXISTS LAB15_2
CREATE DATABASE LAB15_2 ON(
    NAME=LAB15_2DB,
    FILENAME='C:\data\lab15_2bd.mdf',
    SIZE=10,
    MAXSIZE=UNLIMITED,
    FILEGROWTH=5%
)

LOG ON(
    NAME=LAB15_2LOG,
    FILENAME='C:\data\lab15_2log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)

USE LAB15_2;
DROP TABLE IF EXISTS Artist
CREATE TABLE Artist(
	ArtistId INT PRIMARY KEY,
	ArtistName NVARCHAR(60) NOT NULL,
	Country NVARCHAR(60) NOT NULL
)

DROP TRIGGER IF EXISTS ArtistInsertTrigger
GO
CREATE TRIGGER ArtistInsertTrigger ON Artist
INSTEAD OF INSERT 
AS
BEGIN
	IF (EXISTS (SELECT ArtistId FROM LAB15_2.dbo.Artist INTERSECT SELECT ArtistId FROM inserted))
	BEGIN
        RAISERROR('Row with such id is already exists', -1, 1)
		RETURN
    END
	INSERT INTO LAB15_2.dbo.Artist SELECT * FROM inserted
RETURN
END
GO

DROP TRIGGER IF EXISTS ArtistDeleteTrigger
GO
CREATE TRIGGER ArtistDeleteTrigger ON Artist
INSTEAD OF DELETE 
AS
BEGIN
	DELETE S FROM LAB15_1.dbo.Song AS S INNER JOIN deleted AS D ON S.ArtistId = D.ArtistId
	DELETE S FROM LAB15_2.dbo.Artist AS S INNER JOIN deleted AS D ON S.ArtistId = D.ArtistId
RETURN
END
GO


DROP TRIGGER IF EXISTS ArtistUpdateTrigger
GO
CREATE TRIGGER ArtistUpdateTrigger ON Artist
INSTEAD OF UPDATE 
AS
BEGIN
    IF (UPDATE(ArtistId)) 
	BEGIN
	    RAISERROR('Can not update PK',-1, 1)
		RETURN
	END
	IF (UPDATE(ArtistName))
		UPDATE LAB15_2.dbo.Artist SET ArtistName = inserted.ArtistName FROM inserted WHERE Artist.ArtistId = inserted.ArtistId

	IF (UPDATE(Country))
		UPDATE LAB15_2.dbo.Artist SET Country = inserted.Country FROM inserted WHERE Artist.ArtistId = inserted.ArtistId
	RETURN
END
GO



USE LAB15_1;
--SONG--

DROP TRIGGER IF EXISTS SongInsertTrigger
GO
CREATE TRIGGER SongInsertTrigger ON Song
INSTEAD OF INSERT 
AS
BEGIN
	IF (EXISTS (SELECT * FROM Song as S 
			JOIN LAB15_2.dbo.Artist A ON S.ArtistId = A.ArtistId
			WHERE A.ArtistId IS NULl))
		BEGIN
			RAISERROR('No such ArtistId', -1, 1)
			RETURN
		END
	IF (EXISTS (SELECT SongId FROM LAB15_1.dbo.Song INTERSECT SELECT SongId FROM inserted))
		BEGIN
			RAISERROR('Row with such id is already exists', -1, 1)
			RETURN
		END

	INSERT INTO LAB15_1.dbo.Song SELECT * FROM inserted
	RETURN
END
GO

DROP TRIGGER IF EXISTS SongDeleteTrigger
GO
CREATE TRIGGER SongDeleteTrigger ON Song
INSTEAD OF DELETE 
AS
BEGIN
	DELETE S FROM LAB15_1.dbo.Song AS S INNER JOIN deleted AS D ON S.SongId = D.SongId
	RETURN
END
GO


DROP TRIGGER IF EXISTS SongUpdateTrigger
GO
CREATE TRIGGER SongUpdateTrigger ON Song
INSTEAD OF UPDATE 
AS
BEGIN
    IF (UPDATE(SongId)) 
	BEGIN
	    RAISERROR('Can not update PK',-1, 1)
		RETURN
	END

	IF(UPDATE(ArtistId))
	BEGIN
	    RAISERROR('Can not update FK',-1, 1)
		RETURN
	END

	IF (UPDATE(SongName))
		UPDATE LAB15_1.dbo.Song SET SongName = inserted.SongName FROM inserted WHERE Song.SongID = inserted.SongID

	IF (UPDATE(Duration))
		UPDATE LAB15_1.dbo.Song SET Duration = inserted.Duration FROM inserted WHERE Song.SongID = inserted.SongID

	IF (UPDATE(AgeLimitation))
		UPDATE  LAB15_1.dbo.Song SET AgeLimitation = inserted.AgeLimitation FROM inserted WHERE Song.SongID = inserted.SongID

	IF (UPDATE(FilePath))
		UPDATE  LAB15_1.dbo.Song SET FilePath = inserted.FilePath FROM inserted WHERE Song.SongID = inserted.SongID
	RETURN
END
GO

SELECT * FROM LAB15_2.dbo.Artist

INSERT INTO LAB15_2.dbo.Artist(ArtistId, ArtistName, Country) VALUES
(1, 'Name1', 'UK'),
(2, 'Name1', 'UK')

SELECT * FROM LAB15_2.dbo.Artist

INSERT INTO LAB15_1.dbo.Song(SongID, SongName, ArtistId) VALUES
(1, 'SONG1', 2)
SELECT * FROM LAB15_1.dbo.Song

UPDATE LAB15_2.dbo.Artist SET ArtistName = ('Prefix' + ArtistName) WHERE ArtistName = 'Name1'
SELECT * FROM LAB15_1.dbo.Song

DELETE FROM LAB15_2.dbo.Artist WHERE ArtistID =  2

SELECT * FROM LAB15_2.dbo.Artist
SELECT * FROM LAB15_1.dbo.Song





SELECT * FROM LAB15_2.dbo.Artist