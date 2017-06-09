defmodule PricingTest do
  use ExUnit.Case
  doctest CabifyChallenge

  test "default price" do
    price = %Pricing.DefaultPrice{ price_per_unit: 1.95 }

    assert Pricing.Calculator.calculate(price, 0) == 0
    assert Pricing.Calculator.calculate(price, 1) == 1.95
    assert Pricing.Calculator.calculate(price, 2) == 3.90
  end
  
end
