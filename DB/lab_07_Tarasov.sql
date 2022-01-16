USE LAB6;

DROP VIEW IF EXISTS UserView;
GO

CREATE VIEW UserView AS 
	SELECT * FROM [User] WHERE (DATEDIFF(YEAR, [User].DateOfBirth, GETDATE()) >= 18)
GO

SELECT * FROM UserView
GO


DROP VIEW IF EXISTS SongsByCountryView;
GO

CREATE VIEW SongsByCountryView AS
	SELECT Country, Count(*) AS [Count] FROM Song JOIN Artist
	ON Song.AritstId = Artist.ArtistId
	GROUP BY Country
GO

SELECT * FROM SongsByCountryView


DROP INDEX IF EXISTS [User].UserIndex 
CREATE INDEX UserIndex ON [User](DateOfBirth) INCLUDE(UserName)
GO


SElECT UserName FROM [User] 
WHERE DateOfBirth < '2002-01-01'
GO


DROP INDEX IF EXISTS [User].UserIndex_1 
CREATE INDEX UserIndex_1 ON [User](UserName) INCLUDE(DateOfBirth)
GO

SELECT UserName FROM [User] WITH (INDEX(UserIndex_1))
WHERE UserName LIKE 'A%'


SET NUMERIC_ROUNDABORT OFF;
SET ANSI_PADDING, ANSI_WARNINGS, CONCAT_NULL_YIELDS_NULL, ARITHABORT,
   QUOTED_IDENTIFIER, ANSI_NULLS ON;


DROP VIEW IF EXISTS UserIndexedView;
GO
CREATE VIEW UserIndexedView WITH SCHEMABINDING AS
	SELECT UserId, UserName, DateOfBirth FROM dbo.[User]
		WHERE DateOfBirth < CONVERT(datetime, '2002.01.01', 102)
GO

DROP INDEX IF EXISTS UserIndexedView.View_Index
CREATE UNIQUE CLUSTERED INDEX View_Index ON UserIndexedView(UserId, UserName)

DROP INDEX IF EXISTS UserIndexedView.View_Index1
CREATE UNIQUE NONCLUSTERED INDEX View_Index1 ON UserIndexedView(UserId, UserName)

SElECT * FROM UserIndexedView WITH(INDEX(View_Index))
	WHERE UserName Like('A%')
GO

SELECT [Count] FROM SongsByCountryView WITH(INDEX(Ind))
WHERE Country = 'UK'