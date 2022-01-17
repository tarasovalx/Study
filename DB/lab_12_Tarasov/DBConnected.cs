using Microsoft.Data.SqlClient;
using System;
using System.Data;

namespace DB_lab_12
{
    public class DBConnected
    {
        private static readonly string tableName = "Artist";

        private readonly SqlCommand insertCommand;
        private readonly SqlParameter insertNameParam;

        private readonly SqlCommand insertCommandWithGetId;
        private readonly SqlParameter insertWithGetIdNameParam;


        private readonly SqlCommand deleteCommand;
        private readonly SqlParameter deleteIdParam;

        private readonly SqlCommand updateCommand;
        private readonly SqlParameter updateIdParam;
        private readonly SqlParameter updateNameParam;

        public DBConnected()
        {
            insertCommand = new SqlCommand();
            insertNameParam = new SqlParameter("@ArtistName", SqlDbType.NVarChar, 60);
            insertCommand.Parameters.Add(insertNameParam);
            insertCommand.CommandText = $"INSERT INTO {tableName}(ArtistName) VALUES (@ArtistName)";

            insertCommandWithGetId = new SqlCommand();
            insertWithGetIdNameParam = new SqlParameter("@ArtistName", SqlDbType.NVarChar, 60);
            insertCommandWithGetId.Parameters.Add(insertWithGetIdNameParam);
            insertCommandWithGetId.CommandText = $"INSERT INTO {tableName}(ArtistName) VALUES (@ArtistName)"
                                               + $"\nSELECT IDENT_CURRENT('{tableName}')";

            deleteCommand = new SqlCommand();
            deleteIdParam = new SqlParameter("@ArtistId", SqlDbType.Int, 1);
            deleteCommand.Parameters.Add(deleteIdParam);
            deleteCommand.CommandText = $"DELETE FROM {tableName} WHERE ArtistId = @ArtistId";

            updateCommand = new SqlCommand();
            updateIdParam = new SqlParameter("@ArtistId", SqlDbType.Int, 1);
            updateNameParam = new SqlParameter("@ArtistName", SqlDbType.NVarChar, 60);
            updateCommand.Parameters.Add(updateIdParam);
            updateCommand.Parameters.Add(updateNameParam);
            updateCommand.CommandText = $"UPDATE {tableName} SET ArtistName = @ArtistName WHERE ArtistId = @ArtistId";
        }

        public void ShowTableContents()
        {
            using (var connection = SqlConnectionExtension.Connect())
            {
                try
                {
                    SqlCommand command = connection.CreateCommand();
                    command.Connection = connection;

                    command.CommandText = $"SELECT * FROM {tableName}";

                    SqlDataReader reader = command.ExecuteReader();

                    foreach (var t in reader.GetColumnSchema())
                        Console.Write($"{t.ColumnName} ");
                    Console.Write("\n");


                    while (reader.Read())
                    {
                        for (int i = 0; i < reader.FieldCount; i++)
                            Console.Write($"{reader.GetValue(i)} ");
                        Console.Write("\n");
                    }
                    reader.Close();
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }
            }
        }

        public void InsertIntoTable(string ArtistName)
        {
            using (var connection = SqlConnectionExtension.Connect())
            {
                var command = insertCommand;
                command.Connection = connection;
                insertNameParam.Value = ArtistName;

                try
                {
                    command.ExecuteNonQuery();
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }

            }
        }

        public void InsertIntoTableWithGetKeys(string ArtistName)
        {
            using (var connection = SqlConnectionExtension.Connect())
            {
                var command = insertCommandWithGetId;
                insertWithGetIdNameParam.Value = ArtistName;

                command.Connection = connection;

                try
                {
                    decimal id = (decimal)command.ExecuteScalar();
                    Console.WriteLine($"INSERTED NAME: {ArtistName}");
                    Console.WriteLine($"INSERTED ID: {id}");
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }

            }
        }

        public void UpdateByCondition(int ArtistId, string ArtistName)
        {
            using (var connection = SqlConnectionExtension.Connect())
            {
                var command = updateCommand;
                command.Connection = connection;
                updateIdParam.Value = ArtistId;
                updateNameParam.Value = ArtistName;

                try
                {
                    command.ExecuteNonQuery();
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }
            }
        }

        public void DeleteByCondition(int ArtistId)
        {
            using (var connection = SqlConnectionExtension.Connect())
            {
                var command = deleteCommand;
                command.Connection = connection;

                deleteIdParam.Value = ArtistId;
                try
                {
                    command.ExecuteNonQuery();
                }
                catch (SqlException e)
                {
                    Console.WriteLine(e);
                }
            }
        }
    }
}