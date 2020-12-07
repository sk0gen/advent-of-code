using System;
using System.Collections.Generic;
using System.Diagnostics;

namespace AOC_2020.Days
{

    public abstract class Day
    {
        private Result Result { get; set; }

        protected readonly List<string> Lines = new List<string>();

        protected Day()
        {
            Lines = DataManager.Read(int.Parse(GetType().Name.Substring(3)));
        }

        /**
     * Execute a given day; outputting part 1, part 2, and the time taken.
     */
        public void Run()
        {
            var stopWatch = new Stopwatch();
            stopWatch.Start();
            Result = ParseAndEvaluate();
            stopWatch.Stop();
            // Get the elapsed time as a TimeSpan value.
            var ts = stopWatch.Elapsed;
            var elapsedTime = $"{ts.Hours:00}:{ts.Minutes:00}:{ts.Seconds:00}.{ts.Milliseconds / 10:00}";
            Console.WriteLine(GetType().Name);
            Console.WriteLine("✅ Part 1: " + Result.Part1);
            Console.WriteLine("✅ Part 2: " + Result.Part2);
            Console.WriteLine($"Completed in {elapsedTime}\n");
        }

        /**
     * Parse and then evaluate a day's code.
     * This is not guaranteed to be repeatable without constructing a new instance of the class.
     *
     * @return A {@link Result} holding data of the first and second part
     */
        public Result ParseAndEvaluate()
        {
            Parse();
            Result = Evaluate();
            return Result;
        }

        /**
     * This internal method is what actually evaluates the result of part 1 and part 2.
     */
        protected abstract Result Evaluate();

        /**
         * This internal method can be overridden to parse the {@link #lines} of the day into something more useful for
         * the challenge.
         * <p>
         *     This method will automatically be run before {@link #evaluate()}.
         */
        protected virtual void Parse()
        {
        }
    }

}