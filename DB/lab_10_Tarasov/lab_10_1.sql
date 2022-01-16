USE LAB10

--Dirty read--
BEGIN TRANSACTION
	UPDATE TestTable SET ColInt = ColInt + 100 WHERE id = 1
	WAITFOR DELAY '00:00:05'
	SELECT resource_type, resource_subtype, request_mode FROM sys.dm_tran_locks
ROLLBACK TRANSACTION


--Unrepeatable read--
BEGIN TRANSACTION
	UPDATE TestTable SET ColInt = ColInt + 100 WHERE id = 1
	SELECT resource_type, resource_subtype, request_mode FROM sys.dm_tran_locks
COMMIT TRANSACTION


--Phantom read--
BEGIN TRANSACTION
    INSERT INTO TestTable(ColInt, ColChar)
	VALUES(10, 'Name_C')
    SELECT resource_type, resource_subtype, request_mode FROM sys.dm_tran_locks
COMMIT TRANSACTION
GO