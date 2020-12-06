using System;

namespace AOC_2020.Days
{

    public class Day5 : SeparatedDay
    {
        private const int Columns = 8;

        public override object part1()
        {
            return HighestSeatPart1();
        }

        public override object part2()
        {
            return FindMySeat();
        }

        private int HighestSeatPart1()
        {
            int highest = default;
            foreach (var seatCoding in Lines)
            {
                if (CalculateSeatIdBinary(seatCoding) is var seatId && seatId > highest)
                {
                    highest = seatId;
                }
            }

            return highest;
        }

        private int CalculateSeatId(string seatCode)
        {
            int column = default;
            int row = default;

            // get row value
            row += seatCode[0] == 'B' ? 64 : 0;
            row += seatCode[1] == 'B' ? 32 : 0;
            row += seatCode[2] == 'B' ? 16 : 0;
            row += seatCode[3] == 'B' ? 8 : 0;
            row += seatCode[4] == 'B' ? 4 : 0;
            row += seatCode[5] == 'B' ? 2 : 0;
            row += seatCode[6] == 'B' ? 1 : 0;

            // get column value
            column += seatCode[7] == 'R' ? 4 : 0;
            column += seatCode[8] == 'R' ? 2 : 0;
            column += seatCode[9] == 'R' ? 1 : 0;

            return row * Columns + column;
        }

        private object FindMySeat()
        {
            var TakenSeatIds = new int[Lines.Count];

            for (var i = 0; i < Lines.Count; i++)
            {
                TakenSeatIds[i] = CalculateSeatIdBinary(Lines[i]);
            }

            Array.Sort(TakenSeatIds);

            for (int i = 1, seatId = TakenSeatIds[0] + 1; i < TakenSeatIds.Length - 1; i++, seatId++)
            {
                if (seatId != TakenSeatIds[i])
                {
                    return seatId;
                }
            }

            return -1;
        }

        private int CalculateSeatIdBinary(string seatCode)
        {
            var id = 0;
            for (var i = 0; i < seatCode.Length; i++)
            {
                if (seatCode[i] == 'B' || seatCode[i] == 'R')
                {
                    id |= 1 << (9 - i);
                }
            }
            
            return id;
        }
    }

}