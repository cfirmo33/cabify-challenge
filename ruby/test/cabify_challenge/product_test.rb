require 'minitest/autorun'
require 'cabify_challenge'

module CabifyChallenge
  class ProductTest < Minitest::Test
    def test_product_from_hash
      prod = Product.from_hash(
        code: "VOUCHER",
        description: "Cabify Voucher",
        price: 5.00,
        discount: "2x1")

      assert_equal "VOUCHER", prod.code
      assert_equal "Cabify Voucher", prod.description
      assert prod.price.is_a?(Pricing::GetTwoPayOnePrice)
    end

    def test_new_product_catalog_is_empty
      assert ProductCatalog.new.empty?
    end

    def test_product_catalog_find
      prod = Product.from_hash(code: "VOUCHER", description: "Cabify Voucher", price: 5.00)
      cat = ProductCatalog.new(prod)

      assert_equal "Cabify Voucher", cat.find("VOUCHER").description
      assert_nil cat.find("NAZGUL")
    end

    def test_product_catalog_includes
      prod = Product.from_hash(code: "VOUCHER", description: "Cabify Voucher", price: 5.00)
      cat = ProductCatalog.new(prod)

      assert cat.includes?("VOUCHER")
      refute cat.includes?("NAZGUL")
    end

    def test_default_product_catalog_is_not_empty
      cat = ProductCatalog.default

      refute cat.empty?
    end

  end
end
