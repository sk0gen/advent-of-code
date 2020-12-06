namespace AOC_2020.Days
{

    public abstract class SeparatedDay : Day
    {
        /**
     * @return The result of part 1
     */
        public abstract object part1();

        /**
     * @return The result of part 2
     */
        public abstract object part2();
    
        protected override Result Evaluate()
        {
            return new Result(part1(), part2());
        }
    }

}