using System;
using System.Data;

namespace DB_lab_12
{

    class Program
    {
        private static readonly DBConnected connectedDB = new DBConnected();
        private static readonly DBDisconnected nonConnectedDB = new DBDisconnected();


        static void Main()
        {
            runConnectedLayer();
            Console.ReadKey();
            runDisconnectedLayer();
        }

        private static void runDisconnectedLayer()
        {
            Console.WriteLine("Disconnected layer");

            nonConnectedDB.ConfigureDisconnectedLayer();
            nonConnectedDB.ShowTableContents();
            nonConnectedDB.InsertIntoTable(newRow =>
            {
                var inserting = new[] {
                    ("Van der graaf generator"),
                    ("Frank Synatra"),
                };
                var res = new DataRow[inserting.Length];
                for (var i = 0; i < inserting.Length; i++)
                {
                    res[i] = newRow();
                    res[i]["ArtistName"] = inserting[i];
                }


                return res;
            }
            );
            nonConnectedDB.ShowTableContents();
            nonConnectedDB.UpdateByCondition(
                row => row["ArtistName"] = "Frank Synatra",
                row => row["ArtistName"].ToString() == "Frank Sinatra");


            nonConnectedDB.ShowTableContents();

            Console.ReadKey();

            nonConnectedDB.UpdateOnServer();
            Console.WriteLine("____UpdatedDataTabe____");
            nonConnectedDB.ShowTableContents();
            nonConnectedDB.Dispose();
        }

        private static void runConnectedLayer()
        {
            connectedDB.ShowTableContents();
            connectedDB.InsertIntoTableWithGetKeys("Frank Sinatra");
            
            connectedDB.ShowTableContents();
            connectedDB.UpdateByCondition(4, "Frank Synatra");
            connectedDB.ShowTableContents();
            connectedDB.DeleteByCondition(4);
            connectedDB.ShowTableContents();

        }
    }
}