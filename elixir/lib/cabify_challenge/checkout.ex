defmodule Checkout do
  @default_rules %{
    "VOUCHER" => %Pricing.GetTwoPayOne { price_per_unit: 5.00 },
    "TSHIRT" => %Pricing.BulkDiscount {
     price_per_unit: 20.00,
     discount_price_per_unit: 19.00,
     min_bulk_units: 3
    },
    "MUG" => %Pricing.DefaultPrice { price_per_unit: 7.50 }
  }

  defstruct rules: @default_rules, cart: %{}

  def scan(co, product_name) do
    %Checkout {
      rules: co.rules,
      cart: Map.put(co.cart, product_name, Map.get(co.cart, product_name, 0) + 1)
    }
  end

  def total(co) do
    Enum.reduce(co.cart, 0.0, fn({k, v}, acc) ->
      price = Map.get(co.rules, k)
      acc + Pricing.Calculator.calculate(price, v)
    end)
  end
end
