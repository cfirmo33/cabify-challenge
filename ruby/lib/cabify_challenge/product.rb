require 'yaml'

module CabifyChallenge

  # A product from the online store
  class Product
    attr_accessor :code
    attr_accessor :description
    attr_accessor :price

    # Instantiate a new `Product` from its hash representation
    def self.from_hash(hash)
      new.tap do |prod|
        prod.code = hash[:code]
        prod.description = hash[:description]
        prod.price = Pricing.from_hash(hash)
      end
    end

    # A convenience method to calculate the price
    def calculate_price(units)
      price.calculate(units)
    end
  end

  # A catalog of products from the online store
  class ProductCatalog

    # Load the default catalog
    def self.default
      from_file(File.expand_path('../../resources/catalog.yml', __FILE__))
    end

    # Load catalog from the given file
    def self.from_file(file)
      products = YAML.load_file(file).map do |item|
        Product.from_hash(item)
      end
      new.tap do |cat|
        cat.add(products)
      end
    end

    # Initialize the catalog
    #
    # @param products [Array<Product>]
    def initialize(*products)
      @products = Hash.new
      add(products)
    end

    # Check whether the catalog is empty
    def empty?
      @products.empty?
    end

    # Add a new product to the catalog
    #
    # @param products [Array<Product>]
    def add(*products)
      products.flatten.each do |prod|
        @products[prod.code] = prod
      end
    end

    # Find a product in the catalog
    #
    # @param code [String]
    # @return [Product | nil]
    def find(code)
      @products[code]
    end

    # Check whether the catalog includes a product with the given code
    #
    # @param code [String]
    # @return [Boolean]
    def includes?(code)
      @products.include?(code)
    end
  end
end