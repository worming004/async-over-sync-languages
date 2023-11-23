defmodule ElixirExample do
  # 1,625s
  @count 100_000

  # 1,002s
  # @count 200
  import Number.Delimit

  def main_timed do
    {time_in_microseconds, _ret_val} = :timer.tc(fn -> main() end)
    IO.puts("elapsed times in milliseconds #{number_to_delimited(time_in_microseconds / 1000)}")
  end

  def main do
    case Async.make_async_for(self(), &Sync.sync_write/2, @count) do
      :ok -> IO.puts("program ended now")
      _ -> IO.puts("something wrong happened")
    end
  end
end
