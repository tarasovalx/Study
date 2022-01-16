USE LAB6
GO

DROP PROCEDURE IF EXISTS selection
GO
CREATE PROCEDURE dbo.selection @cursor CURSOR VARYING OUTPUT AS
    SET @cursor = CURSOR
    FORWARD_ONLY STATIC FOR
    SELECT ArtistName, Country FROM Artist
	OPEN @cursor
GO

DECLARE @artists_cursor CURSOR

EXECUTE dbo.selection @cursor = @artists_cursor OUTPUT

FETCH NEXT FROM @artists_cursor
WHILE (@@FETCH_STATUS = 0)
BEGIN
	FETCH NEXT FROM @artists_cursor
END

CLOSE @artists_cursor
DEALLOCATE @artists_cursor

GO

-----
--2--
-----
DROP FUNCTION IF EXISTS getRandomInRange
DROP VIEW IF EXISTS randomView
GO

CREATE VIEW randomView AS
	SELECT RAND() randomView
GO


CREATE FUNCTION getRandomInRange(@a int, @b int)
	RETURNS INT
	AS
		BEGIN
			DECLARE @rnd FLOAT

            SELECT @rnd = randomView FROM dbo.randomView
            RETURN ROUND(@rnd* (@b - @a), 0) + @a
		END
GO

DROP PROCEDURE IF EXISTS dbo.selectionWithRandom 
GO

CREATE PROCEDURE dbo.selectionWithRandom @cursor CURSOR VARYING OUTPUT AS
    SET @cursor = CURSOR FORWARD_ONLY STATIC FOR
    SELECT ArtistName, Country, dbo.getRandomInRange(1,10) AS Random
	FROM dbo.Artist
    OPEN @cursor;
GO

DECLARE @s_cursor CURSOR
EXECUTE dbo.selectionWithRandom @cursor = @s_cursor OUTPUT

FETCH NEXT FROM @s_cursor
WHILE (@@FETCH_STATUS = 0)
	BEGIN;
		FETCH NEXT FROM @s_cursor;
	END;
CLOSE @s_cursor;
DEALLOCATE @s_cursor;


-----
--3--
-----
DROP FUNCTION IF EXISTS dbo.checkForUk
GO
CREATE FUNCTION dbo.checkForUk(@country NVARCHAR(20))
    RETURNS INT
    AS
        BEGIN
            DECLARE @res INT
            IF @country = N'UK' 
				SET @country = 1
			ELSE 
				SET @country = 0
            RETURN (@country)
        END

GO

DROP PROCEDURE IF EXISTS dbo.updatedProcedure
GO


CREATE PROCEDURE dbo.updatedProcedure AS
    DECLARE @cursor CURSOR
    DECLARE @name NVARCHAR(25)
    DECLARE @country NVARCHAR(20)

    EXECUTE dbo.selection @cursor = @cursor OUTPUT

    FETCH NEXT FROM @cursor INTO @name, @country
    WHILE (@@FETCH_STATUS=0)
    BEGIN
        IF (dbo.checkForUk(@country) > 0)
            PRINT @name + ' UK'
        ELSE
            print @name + ' NOT'

        FETCH NEXT FROM @cursor INTO @name, @country
    END

    CLOSE @cursor
    DEALLOCATE @cursor

GO
EXECUTE dbo.updatedProcedure
GO

--4--
DROP FUNCTION IF EXISTS dbo.tableFunction
GO

CREATE FUNCTION dbo.tableFunction() 
RETURNS @outTable TABLE (
    ArtistName NVARCHAR(20) NOT NULL,
    Country NVARCHAR(20) NOT NULL,
    rnd INT NOT NULL
)
AS
    BEGIN
        INSERT @outTable 
        SELECT ArtistName, Country, dbo.getRandomInRange(0,10) as Score 
        FROM Artist 
		WHERE dbo.checkForUk(Country) > 0 
        RETURN 
    END
GO
 
DROP FUNCTION IF EXISTS dbo.tableFunction2
GO

CREATE FUNCTION dbo.tableFunction2() RETURNS TABLE 
	AS RETURN SELECT ArtistName, Country, dbo.getRandomInRange(0,10) as Score 
			  FROM Artist 
			  WHERE dbo.checkForUk(Country) > 0 
GO

ALTER PROCEDURE dbo.selectionWithRandom @cursor CURSOR VARYING OUTPUT
AS
    SET @cursor = CURSOR 
	FORWARD_ONLY STATIC FOR 
	SELECT * FROM dbo.tableFunction()
	OPEN @cursor
GO

DECLARE @table_cursor CURSOR
EXECUTE dbo.selectionWithRandom @cursor = @table_cursor OUTPUT

DECLARE @name NVARCHAR(20)
DECLARE @country NVARCHAR(20)
DECLARE @rnd INT

FETCH NEXT FROM @table_cursor INTO @name, @country, @rnd
WHILE (@@FETCH_STATUS = 0)
	BEGIN
		FETCH NEXT FROM @table_cursor INTO @name, @country, @rnd
        PRINT @name + ' '+ @country + ' ' + CAST(@rnd as NVARCHAR(2))
	END

CLOSE @table_cursor
DEALLOCATE @table_cursor
GO
