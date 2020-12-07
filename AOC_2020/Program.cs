using System;
using AOC_2020.Days;

namespace AOC_2020
{

    internal class Program
    {
        private static void Main(string[] args)
        {
            var input = "";
            do
            {
                Console.WriteLine("Advent of code 2020!!");
                Console.WriteLine("1. Run AOC");
                Console.WriteLine("0. Clear Console");
                Console.WriteLine("-1. Exit");
                input = Console.ReadLine();
                switch (input)
                {
                    case "1":
                    {
                        var days = ReflectionHelper.GetInheritedClasses(typeof(SeparatedDay));
                        foreach (var day in days)
                        {
                            var type = (SeparatedDay) Activator.CreateInstance(day);
                            type.Run();
                        }

                        break;
                    }
                    
                    case "0":
                    {
                        Console.Clear();
                        break;
                    }
                }
            } while (input != "-1");
        }
    }

}