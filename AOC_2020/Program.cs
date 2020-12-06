using System;
using AOC_2020.Days;

namespace AOC_2020
{

    class Program
    {
        static void Main(string[] args)
        {
            var input = "";
            do
            {
                Console.WriteLine("Advent of code 2020!!");
                Console.WriteLine("1. Day_01");
                Console.WriteLine("2. Day_02");
                Console.WriteLine("3. Day_03");
                Console.WriteLine("4. Day_04");
                Console.WriteLine("5. Day_05");
                Console.WriteLine("0. Clear Console");
                Console.WriteLine("-1. Exit");
                input = Console.ReadLine();
                switch (input)
                {
                    case "1":
                    {
                        var day01 = new Day1();
                        day01.Run();
                        break;
                    }
                    case "2":
                    {
                        var day02 = new Day2();
                        day02.Run();
                        break;
                    }
                    case "3":
                    {
                        var day03 = new Day3();
                        day03.Run();
                        break;
                    }
                    case "4":
                    {
                        var day04 = new Day4();
                        day04.Run();
                        break;
                    }
                    case "5":
                    {
                        var day05 = new Day5();
                        day05.Run();
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