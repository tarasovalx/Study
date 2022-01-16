USE master;
GO

DROP DATABASE IF EXISTS LAB11;
GO
CREATE DATABASE LAB11 ON(
    NAME=LAB11_dat,
    FILENAME='C:\data\lab11bd.mdf',
    SIZE=10,
    MAXSIZE=UNLIMITED,
    FILEGROWTH=5%
)
LOG ON(
    NAME=LAB11_log,
    FILENAME='C:\data\lab11log.ldf',
    SIZE=5MB,
    MAXSIZE=25MB,
    FILEGROWTH=5MB
)
GO

USE LAB11;

DROP TABLE IF EXISTS Playlist_Song_int
DROP TABLE IF EXISTS Album_Song_int
DROP TABLE IF EXISTS Song

DROP TABLE IF EXISTS Playlist
DROP TABLE IF EXISTS Album
DROP TABLE IF EXISTS [User]
DROP TABLE IF EXISTS Artist


CREATE SEQUENCE UserName_Seq
	START WITH 1
    INCREMENT BY 1


CREATE SEQUENCE ArtistName_Seq 
	START WITH 1
    INCREMENT BY 1


CREATE TABLE [User](
	UserId int IDENTITY(1,1) PRIMARY KEY,
	UserName NVARCHAR(60) NOT NULL UNIQUE
		DEFAULT (FORMATMESSAGE('UserName_%s',CAST((NEXT VALUE FOR UserName_Seq) as VARCHAR(max)))),
	PasswordHash NVARCHAR(100),
	DateOfBirth Date CHECK (DateOfBirth < DATEADD(year, -12, GETDATE())) DEFAULT DATEADD(year, -12, GETDATE())
)

CREATE TABLE Playlist(
	PlaylistId INT IDENTITY(1,1) PRIMARY KEY,
	PlaylistName NVARCHAR(60) NOT NULL,
	[Description] NVARCHAR(120),
	UserId INT
	CONSTRAINT Playlist_FK FOREIGN KEY (UserId) REFERENCES [User](UserId)
	ON UPDATE CASCADE
	ON DELETE CASCADE
)


CREATE TABLE Artist(
	ArtistId INT IDENTITY(1,1) PRIMARY KEY,
	ArtistName NVARCHAR(60) NOT NULL UNIQUE DEFAULT ('ArtirstName_' + (NEXT VALUE FOR ArtistName_Seq)),
	About NVARCHAR(120)
)

CREATE TABLE Song(
	SongId int IDENTITY(1,1) PRIMARY KEY,
	SongName NVARCHAR(60) NOT NULL,
	ArtistId INT,
	CONSTRAINT Song_FK FOREIGN KEY (ArtistId) REFERENCES Artist(ArtistId),
	Duration Float,
	AgeLimitation bit DEFAULT 0,
	FilePath NVARCHAR(200)
)

CREATE TABLE Album(
	AlbumId INT IDENTITY(1,1) PRIMARY KEY,
	AlbumName NVARCHAR(60),
	[Description] NVARCHAR(120),
	ArtistId INT
	CONSTRAINT Album_FK FOREIGN KEY (ArtistId) REFERENCES Artist(ArtistId)
	ON UPDATE CASCADE
	ON DELETE CASCADE
)

CREATE TABLE Playlist_Song_int(
	PlaylistId INT,
	SongId INT,

	CONSTRAINT Playlist_Song_int_FK1 FOREIGN KEY (PlaylistId) REFERENCES Playlist(PlaylistId)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	CONSTRAINT Playlist_Song_int_FK2 FOREIGN KEY (SongId) REFERENCES Song(SongId)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	CONSTRAINT PK_constraint_Playlist_Song_int PRIMARY KEY (PlaylistId, SongId)
)


CREATE TABLE Album_Song_int(
	AlbumId INT,
	SongId INT,

	CONSTRAINT Album_Song_int_FK1 FOREIGN KEY (AlbumId) REFERENCES Album(AlbumId)
		ON UPDATE CASCADE
		ON DELETE CASCADE,
	CONSTRAINT Album_Song_int_FK2 FOREIGN KEY (SongId) REFERENCES Song(SongId)
		ON UPDATE CASCADE
		ON DELETE CASCADE,

	CONSTRAINT PK_constraint_Album_Song_int PRIMARY KEY (AlbumId, SongId)
)


