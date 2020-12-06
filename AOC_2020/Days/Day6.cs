using System;
using System.Collections.Generic;

namespace AOC_2020.Days
{

    public class Day6 : SeparatedDay
    {
        private string[] GroupFormData;

        protected override object part1()
        {
            return CountGroupsYesAnswers();
        }

        protected override object part2()
        {
            return null;
        }

        private int CountGroupsYesAnswers()
        {
            int count = default;
            for (var i = 0; i < GroupFormData.Length; ++i)
            {
                count += CountYesInGroup(GroupFormData[i]);
            }


            return count;
            int CountYesInGroup(string groupData)
            {
                var yesAnswers = new HashSet<char>(64);
                foreach (var yesAnswer in groupData)
                {
                    yesAnswers.Add(yesAnswer);
                }
                
                yesAnswers.Remove('\n');
                return yesAnswers.Count;
            }
            
        }

        protected override void Parse()
        {
            var fileString = string.Join("\n", Lines.ToArray());
            GroupFormData = fileString.Split("\n\n");
        }
    }

}