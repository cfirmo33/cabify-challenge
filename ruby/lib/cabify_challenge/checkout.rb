module CabifyChallenge
  class Checkout

    DEFAULT_PRICING_RULES = {
      "VOUCHER" => Pricing::GetTwoPayOnePrice.new(5.00),
      "TSHIRT"  => Pricing::BulkDiscountPrice.new(20.00, 19.00, 3),
      "MUG"     => Pricing::DefaultPrice.new(7.50)
    }.freeze

    # Create a new checkout with the given pricing rules.
    #
    # @param pricing_rules [Hash<String, #calculate>] The pricing rules defining a price for each
    #   product; the price would be one of those declared in `Pricing` module
    def initialize(pricing_rules = DEFAULT_PRICING_RULES)
      @pricing_rules = pricing_rules

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
        rule = @pricing_rules[product_name]
        accum + rule.calculate(units)
      end
    end

    private

    def has_price_for?(product_name)
      @pricing_rules.key?(product_name)
    end
  end
end
