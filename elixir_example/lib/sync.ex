defmodule Sync do
  def sync_write(id, duration) do
    IO.puts("Start with ID #{id}")
    :timer.sleep(duration)
    IO.puts("Stop with ID #{id}")
  end
end
