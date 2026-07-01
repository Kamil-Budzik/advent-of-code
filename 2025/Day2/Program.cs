var line = File.ReadAllText("./day2.txt").Trim();
var parts = line.Split(',');

long total = 0;

foreach (var part in parts)
{
    var ids = part.Split('-');
    long first = long.Parse(ids[0]);
    long last = long.Parse(ids[1]);

    for (long id = first; id <= last; id++)
    {
        if (isInvalid(id.ToString())) total += id;
    }
}

Console.WriteLine(total);

bool isInvalid(string id)
{
    if (id.Length % 2 != 0) return false;
    int half = id.Length / 2;
    return id[..half] == id[half..];
}
