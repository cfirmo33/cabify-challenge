defmodule CheckoutTest do
  use ExUnit.Case
  doctest CabifyChallenge

  test "scenary #1" do
    co = %Checkout {}
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("MUG")

    assert Checkout.total(co) == 32.50
  end

  test "scenary #2" do
    co = %Checkout {}
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("VOUCHER")

    assert Checkout.total(co) == 25.00
  end

  test "scenary #3" do
    co = %Checkout {}
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("TSHIRT")

    assert Checkout.total(co) == 81.00
  end

  test "scenary #4" do
    co = %Checkout {}
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("VOUCHER")
      |> Checkout.scan("MUG")
      |> Checkout.scan("TSHIRT")
      |> Checkout.scan("TSHIRT")

    assert Checkout.total(co) == 74.50
  end
end
