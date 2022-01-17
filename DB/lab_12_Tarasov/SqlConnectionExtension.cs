using Microsoft.Data.SqlClient;
using System;
using System.Configuration;

namespace DB_lab_12
{
    public static class SqlConnectionExtension
    {
        public static void Disconnect(this SqlConnection connection)
        {
            try
            {
                connection.Close();
                Console.WriteLine("Connection closed");
            }
            catch (Exception e)
            {
                Console.Write(e);
                Environment.Exit(-1);
            }
        }

        public static SqlConnection Connect()
        {
            var connection = new SqlConnection(ConfigurationManager.AppSettings["connectionString"]);
            try
            {
                connection.Open();
                Console.WriteLine("Connection openned");
            }
            catch (Exception e)
            {
                Console.Write(e);
                Environment.Exit(-1);
            }
            return connection;
        }
    }
}
