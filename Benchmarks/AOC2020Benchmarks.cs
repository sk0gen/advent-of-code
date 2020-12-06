using System.Reflection;
using AOC_2020.Days;
using BenchmarkDotNet.Attributes;

namespace Benchmarks
{
    [MemoryDiagnoser]
    public class AOC2020Benchmarks
    {
        
        private readonly Day1 _day1 = new Day1();
        private readonly Day2 _day2 = new Day2();
        private readonly Day3 _day3 = new Day3();
        private readonly Day4 _day4 = new Day4();
        private readonly Day5 _day5 = new Day5();
        private readonly Day6 _day6 = new Day6();

        public AOC2020Benchmarks()
        {
            _day1.Run();
            _day2.Run();
            _day3.Run();
            _day4.Run();
            _day5.Run();
            _day6.Run();
        }

        [Benchmark]
        public object Day1Part1() => _day1.part1();
        [Benchmark]
        public object Day1Part2() => _day1.part1();
        [Benchmark]
        public object Day2Part1() => _day2.part1();
        [Benchmark]
        public object Day2Part2() => _day2.part1();
        [Benchmark]
        public object Day3Part1() => _day3.part1();
        [Benchmark]
        public object Day3Part2() => _day3.part1();
        [Benchmark]
        public object Day4Part1() => _day4.part1();
        [Benchmark]
        public object Day4Part2() => _day4.part1();
        [Benchmark]
        public object Day5Part1() => _day5.part1();
        [Benchmark]
        public object Day5Part2() => _day5.part1();
        [Benchmark]
        public object Day6Part1() => _day6.part1();
        [Benchmark]
        public object Day6Part2() => _day6.part1();
    }

}