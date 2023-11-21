defmodule Async do
  def make_async_for(return_pid, func, count) do
    Enum.each(0..count, fn v -> make_async(self(), func, Integer.to_string(v), 1000) end)

    Enum.each(0..count, fn _v ->
      receive do
        :ok -> :ok
      end
    end)

    send(return_pid, :ok)
  end

  defp make_async(pid, func, id, duration) do
    spawn(fn ->
      func.(id, duration)
      send(pid, :ok)
    end)
  end
end
