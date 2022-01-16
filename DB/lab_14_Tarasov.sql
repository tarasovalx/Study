USE LAB13_1

DROP TABLE IF EXISTS Song
CREATE TABLE Song(
	SongId INT PRIMARY KEY,
	SongName NVARCHAR(60) NOT NULL,
	Duration INT,
	AgeLimitation bit DEFAULT 0,

	--FilePath NVARCHAR(200)
)

GO

USE LAB13_2

DROP TABLE IF EXISTS Song
CREATE TABLE Song(
	SongId INT PRIMARY KEY,
	--SongName NVARCHAR(60) NOT NULL,
	--Duration Float,
	--AgeLimitation bit DEFAULT 0,

	FilePath NVARCHAR(200)
)
GO

USE LAB13_3

DROP VIEW IF EXISTS SongView
GO

SELECT * FROM LAB13_1.dbo.Song
GO
CREATE VIEW SongView AS
    SELECT A.*, B.FilePath
		FROM LAB13_1.dbo.Song as A, LAB13_2.dbo.Song as B
		WHERE A.SongId = B.SongId
GO

DROP TRIGGER IF EXISTS onInsertTrigger
GO

CREATE TRIGGER onInsertTrigger ON SongView
INSTEAD OF INSERT
AS
BEGIN
    IF (EXISTS (SELECT SongId FROM LAB13_1.dbo.Song INTERSECT SELECT SongId FROM inserted))
	BEGIN
        RAISERROR('Row with such id is already exists', -1, 1)
		RETURN
    END

    INSERT INTO LAB13_1.dbo.Song SELECT SongId, SongName, Duration, AgeLimitation FROM inserted
    INSERT INTO LAB13_2.dbo.Song SELECT SongId, FilePath FROM inserted
END
GO

DROP TRIGGER IF EXISTS onDeleteTrigger
GO
CREATE TRIGGER onDeleteTrigger ON SongView
INSTEAD OF DELETE
AS
BEGIN
	--DELETE FROM LAB13_1.dbo.Song 
	--	WHERE (LAB13_1.dbo.Song.SongID IN (SELECT SongId FROM deleted))


    DELETE S FROM LAB13_1.dbo.Song AS S INNER JOIN deleted AS D ON S.SongId = D.SongId
    DELETE S FROM LAB13_2.dbo.Song AS S INNER JOIN deleted AS D ON S.SongId = D.SongId
END
GO

DROP TRIGGER IF EXISTS onUpdateTrigger
GO
CREATE TRIGGER onUpdateTrigger ON SongView
INSTEAD OF UPDATE
AS
BEGIN
    IF (UPDATE(SongId)) 
	    RAISERROR('Can not update PK',-1, 1)

    UPDATE LAB13_1.dbo.Song SET SongName = I.SongName, Duration=I.Duration, AgeLimitation = I.AgeLimitation FROM LAB13_1.dbo.Song S INNER JOIN inserted I ON S.SongId=I.SongId
    UPDATE LAB13_2.dbo.Song SET FilePath=I.FilePath FROM LAB13_1.dbo.Song S INNER JOIN inserted I ON S.SongId=i.SongId
END
GO

SELECT * FROM SongView

INSERT INTO SongView VALUES
(1,'Name1', 120, 0, 'path/Name1'),
(2,'Name2', 120, 0, 'path/Name2'),
(3,'Name3', 120, 0, 'path/Name2'),
(4,'Name4', 120, 0, 'path/Name3'),
(5,'Name5', 120, 1, 'path/Name4')

SELECT * FROM SongView

INSERT INTO SongView VALUES
(1,'Name1', 120, 1, 'path/Name1')

UPDATE SongView SET SongName = 'Prefix'+ SongName WHERE AgeLimitation = 1

DELETE FROM SongView WHERE SongName LIKE('Prefix%')

SELECT * FROM SongView