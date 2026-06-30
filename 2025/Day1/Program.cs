
int currDial = 50;
int total = 0;


foreach (var line in File.ReadLines("./day1.txt"))
{
    char direction = line[0];
    int steps = int.Parse(line[1..]);
    currDial = CalcTheDial(direction, steps);
    // answer to part 1
    // if (currDial == 0) total++;
}

Console.WriteLine(total);


int CalcTheDial(char dir, int num)
{
    int value = currDial;

    for (int i = 0; i < num; i++)
    {

        if (dir == 'L' && value == 0)
        {
            value = 99;
            continue;
        }

        if (dir == 'R' && value == 99)
        {
            value = 0;
            total++;
            continue;
        }


        if (dir == 'L')
        {
            value -= 1;
            if (value == 0) total++;
        }
        else
        {
            value += 1;
        }


    }


    return value;
}
