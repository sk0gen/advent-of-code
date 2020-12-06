using System;
using System.Collections.Generic;
using System.Linq;

namespace AOC_2020.Days
{

    public class Day1 : SeparatedDay
    {
        private const int yearTarget = 2020;
        private HashSet<int> Expenses;

        protected override object part1()
        {
            return FindSumWithTwoNumberEqualsTarget(yearTarget);
        }

        protected override object part2()
        {
            return FindSumWithThreeNumbersEqualsTarget(yearTarget);
        }

        protected override void Parse()
        {
            Expenses = Lines.Select(int.Parse).ToHashSet();
        }

        private int? FindSumWithTwoNumberEqualsTarget(int target)
        {
            foreach (var expense in Expenses)
            {
                if (Expenses.Contains(target - expense))
                {
                    return (target - expense) * expense;
                }
            }

            return null;
        }

        private int? FindSumWithThreeNumbersEqualsTarget(int target)
        {
            foreach (var expense in Expenses)
            {
                var newTarget = target - expense;
                var result = FindSumWithTwoNumberEqualsTarget(newTarget);
                if (result != null)
                {
                    return expense * result;
                }
            }
            return null;
        }
    }

}