CREATE INDEX User_DOBIndex ON [User](DateOfBirth) INCLUDE(UserName)
GO

CREATE INDEX User_NameIndex ON [User](UserName) INCLUDE(DateOfBirth)
GO


GO
CREATE TRIGGER onInsertToSongTrigger ON Song
AFTER INSERT, UPDATE AS
BEGIN
    IF EXISTS(SELECT * FROM inserted as S1
			  INNER JOIN Song S2 on ((S1.ArtistId = S2.ArtistId)
									AND (S1.SongId != S2.SongId)
									AND  (S1.SongName = S2.SongName)))
		BEGIN
			RAISERROR(N'INVALID NAMES IN BATCH', -1, 1);
			ROLLBACK TRANSACTION
		END
    RETURN
END
GO


CREATE VIEW SongsWithArtistView AS
	SELECT SongID, SongName, Song.ArtistId, ArtistName FROM Song 
		LEFT JOIN Artist ON Song.ArtistId = Artist.ArtistId
GO

CREATE FUNCTION UserPlaylists(@UserId INT) RETURNS TABLE AS
	RETURN
	SELECT * FROM Playlist
		WHERE UserId = @UserId
GO

CREATE FUNCTION GetArtistId(@Name VARCHAR(100)) RETURNS INT AS
BEGIN
	DECLARE @res INT
	SELECT @res = ArtistId FROM Artist
		WHERE ArtistName = @Name
	RETURN @res
END
GO

SELECT * FROM SongsWithArtistView


INSERT INTO [User](PasswordHash) VALUES 
('pass'),
('pass'),
('pass'),
('pass')


INSERT INTO Artist(ArtistName) VALUES 
('Pink Phloyd'),
('Queen'),
('Deep Purple')


INSERT INTO Playlist(PlaylistName, UserId) VALUES 
('Playlist 1', 1),
('Playlist 2', 1)


INSERT INTO Song(SongName, ArtistId) VALUES
('Another brick in the wall', 1),
('High Hopes', 1),

('Bohemian Rhapsody', 2),
('Show must go on', 2),

('Hush', dbo.GetArtistId('Deep Purple'))

INSERT INTO Song(SongName, ArtistId, Duration) VALUES
('Money', 1, 3),
('Time', 1, 4)


SELECT * FROM Playlist
	WHERE UserId = 1

SELECT * FROM Artist
	WHERE Artist.ArtistId = dbo.GetArtistId('Deep Purple')


INSERT INTO Playlist_Song_int(PlaylistId, SongId) VALUES 
(1, 1), 
(1, 2),
(2, 1),
(2, 2),
(2, 3),
(2, 6),
(2, 7)


--USER SONGS--
GO
CREATE FUNCTION GetUserSongs(@id INT) RETURNS TABLE AS
RETURN
	SELECT DISTINCT Playlist_Song_int.SongId AS SongId FROM UserPlaylists(@id) as P 
		INNER JOIN Playlist_Song_int 
		ON Playlist_Song_int.PlaylistId = P.PlaylistId
GO

SELECT * FROM GetUserSongs(1)


SELECT Artist.ArtistId,
	   Count(*) AS SongsCount,
	   AVG(Duration) AS AvgDuration,
	   MIN(Duration) as MinDuration,
	   MAX(Duration) as MaxDuration 
	FROM Artist 
	LEFT JOIN Song On Artist.ArtistId = Song.ArtistId
	GROUP BY Artist.ArtistId
	ORDER BY SongsCount ASC


--User's playlists duration
SELECT p.PlaylistId,  p.PlaylistName, SUM(Duration) as PlaylistDuration FROM UserPlaylists(1) AS P
	JOIN Playlist_Song_int ON P.PlaylistId = Playlist_Song_int.PlaylistId
	JOIN Song ON Playlist_Song_int.SongId = Song.SongId
	GROUP BY p.PlaylistId, p.PlaylistName


DROP VIEW IF EXISTS PlaylistsDuration
GO
CREATE VIEW PlaylistsDuration AS
	SELECT Playlist.PlaylistId AS Id, Playlist.PlaylistName AS [Name], SUM(Duration) AS Duration FROM Playlist_Song_int
		JOIN Playlist ON Playlist_Song_int.PlaylistId = Playlist.PlaylistId
		JOIN Song ON Playlist_Song_int.SongId = Song.SongId
	
		GROUP BY Playlist.PlaylistId, Playlist.PlaylistName
		
