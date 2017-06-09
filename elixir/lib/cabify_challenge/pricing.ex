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

  defmodule GetTwoPayOne do
    defstruct price_per_unit: 0
  end

  defimpl Calculator, for: GetTwoPayOne do
    def calculate(price, units), do: price.price_per_unit * (units - div(units, 2))
  end

  defmodule BulkDiscount do
    defstruct price_per_unit: 0,
              discount_price_per_unit: 0,
              min_bulk_units: 0
  end

  defimpl Calculator, for: BulkDiscount do
    def calculate(price, units) do
      if units < price.min_bulk_units do
        units * price.price_per_unit
      else
        units * price.discount_price_per_unit
      end
    end
  end
end
