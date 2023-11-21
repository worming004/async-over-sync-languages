Action<string, int> syncConsoleWrite = (string id, int wait) =>
{
    System.Console.WriteLine($"Action with id {id} started");
    Thread.Sleep(wait);
    System.Console.WriteLine($"Action with id {id} stopped");
};

var allTasks = new List<Task>();

// const int count = 100000; // 100000 is too long ..
const int count = 200;

var start = DateTime.Now;
for (var i = 0; i < count; i++)
{
    var id = i.ToString();
    allTasks.Add(Task.Run(() =>
    {
        syncConsoleWrite(id, 1000);
    }));
}

await Task.WhenAll(allTasks);
var elapsed = DateTime.Now.Subtract(start);

System.Console.WriteLine($"elapsed time {elapsed}");
