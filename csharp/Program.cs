Action<string, int> syncConsoleWrite = (string id, int wait) =>
{
    System.Console.WriteLine($"Action with id {id} started");
    Thread.Sleep(wait * 1000);
    System.Console.WriteLine($"Action with id {id} stopped");
};

var allTasks = new List<Task>();

const int count = 100000;


var start = DateTime.Now;
for (var i = 0; i < count; i++)
{
    var id = i.ToString();
    allTasks.Add(Task.Run(() =>
    {
        syncConsoleWrite(id, 1000);
    }));
}

var elapsed = DateTime.Now.Subtract(start);
await Task.WhenAll(allTasks);

System.Console.WriteLine($"elapsed time {elapsed}");
