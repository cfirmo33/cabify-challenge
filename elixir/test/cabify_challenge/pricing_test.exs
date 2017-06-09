defmodule PricingTest do
  use ExUnit.Case
  doctest CabifyChallenge

  test "default price" do
    price = %Pricing.DefaultPrice{ price_per_unit: 1.95 }

    assert Pricing.Calculator.calculate(price, 0) == 0
    assert Pricing.Calculator.calculate(price, 1) == 1.95
    assert Pricing.Calculator.calculate(price, 2) == 3.90
  end

  test "get-two-pay-one price" do
    price = %Pricing.GetTwoPayOne{ price_per_unit: 1.95 }

    assert Pricing.Calculator.calculate(price, 0) == 0
    assert Pricing.Calculator.calculate(price, 1) == 1.95
    assert Pricing.Calculator.calculate(price, 2) == 1.95
    assert Pricing.Calculator.calculate(price, 3) == 3.90
    assert Pricing.Calculator.calculate(price, 4) == 3.90
  end

  test "bulk discount price" do
    price = %Pricing.BulkDiscount {
      price_per_unit: 1.95,
      discount_price_per_unit: 1.75,
      min_bulk_units: 3
    }

    assert Pricing.Calculator.calculate(price, 0) == 0
    assert Pricing.Calculator.calculate(price, 1) == 1.95
    assert Pricing.Calculator.calculate(price, 2) == 3.90
    assert Pricing.Calculator.calculate(price, 3) == 5.25
    assert Pricing.Calculator.calculate(price, 4) == 7.00
  end
  
end
