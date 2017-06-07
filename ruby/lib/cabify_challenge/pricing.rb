module CabifyChallenge
  module Pricing

    # A price policy that applies no discount at all.
    class DefaultPrice

      # Initialize this default price with the given price per unit.
      #
      # @param price_per_unit [Float]
      def initialize(price_per_unit)
        @price_per_unit = price_per_unit
      end

      # Calculate the price for the given units
      #
      # @param units [Integer]
      # @return Float
      def calculate(units)
        @price_per_unit * units
      end
    end

    # A price policy that applies a discount of 2-for-1.
    class GetTwoPayOnePrice

      # Initialize this price with the given price per unit.
      #
      # @param price_per_unit [Float]
      def initialize(price_per_unit)
        @price_per_unit = price_per_unit
      end

      # Calculate the price for the given units
      #
      # @param units [Integer]
      # @return Float
      def calculate(units)
        @price_per_unit * (units - units.to_i / 2)
      end
    end

    # A price policy that applies a discount for bulk purchases.
    class BulkDiscountPrice

      # Initialize this price with the given price per unit.
      #
      # @param price_per_unit [Float] The price per unit if bulk minimum is not reached
      # @param discount_price_per_unit [Float] The price per unit if bulk minimum is reached
      # @param bulk_minimum_units [Integer] The minimum units to consider a bulk purchase
      def initialize(price_per_unit, discount_price_per_unit, bulk_minimum_units)
        @price_per_unit = price_per_unit
        @discount_price_per_unit = discount_price_per_unit
        @bulk_minimum_units = bulk_minimum_units
      end

      # Calculate the price for the given units
      #
      # @param units [Integer]
      # @return Float
      def calculate(units)
        if units < @bulk_minimum_units
          @price_per_unit * units
        else
          @discount_price_per_unit * units
        end
      end
    end
  end
end
