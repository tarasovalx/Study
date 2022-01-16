USE LAB6
GO

DROP TRIGGER IF EXISTS onInsertToSongTrigger 
DROP TRIGGER IF EXISTS onDeleteFromSongTrigger
DROP TRIGGER IF EXISTS onUpdateSongTrigger
GO


CREATE TRIGGER onInsertToSongTrigger ON Song
AFTER INSERT
AS
BEGIN
    IF EXISTS(SELECT * FROM inserted as S1
			  INNER JOIN Song S2 on ((S1.AritstId = S2.AritstId)
									AND (S1.SongId != S2.SongId)
									AND  (S1.SongName = S2.SongName)))
		BEGIN
			PRINT 'INVALID NAMES IN BATCH'
			ROLLBACK TRANSACTION
		END
    RETURN
	
END
GO


SELECT * FROM Song

INSERT INTO Song(SongName, AritstId) VALUES
('NewSong', 2),
('NewSong', 2),
('KEK', 2)
GO

CREATE TRIGGER onUpdateSongTrigger ON Song
AFTER UPDATE AS
BEGIN
    IF EXISTS(SELECT * FROM inserted as S1
			  INNER JOIN Song S2 on ((S1.AritstId = S2.AritstId)
									AND (S1.SongId != S2.SongId)
									AND  (S1.SongName = S2.SongName)))
		BEGIN
			RAISERROR(N'INVALID NAMES IN BATCH', -1, 20);
			ROLLBACK TRANSACTION;
		END
    RETURN
END
GO

UPDATE Song
SET SongName = 'Bohemian rapsody'
WHERE SongName = 'Do not stop me now' 
GO


CREATE TRIGGER onDeleteFromSongTrigger ON Song
AFTER DELETE
AS
BEGIN
	IF (SELECT COUNT(*) FROM Song) = 0
		PRINT 'Last Song deleted'
	RETURN
END
GO

SELECT * FROM Song


SELECT SongID, SongName, ArtistName FROM Song 
	LEFT JOIN Artist ON Song.AritstId = Artist.ArtistId
GO



DROP VIEW IF EXISTS SongsWithArtistView
GO
CREATE VIEW SongsWithArtistView AS
	SELECT SongID, SongName, ArtistId, ArtistName FROM Song 
		LEFT JOIN Artist ON Song.AritstId = Artist.ArtistId
GO

SELECT * FROM SongsWithArtistView


DROP TRIGGER IF EXISTS onInsertSongsWithArtistViewTrigger 
DROP TRIGGER IF EXISTS onDeleteSongsWithArtistViewTrigger
DROP TRIGGER IF EXISTS onUpdateSongsWithArtistViewTrigger
GO

CREATE TRIGGER onInsertSongsWithArtistViewTrigger ON  SongsWithArtistView
INSTEAD OF INSERT
AS
BEGIN
	/*
	DECLARE MY_CURSOR CURSOR LOCAL STATIC READ_ONLY FORWARD_ONLY FOR 
	SELECT SongName, ArtistName FROM inserted
	OPEN MY_CURSOR
	
	DECLARE @ArtistName NVARCHAR(60), @SongName NVARCHAR(60)

	FETCH NEXT FROM MY_CURSOR INTO @SongName, @ArtistName
	WHILE @@FETCH_STATUS = 0
	BEGIN 
		 PRINT @SongName + ' ' + @ArtistName
		 INSERT INTO Artist(ArtistName) VALUES (@ArtistName)
		 INSERT INTO Song(SongName, AritstId) VALUES (@SongName, SCOPE_IDENTITY())
	     FETCH NEXT FROM MY_CURSOR INTO @SongName, @ArtistName
	END
	*/
	--INSERT INTO Artist(ArtistName)  
	INSERT INTO Artist(ArtistName) (SELECT DISTINCT ArtistName FROM inserted WHERE ArtistName NOT IN (SELECT ArtistName FROM Artist)) 
	INSERT INTO Song(SongName, AritstId) (SELECT T.SongName, T.ArtistId FROM 
		(SELECT SongName, Artist.ArtistId as ArtistId FROM inserted LEFT JOIN Artist on inserted.ArtistName = Artist.ArtistName) AS T)
	RETURN
END
GO

CREATE TRIGGER onUpdateSongsWithArtistViewTrigger ON  SongsWithArtistView
INSTEAD OF UPDATE
AS
BEGIN
	IF UPDATE(SongName)
		UPDATE Song SET SongName = A.SongName FROM (SELECT SongName, SongId FROM inserted) AS A
		WHERE (Song.SongId = A.SongId)
	IF UPDATE(ArtistName)
		UPDATE Artist SET ArtistName = A.ArtistId FROM (SELECT ArtistName, ArtistId FROM inserted) AS A
		WHERE (Artist.ArtistId = A.ArtistId)
END
GO

CREATE TRIGGER onDeleteSongsWithArtistViewTrigger ON  SongsWithArtistView
INSTEAD OF DELETE
AS
BEGIN
	DELETE FROM Song WHERE SongId in (SELECT SongId FROM deleted)
END
GO

INSERT INTO SongsWithArtistView(SongName, ArtistName) VALUES 
('Hush', 'DeepPurple'), ('Wall', 'Pink Phloyd')
	
UPDATE SongsWithArtistView SET SongName ='WALL' WHERE SongName = 'Another brick in the wall'
--('Hush', 'Deep Purple'), ('Wall', 'Pink Phloyd')
	

SELECT * FROM SongsWithArtistView
SELECT * FROM Artist 
SELECT * FROM Song

DELETE FROM ARTIST WHERE ArtistName in ('Pink Phloyd', 'Deep Purple')

