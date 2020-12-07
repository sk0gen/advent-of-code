using System.Linq;
using System.Text;

namespace AOC_2020.Days
{

    public class Day7 : SeparatedDay
    {
        public override object part1()
        {
            return null;
        }

        public override object part2()
        {
            return null;
        }

        private int SolvePart1()
        {
            var bags = Lines.Where(x => !x.Contains("no")).ToList();

            return 0;
        }
        protected override void Parse()
        {
            for (var index = 0; index < Lines.Count; index++)
            {
                var line = Lines[index];
                var sb = new StringBuilder(line);
                sb.Replace(" bags.", "").Replace("bag ", "").Replace("bags","").Replace("bag","").Replace("  "," ");
                Lines[index] = sb.ToString();
            }
        }
    }

}