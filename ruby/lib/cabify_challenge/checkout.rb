module CabifyChallenge
  class Checkout

    # Create a new checkout with the given product catalog
    #
    # @param product_catalog [ProductCatalog]
    def initialize(product_catalog = ProductCatalog.default)
      @product_catalog = product_catalog

      # This cart maintains a relation of product names and the number of units scanned.
      # It is defaulted to 0 to facilitate the scan of products
      @cart = Hash.new(0)
    end

    # Scan a new product and put it in the shopping cart.
    #
    # @param product_name [String]
    def scan(product_name)
      raise ArgumentError.new("No price for '#{product_name}'") unless has_price_for?(product_name)
      # Thanks to the default value of the cart hash, we just have to increment it.
      @cart[product_name] += 1
    end

    # Calculate the total price of the scanned products.
    #
    # @return [Float]
    def total
      @cart.inject(0) do |accum, (product_name, units)|
        product = @product_catalog.find(product_name)
        accum + product.calculate_price(units)
      end
    end

    private

    def has_price_for?(product_name)
      @product_catalog.includes?(product_name)
    end
  end
end
