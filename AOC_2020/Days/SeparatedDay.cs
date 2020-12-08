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
    
        protected override (object, object) Evaluate()
        {
            return (part1(), part2());
        }
    }

}