require 'minitest/autorun'
require 'cabify_challenge'

module CabifyChallenge
  class PricingTest < Minitest::Test

    def test_default_price
      price = Pricing::DefaultPrice.new(1.95)

      assert_equal 0.0, price.calculate(0)
      assert_equal 1.95, price.calculate(1)
      assert_equal 3.90, price.calculate(2)
    end

    def test_get_two_pay_one_price
      price = Pricing::GetTwoPayOnePrice.new(1.95)

      assert_equal 0.0, price.calculate(0)
      assert_equal 1.95, price.calculate(1)
      assert_equal 1.95, price.calculate(2)
      assert_equal 3.90, price.calculate(3)
      assert_equal 3.90, price.calculate(4)
    end

    def test_bulk_discount_price
      price = Pricing::BulkDiscountPrice.new(1.95, 1.75, 3)

      assert_equal 0.0, price.calculate(0)
      assert_equal 1.95, price.calculate(1)
      assert_equal 3.90, price.calculate(2)
      assert_equal 5.25, price.calculate(3)
      assert_equal 7.0, price.calculate(4)
    end
  end
end
