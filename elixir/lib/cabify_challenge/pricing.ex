defmodule Pricing do

  defprotocol Calculator do
    def calculate(price, units)
  end

  defmodule DefaultPrice do
    defstruct price_per_unit: 0
  end

  defimpl Calculator, for: DefaultPrice do
    def calculate(price, units), do: price.price_per_unit * units
  end
end
