require 'minitest/unit'
require 'minitest/autorun'

require 'cabify_challenge'

module CabifyChallenge
  class CheckoutTest < Minitest::Unit::TestCase

    def test_checkout_scenary_1
      co = Checkout.new
      %w(VOUCHER TSHIRT MUG).each do |product_name|
        co.scan(product_name)
      end
      assert_equal 32.50, co.total
    end

    def test_checkout_scenary_2
      co = Checkout.new
      %w(VOUCHER TSHIRT VOUCHER).each do |product_name|
        co.scan(product_name)
      end
      assert_equal 25.00, co.total
    end

    def test_checkout_scenary_3
      co = Checkout.new
      %w(TSHIRT TSHIRT TSHIRT VOUCHER TSHIRT).each do |product_name|
        co.scan(product_name)
      end
      assert_equal 81.00, co.total
    end

    def test_checkout_scenary_4
      co = Checkout.new
      %w(VOUCHER TSHIRT VOUCHER VOUCHER MUG TSHIRT TSHIRT).each do |product_name|
        co.scan(product_name)
      end
      assert_equal 74.50, co.total
    end
  end
end
