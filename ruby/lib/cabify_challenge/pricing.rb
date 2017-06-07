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
  end
end
