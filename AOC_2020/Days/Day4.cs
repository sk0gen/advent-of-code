using System;
using System.Text.RegularExpressions;

namespace AOC_2020.Days
{

    public class Day4 : SeparatedDay
    {
        readonly string[] requiredTags = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"};
        readonly string[] patterns =
        {
            @"byr:([0-9]+)",
            @"iyr:([0-9]+)",
            @"eyr:([0-9]+)",
            @"hgt:([0-9]+)(cm|in)",
            @"hcl:#[0-9a-f]{6}",
            @"ecl:(amb|blu|brn|gry|grn|hzl|oth)",
            @"pid:([0-9]){9}\b"
        };

        private readonly (int, int)[] years = new[]
        {
            (1920, 2002),
            (2010, 2020),
            (2020, 2030),
        };
        private string[] passportsData;

        protected override object part1()
        {
            return CountValidPasswords();
        }

        protected override object part2()
        {
            return CountValidPassword2();
        }

        private int CountValidPasswords()
        {
            int count = default;
            foreach (var t in passportsData)
            {
                if (HasRequiredTags(t))
                {
                    count++;
                }
            }

            return count;
        }

        private bool HasRequiredTags(string t)
        {
            var tags = t.Split(requiredTags, StringSplitOptions.RemoveEmptyEntries);
            return tags.Length == 8 || tags.Length == 7 && !t.Contains("cid");
        }

        private int CountValidPassword2()
        {
            int count = default;
            foreach (var passport in passportsData)
            {
                if (PassportValidationPartTwo(passport))
                {
                    count++;
                }
            }

            return count;
        }

        private bool PassportValidationPartTwo(string passport)
        {
            int total = default;
            for (var i = 0; i < patterns.Length; i++)
            {
                var regex = new Regex(patterns[i]);

                var match = regex.Match(passport);

                if (!match.Success)
                {
                    return false;
                }

                switch (i)
                {
                    case > 3:
                        continue;
                    case < 3 when !CheckYear(int.Parse(match.Groups[1].Value), years[i]):
                        return false;
                    case 3 when !CheckHeight(int.Parse(match.Groups[1].Value), match.Groups[2].Value):
                        return false;
                }
            }

            return true;
        }

        static bool CheckYear(int year, (int lower, int upper) range)
        {
            return year >= range.lower && year <= range.upper;
        }

        static bool CheckHeight(int height, string unitsStr)
        {
            return unitsStr switch
            {
                "cm" => height is >= 150 and <= 193,
                _ => height is >= 59 and <= 76
            };
        }

        protected override void Parse()
        {
            var fileString = string.Join("\n", Lines.ToArray());
            passportsData = fileString.Split("\n\n");
        }
    }

}