 USE LAB10
 
 SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED
 BEGIN TRANSACTION
	 SELECT * FROM TestTable
	 SELECT resource_type, resource_subtype, request_mode FROM sys.dm_tran_locks
 COMMIT TRANSACTION
 GO

SET TRANSACTION ISOLATION LEVEL READ COMMITTED
--SET TRANSACTION ISOLATION LEVEL REPEATABLE READ
--SET TRANSACTION ISOLATION LEVEL SERIALIZABLE
BEGIN TRANSACTION
    SELECT * FROM TestTable 
    WAITFOR DELAY '00:00:05'  
    SELECT * FROM TestTable
    SELECT resource_type, resource_subtype, request_mode FROM sys.dm_tran_locks
COMMIT TRANSACTION
GO