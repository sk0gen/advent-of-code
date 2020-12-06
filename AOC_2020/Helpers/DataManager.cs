using System;
using System.Collections.Generic;
using System.IO;

namespace AOC_2020.Days
{

    public static class DataManager
    {
        public static List<string> Read(int day)
        {
            var path = GetPath(day);
            if (path == null)
                return new List<string>();
            var lines = GetDataFromFile(path);

            if (lines.Count != 0)
            {
                return lines;
            }

            return new List<string>();
        }

        private static string GetPath(int day)
        {
            string[] paths = {"Resources", day.ToString()};
            var fullPath = Path.Combine(paths);
            return fullPath;
        }

        private static List<string> GetDataFromFile(string path)
        {
            try
            {
                if (path != null && File.Exists(path))
                {
                    return new List<string>(File.ReadAllLines(path));
                }
            }
            catch (IOException e)
            {
                Console.WriteLine(e.StackTrace);
            }

            return new List<string>();
        }
    }

}