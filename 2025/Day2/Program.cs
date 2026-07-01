var line = File.ReadAllText("./day2.txt").Trim();
var parts = line.Split(',');

long totalPart1 = 0;
long totalPart2 = 0;

foreach (var part in parts)
{
    var ids = part.Split('-');
    long first = long.Parse(ids[0]);
    long last = long.Parse(ids[1]);

    for (long id = first; id <= last; id++)
    {
        if (isInvalidPart1(id.ToString())) totalPart1 += id;
        if (isInvalidPart2(id.ToString())) totalPart2 += id;
    }
}

Console.WriteLine(totalPart1);
Console.WriteLine(totalPart2);

bool isInvalidPart1(string id)
{
    if (id.Length % 2 != 0) return false;
    int half = id.Length / 2;
    return id[..half] == id[half..];
}

bool isInvalidPart2(string id)
{
    for (int unitLength = 1; unitLength <= id.Length / 2; unitLength++)
    {
        if (id.Length % unitLength != 0) continue;

        string unit = id[..unitLength];
        bool allMatch = true;

        for (int i = unitLength; i < id.Length; i += unitLength)
        {
            if (id[i..(i + unitLength)] != unit)
            {
                allMatch = false;
                break;
            }
        }

        if (allMatch) return true;
    }
    return false;
}
