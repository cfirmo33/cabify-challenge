# Cabify Code Challenge - Elixir

This is a implementation of the Cabify Code Challenge [you may find here][1]
for the Elixir programming language. The proposed exercise is for Ruby, but this
has been adapted to use Elixir.

Please read the challenge before reading further to understand the rest of
the document and the code.

## Disclaimer

I've never coded in Elixir until now. I just received once a basic
introduction  to the language. This solution to the exercise is the result to
learn Elixir basics by reading the language guide. Fortunatelly, the concepts
behind it (functional programming, protocols, high order, etc) are so familiar
to me that I could complete the challenge in less than I had estimated. But...
do not expect perfect code here.

## How to run it

The code is organized as a Mix project. There are some unit tests to check the
goals of the challenge are satisfied. You can run them with `mix test`.

## Solution and design

Elixir is a pure functional programming language. That means mutability is
not an option in this land. Fortunatelly, instead of Duck Typing we can
have polymorphic code using protocols, which privides a higher grade of
formalism to define the interfaces.

The proposed design consists in defining a set of modules with their structs
to represent the pricing policy conforming a common protocol. These pricing
modules are declared in the `Pricing` module:

* The `Pricing.DefaultPrice`, which represents the regular pricing
approach: each item has a fixed cost.
* The `Pricing.GetTwoPayOne`, which applies a 2-for-1 discount.
* The `Pricing.BulkDiscount`, which applies a bulk discount.

Each module provides an struct that contains the necessary information to
apply the price to an amount of product units and calculate the applied price.
This is done through the `Calculator` protocol using the `calculate()`
function.

Appart from that, the `Checkout` module represents the checkout process
according to the definition in the Challenge. It is instantiated with a set
of price rules, which are basically a map that associates each product name
with the price to be applied (an instance from `Pricing` module structs).
Internally, `Checkout` has an struct that maintains a map to count the number
of times each product was scanned. These counters and the pricing rules are
used while calculating the total amount in the checkout.

## Design motivation

The initial naive approach would be to implement the `Checkout` module without
the presence of the `Pricing` types. Perhaps introducing these rules as
conditional statements while calculating the total amount. Nevertheless, this
wouldn't be a clever solution because adding new discount forms would be a
pain in the ass. We might end with dozens of conditional statements in the
class, a very illegible code with a huge cyclomatic complexity and really
hard to test it separated.

In constrast, the proposed solution allow us to add more `Pricing` modules
conforming to the `Pricing.Calculator` protocol to represent new forms of
calculating the cost of a product.

## Design caveats

Some complex form of pricing policies cannot be implemented with the proposed
solution. Particularly things like having two or more discount policies for
the same product or having a discount that depends on the purchased units of
another product are not possible. We would need a more complex interface in
the `Pricing` modules to support that.

## Performance

The `Checkout` class has the following Big O performance:

* `scan`: O(1) assuming insertions in Elixir maps are constant.
* `total`: O(n) assuming traversal of Elixir maps are linear. If required,
this could be reduced to O(1) by just accumulating the current total amount
of the cart when products are added in `scan()`.


[1]: https://gist.github.com/patriciagao/377dca8920ba3b1fc8da
