# Cabify Code Challenge - Ruby

This is a implementation of the Cabify Code Challenge [you may find here][1]
for the Ruby programming language.

Please read the challenge before reading further to understand the rest of
the document and the code.

## How to run it

The code is organized as a Ruby gem. There are some unit tests to check the
goals of the challenge are satisfied. You can run them with `rake test`.

## Solution and design

Ruby is a dynamic-typing, object oriented programming language with a pinch
of functional programming features. This solution is based on using the OOP
features through [Duck Typing][2] with some additional usage of internal
iteration (monadic????) on standard collections (particularly, `each` and `inject`).

The proposed design consists in defining a set of classes to represent the
pricing policy conforming a common Duck Typing interface. These pricing
classes are declared in the `Pricing` module:

* The `Pricing::DefaultPrice`, which represents the regular pricing
approach: each item has a fixed cost.
* The `Pricing::GetTwoPayOnePrice`, which applies a 2-for-1 discount.
* The `Pricing::BulkDiscountPrice`, which applies a bulk discount.

Each class is instantiated with the necessary information to apply the price
to an amount of product units and calculate the applied price. This is done
through the `calculate()` method.

Appart from that, the `Checkout` class represents the checkout process
according to the definition in the Challenge. It is instantiated with a set
of price rules, which are basically a hash that associates each product name
with the price to be applied (an instance from `Pricing` classes). Internally,
the `Checkout` object maintains a hash to count the number of times each
product was scanned. These counters and the pricing rules are used while
calculating the total amount in the checkout.

## Design motivation

The initial naive approach would be to implement the `Checkout` class without
the presence of the `Pricing` classes. Perhaps introducing these rules as
conditional statements while calculating the total amount. Nevertheless, this
wouldn't be a clever solution because adding new discount forms would be a
pain in the ass. We might end with dozens of conditional statements in the
class, a very illegible code with a huge cyclomatic complexituy and really
hard to test it separated.

In constrast, the proposed solution allow us to add more `Pricing` classes
to represent new forms of calculating the cost of a product.

## Design caveats

Some complex form of pricing policies cannot be implemented with the proposed
solution. Particularly things like having two or more discount policies for
the same product or having a discount that depends on the purchased units of
another product are not possible. We would need a more complex interface in
the `Pricing` classes to support that.

## Performance

The `Checkout` class has the following [Big O][3] performance:

* `scan`: O(1) assuming insertions in Ruby hashes are constant.
* `total`: O(n) assuming traversal of Ruby hashes are linear. If required,
this could be reduced to O(1) by just accumulating the current total amount
of the cart when products are added in `scan()`.

## Other consideratings

The resulting gem is intentionally simple. It is not integrated with Bundler
since it was not necessary to demonstrate the goals of the Challenge.

[1]: https://gist.github.com/patriciagao/377dca8920ba3b1fc8da
[2]: https://en.wikipedia.org/wiki/Duck_typing
[3]: https://en.wikipedia.org/wiki/Big_O_notation
