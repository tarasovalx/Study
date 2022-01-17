using Microsoft.Data.SqlClient;
using System;
using System.Collections.Generic;
using System.Data;
using System.Data.Common;

namespace DB_lab_12
{
    public class DBDisconnected : IDisposable
    {
        private DataSet _dataSet;
        private SqlDataAdapter _dataAdapter;
        private SqlConnection connection;

        private static readonly string tableName = "Artist";


        public DBDisconnected() { }

        public void ConfigureDisconnectedLayer()
        {
            connection = SqlConnectionExtension.Connect();
            _dataSet = new DataSet();

            try
            {
                _dataAdapter = new SqlDataAdapter($"SELECT * FROM {tableName}", connection);
                _dataAdapter.Fill(_dataSet, tableName);
            } 
            catch (Exception e)
            {
                Console.WriteLine(e);
                return;
            }

            var insertCommand = new SqlCommand();
            insertCommand.Connection = connection;
            insertCommand.CommandText = $"INSERT INTO {tableName}(ArtistName) VALUES (@ArtistName)"
                                      + $"\nSET @ArtistId = SCOPE_IDENTITY()";

            var p = new SqlParameter("@ArtistId", SqlDbType.Int, 1, "ArtistId");
            p.Direction = ParameterDirection.Output;

            insertCommand.Parameters.Add(p);
            insertCommand.Parameters.Add("@ArtistName", SqlDbType.NVarChar, 60, "ArtistName");

            var updateCommand = new SqlCommand();
            updateCommand.Connection = connection;
            updateCommand.CommandText = $"UPDATE {tableName} SET ArtistName = @ArtistName " +
                                        $"WHERE ArtistId = @oldArtistId";

            var deleteCommand = new SqlCommand();
            var parameter = deleteCommand.Parameters.Add("@ArtistId", SqlDbType.Int, 1, "ArtistId");
            parameter.SourceVersion = DataRowVersion.Original;
            deleteCommand.CommandText = $"DELETE FROM {tableName} WHERE ArtistId = @ArtistId";



            updateCommand.Parameters.Add("@ArtistId", SqlDbType.Int, 1, "ArtistId");
            updateCommand.Parameters.Add("@ArtistName", SqlDbType.NVarChar, 60, "ArtistName");

            parameter = updateCommand.Parameters.Add("@oldArtistId", SqlDbType.Int, 1, "ArtistId");
            parameter.SourceVersion = DataRowVersion.Original;


            _dataAdapter.DeleteCommand = deleteCommand;
            _dataAdapter.InsertCommand = insertCommand;
            _dataAdapter.UpdateCommand = updateCommand;
        }

        public void DeleteByCondition(Func<DataRow, bool> deleteFilter)
        {
            try
            {
                foreach (DataRow row in _dataSet.Tables[tableName].Rows)
                {
                    if (deleteFilter(row))
                        row.Delete();
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }

        public void InsertIntoTable(Func<DataRow, DataRow> rowCreator)
        {
            try
            {
                DataRow row = _dataSet.Tables[tableName].NewRow();
                row = rowCreator(row);

                _dataSet.Tables[tableName].Rows.Add(row);
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }

        public void InsertIntoTable(Func<Func<DataRow>, IEnumerable<DataRow>> rowCreator)
        {
            Func<DataRow> getNewRow = _dataSet.Tables[tableName].NewRow;

            try
            {
                foreach (var row in rowCreator(getNewRow))
                    _dataSet.Tables[tableName].Rows.Add(row);
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }

        public void ShowTableContents()
        {
            try
            {
                using var reader = _dataSet.CreateDataReader();
                foreach (var t in reader.GetColumnSchema())
                    Console.Write($"{t.ColumnName} ");

                Console.WriteLine();
                while (reader.Read())
                {
                    for (int i = 0; i < reader.FieldCount; i++)
                        Console.Write($"{reader.GetValue(i)} ");

                    Console.Write("\n");
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }

        public void UpdateByCondition(Action<DataRow> setter, Func<DataRow, bool> filter)
        {
            try
            {
                foreach (DataRow row in _dataSet.Tables[tableName].Rows)
                {
                    if (filter(row))
                        setter(row);
                }
            }
            catch (Exception e)
            {
                Console.Write(e);
            }
        }

        public void UpdateOnServer()
        {
            try
            {
                _dataAdapter.Update(_dataSet, tableName);
            }
            catch (Exception e)
            {
                Console.Write(e);
                Environment.Exit(-1);
            }
        }

        public void Dispose()
        {
            if (connection.State == ConnectionState.Open)
                connection.Close();
        }
    }
}