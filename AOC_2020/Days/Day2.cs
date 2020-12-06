using System;
using System.Linq;
using System.Threading.Tasks;

namespace AOC_2020.Days
{

    public class Day2 : SeparatedDay
    {
        private PasswordChecker[] PasswordCheckers;

        protected override object part1()
        {
            return PasswordCheckers.Count(x => x.isValidPartOne);
        }

        protected override object part2()
        {
            return PasswordCheckers.Count(x => x.isValidPartTwo);
        }

        protected override void Parse()
        {
            PasswordCheckers = new PasswordChecker[Lines.Count];
            Parallel.For(0, Lines.Count,
                (i) => { PasswordCheckers[i] = PasswordDeserializer.DeserializeSpan(Lines[i].AsSpan()); });
        }
        
        
    }

    public class PasswordDeserializer
    {
        public static PasswordChecker DeserializeSpan(ReadOnlySpan<char> input)
        {
            var minMaxDividerIndex = input.IndexOf('-');
            var passwordStartIndex = input.LastIndexOf(' ') + 1;
            var requiredCharacterIndex = input.IndexOf(' ') + 1;

            return new PasswordChecker()
            {
                RequiredCharacter = input[requiredCharacterIndex],
                Minimum = int.Parse(input.Slice(0, minMaxDividerIndex)),
                Maximum = int.Parse(
                    input.Slice(minMaxDividerIndex + 1, requiredCharacterIndex - minMaxDividerIndex - 1)),
                Password = input[passwordStartIndex..].ToString()
            };
        }
    }

    public class PasswordChecker
    {
        public char RequiredCharacter { get; set; }
        public int Minimum { get; set; }
        public int Maximum { get; set; }
        public string Password { get; set; }

        public bool isValidPartOne
        {
            get
            {
                int count = default;
                for (var i = 0; i < Password.Length; i++)
                {
                    if (RequiredCharacter == Password[i])
                    {
                        count++;
                    }
                }

                return count <= Maximum && count >= Minimum;
            }
        }

        public bool isValidPartTwo => Password[Minimum-1] == RequiredCharacter ^ Password[Maximum-1] == RequiredCharacter;
    }

}