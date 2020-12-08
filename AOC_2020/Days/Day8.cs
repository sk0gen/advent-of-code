using System;
using System.Collections.Generic;

namespace AOC_2020.Days
{

    public class Day8 : SeparatedDay
    {

        public override object part1()
        {
            return SolvePart1();
        }

        public override object part2()
        {
            return SolvePart2();
        }

        private int SolvePart1()
        {
            var visitedIndexes = new HashSet<int>(Lines.Count);
            int accumulator = default;
            int index = default;
            do
            {
                if (visitedIndexes.Contains(index))
                {
                    return accumulator;
                }
                visitedIndexes.Add(index);

                var instructionSpan = Lines[index].AsSpan();
                switch (instructionSpan.Slice(0, 3))
                {
                    case ReadOnlySpan<char> op when op.SequenceEqual("nop"):
                    {
                        index++;
                        break;
                    }
                    case ReadOnlySpan<char> op when op.SequenceEqual("acc"):
                    {
                        index++;
                        accumulator += int.Parse(instructionSpan.Slice(4));
                        break;
                    }
                    case ReadOnlySpan<char> op when op.SequenceEqual("jmp"):
                    {
                        index += int.Parse(instructionSpan.Slice(4));
                        break;
                    }
                    default:
                        throw new InvalidOperationException();
                }
            } while (true);
        }

        private int SolvePart2()
        {
            var _input = Lines.ToArray();
            if (CanComplete(_input, out var accumulator))
                return accumulator;

            for (var i = 0; i < _input.Length; i++)
            {
                var codeSpan = _input[i].AsSpan();
                string oldOp;
                switch (codeSpan.Slice(0, 3))
                {
                    case ReadOnlySpan<char> op when op.SequenceEqual("nop"):
                    {
                        oldOp = _input[i];
                        _input[i] = $"jmp {codeSpan[4..].ToString()}";
                        if (CanComplete(_input, out accumulator))
                        {
                            return accumulator;
                        }
                        _input[i] = oldOp;
                        break;
                    }

                    case ReadOnlySpan<char> op when op.SequenceEqual("jmp"):
                    {
                        oldOp = _input[i];
                        _input[i] = $"nop {codeSpan[4..].ToString()}";
                        if (CanComplete(_input, out accumulator))
                        {
                            return accumulator;
                        }
                        _input[i] = oldOp;
                        break;
                    }
                }
            }

            throw new Exception("could not end");
        }

        private bool CanComplete(string[] code, out int accumulator)
        {
            int index = default;
            var visitedIndexes = new HashSet<int>(code.Length);
            accumulator = default;
            var instructionsCount = Lines.Count;
            do
            {
                if (visitedIndexes.Contains(index))
                {
                    return false;
                }
                visitedIndexes.Add(index);
                
                var instructionSpan = code[index].AsSpan();
                switch (instructionSpan.Slice(0, 3))
                {
                    case ReadOnlySpan<char> op when op.SequenceEqual("nop"):
                    {
                        index++;
                        break;
                    }
                    case ReadOnlySpan<char> op when op.SequenceEqual("acc"):
                    {
                        index++;
                        accumulator += int.Parse(instructionSpan.Slice(4));
                        break;
                    }
                    case ReadOnlySpan<char> op when op.SequenceEqual("jmp"):
                    {
                        visitedIndexes.Add(index);
                        index += int.Parse(instructionSpan.Slice(4));
                        break;
                    }
                    default:
                        throw new InvalidOperationException();
                }
                
                if (index == instructionsCount)
                {
                    return true;
                }
            } while (true);
        }
        
    }

}