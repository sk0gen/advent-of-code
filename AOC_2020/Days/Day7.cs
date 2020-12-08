using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

namespace AOC_2020.Days
{

    public class Day7 : SeparatedDay
    {
        public override object part1()
        {
            return SolvePart1();
        }

        public override object part2()
        {
            return SolvePart2();
        }
        
        
        public int SolvePart1()
        {
            var bagMap = new Dictionary<string, HashSet<string>>();
            foreach (var line in Lines)
            {
                var split = line.Split(" bags contain ");
                var color = split[0];
                foreach (var match in Regex.Matches(split[1], @"(?<=\d ).*?(?= bag)"))
                {
                    if (bagMap.TryGetValue(match.ToString(), out var hs))
                    {
                        hs.Add(color);
                    }
                    else
                    {
                        bagMap[match.ToString()] = new HashSet<string>(new[] { color });
                    }
                }
            }

            var bags = new HashSet<string>();
            var queue = new Queue<string>();
            queue.Enqueue("shiny gold");
            while (queue.Any())
            {
                if (bagMap.TryGetValue(queue.Dequeue(), out var hs))
                {
                    foreach (var item in hs)
                    {
                        bags.Add(item);
                        queue.Enqueue(item);
                    }
                }
            }

            return bags.Count;
        }

        public int SolvePart2()
        {
            var bagMap = new Dictionary<string, List<(int Count, string Color)>>();

            foreach (var line in Lines)
            {
                var split = line.Split(" bags contain ");
                var color = split[0];
                var bags =
                    Regex.Matches(split[1], @"(\d) (.*?)(?= bag)")
                    .OfType<Match>()
                    .Select(m => (int.Parse(m.Groups[1].Value), m.Groups[2].Value));
                bagMap[color] = bags.ToList();
            }

            var total = 0;
            var q = new Queue<(int, string)>();
            q.Enqueue((1, "shiny gold"));
            while (q.Any())
            {
                var (count, color) = q.Dequeue();
                var hs = bagMap[color];                
                total += count;

                foreach (var (Count, Color) in hs)
                {
                    q.Enqueue((count * Count, Color));
                }
            }

            return total - 1;
        }
        
    }

}