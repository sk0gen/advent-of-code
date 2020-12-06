using System.Collections.Generic;
using System.Linq;

namespace AOC_2020.Days
{

    public class Day6 : SeparatedDay
    {
        private string[] GroupsData;

        public override object part1()
        {
            return CountGroupsYesAnswers();
        }

        public override object part2()
        {
            return CountGroupsWhereEveryoneAnswersYesToSameQuestions();
        }

        private int CountGroupsYesAnswers()
        {
            int count = default;
            for (var i = 0; i < GroupsData.Length; ++i)
            {
                count += CountYesInGroup(GroupsData[i]);
            }

            return count;

            int CountYesInGroup(string groupData)
            {
                var yesAnswers = new HashSet<char>(26);
                foreach (var yesAnswer in groupData)
                {
                    yesAnswers.Add(yesAnswer);
                }

                yesAnswers.Remove('\n');
                return yesAnswers.Count;
            }
        }

        private int CountGroupsWhereEveryoneAnswersYesToSameQuestions()
        {
            var groupSize = 0;
            var total = 0;
            var groupAnswers = new int[26];
            for (int i = 0; i < Lines.Count; ++i)
            {
                if (Lines[i] != "")
                {
                    foreach (var answer in Lines[i])
                    {
                        groupAnswers[answer - 97] += 1;
                    }
                    groupSize++;
                }
                else
                {
                    for (var index = 0; index < groupAnswers.Length; index++)
                    {
                        var answer = groupAnswers[index];
                        if (answer == groupSize)
                        {
                            total++;
                        }
                    }

                    groupAnswers = new int[26];
                    groupSize = 0;
                }
            }

            return total;
        }

        protected override void Parse()
        {
            var fileString = string.Join("\n", Lines.ToArray());
            GroupsData = fileString.Split("\n\n");
        }
    }

}