GO


DROP VIEW IF EXISTS AlbumsDuration
GO
CREATE VIEW AlbumsDuration AS
	SELECT Album.AlbumId AS Id, Album.AlbumName AS Name, SUM(Duration) AS Duration FROM Album_Song_int
		JOIN Album ON Album_Song_int.AlbumId = Album.AlbumId
		JOIN Song ON Album_Song_int.SongId = Song.SongId
	
		GROUP BY Album.AlbumId, Album.AlbumName
		HAVING SUM(Duration) IS NOT NULL
GO

SELECT * FROM PlaylistsDuration

SELECT * FROM AlbumsDuration

SELECT * FROM PlaylistsDuration
	UNION ALL
SELECT * FROM AlbumsDuration


--CREATE VIEW UserAgeView AS

SELECT ageGroup, COUNT(*) as cnt FROM 
(SELECT a.UserId, 
		CASE
			WHEN Age BETWEEN 0 AND 18 then '0-18'
			WHEN Age BETWEEN 19 AND 25 then '19-25'
			WHEN Age BETWEEN 26 AND 40 then '25-40'
			WHEN Age > 40 then '40+'
			END AS ageGroup
		FROM (SELECT *, DATEDIFF(YEAR, [User].DateOfBirth, GETDATE()) as Age FROM [User]) AS a) AS a
		GROUP BY ageGroup


--Песни не добавленные ни в один плейлист
DROP VIEW IF EXISTS SongsNotInAnyPlaylist
GO
CREATE VIEW SongsNotInAnyPlaylist AS
	SELECT S.SongId, SongName FROM Playlist_Song_int AS I
		RIGHT JOIN Song S ON I.SongId = S.SongId
		WHERE I.SongId IS NULL
GO

--Песни не добавленные ни в один альбом
DROP VIEW IF EXISTS SongsNotInAnyAlbum
GO
CREATE VIEW SongsNotInAnyAlbum AS
	SELECT S.SongId, S.SongName FROM Album_Song_int AS A
		RIGHT JOIN Song S ON A.SongId = S.SongId
		WHERE A.SongId IS NULL
GO

SELECT * FROM SongsNotInAnyPlaylist
UNION
SELECT * FROM SongsNotInAnyAlbum

SELECT * FROM SongsNotInAnyPlaylist
INTERSECT
SELECT * FROM SongsNotInAnyAlbum

SELECT * FROM SongsNotInAnyAlbum
EXCEPT 
SELECT * FROM SongsNotInAnyPlaylist

SELECT * FROM [User]

SELECT * FROM [User] as U1 
	FULL OUTER JOIN [User] U2 ON U1.UserName = U2.UserName
	WHERE U1.UserName is NULL OR U2.UserName is NULL


SELECT * FROM [User]
	WHERE [User].UserName LIKE 'UserName_%'
UNION ALL
SELECT * FROM [User]
	WHERE [User].PasswordHash IN ('compromated_hash_1')


---

SELECT p.PlaylistId,  p.PlaylistName, SUM(Duration) as PlaylistDuration FROM UserPlaylists(1) AS P
	JOIN Playlist_Song_int ON P.PlaylistId = Playlist_Song_int.PlaylistId
	JOIN Song ON Playlist_Song_int.SongId = Song.SongId
	GROUP BY p.PlaylistId, p.PlaylistName
	HAVING SUM(Duration) IS NOT NULL


SELECT p.PlaylistId,  p.PlaylistName, SUM(Duration) as PlaylistDuration FROM UserPlaylists(1) AS P
	JOIN Playlist_Song_int ON P.PlaylistId = Playlist_Song_int.PlaylistId
	JOIN Song ON Playlist_Song_int.SongId = Song.SongId
	GROUP BY p.PlaylistId, p.PlaylistName
	

SELECT p.PlaylistId,  p.PlaylistName, SUM(Duration) as PlaylistDuration FROM UserPlaylists(1) AS P
	JOIN Playlist_Song_int ON P.PlaylistId = Playlist_Song_int.PlaylistId
	JOIN Song ON Playlist_Song_int.SongId = Song.SongId
	GROUP BY p.PlaylistId, p.PlaylistName
	HAVING Count(*) > 1
	
