namespace AOC_2020.Days
{

    public class Day3 : SeparatedDay
    {
        const char threeSign = '#';

        public override object part1()
        {
            return CountTrees(3, 1);
        }

        public override object part2()
        {
            long treesA = CountTrees(1,1);
            long treesB = CountTrees(3,1);
            long treesC = CountTrees(5,1);
            long treesD = CountTrees(7,1);
            long treesE = CountTrees(1,2);
            
            return treesA * treesB * treesC * treesD * treesE;
        }

        private int CountTrees(int stepRight, int stepDown)
        {
            int treesPassed = default;
            int currentX = default;
            int currentY = default;
            var lineLength = Lines[0].Length;
            var linesCount = Lines.Count;
                
            while (currentY < linesCount)
            {
                if (Lines[currentY][currentX] == threeSign)
                {
                    treesPassed++;
                }

                currentX += stepRight;
                currentY += stepDown;
                if (currentX >= lineLength)
                {
                    currentX -= lineLength;
                }
                
            }

            return treesPassed;
            
        }

        private int CountTreesPartTwo()
        {
            (int right, int down)[] slopes = new (int, int)[] {(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)};

            return 0;
        }
    }

}