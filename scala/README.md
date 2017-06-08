# Cabify Code Challenge - Scala

This is a implementation of the Cabify Code Challenge [you may find here][1]
for the Scala programming language. The proposed exercise is for Go, but this
has been adapted to use Scala.

Please read the challenge before reading further to understand the rest of
the document and the code.

## How to run it

First, you will need SBT (a build tool) installed. You can do that in OSX with
just `brew install sbt`. The code is a SBT artifact with some unit tests to
check the goals of the challenge are satisfied. You can execute the tests with
`sbt test`. Please be patient. Brewed SBT is relatively light. The Scala
compiler, runtime, and the SBT core itself has to be downloaded.

## Solution and design

The original challenge was clearly focused on showing how the concurrency
primitives of Go can be used, particularly gorutines and channels. This
implementation respect the same constraints: make the `Checkout` object
thread-safe. But, instead of corutines and channels, here I used an
[actor model][2].

As the rest of solutions for this challenge, the different pricing policies
are declared separately. See `Pricing.scala` to find the following types:

* The `Pricing.Default`, which represents the regular pricing
approach: each item has a fixed cost.
* The `Pricing.GetTwoPayOne`, which applies a 2-for-1 discount.
* The `Pricing.BulkDiscount`, which applies a bulk discount.

Each type is instantiated with the necessary information to apply the price
to an amount of product units and calculate the applied price. All them
satisfy a `Pricing` trait (similar to a Java interface, but more powerful)
so they can handled polymorphically.

Appart from that, the `Checkout` type in `Checkout.scala` represents the checkout
process according to the definition in the Challenge. Since the challenge says 
it is mandatory to make it thread-safe, under the hoods it uses an actor we
can communicate from the `scan()` and `total()` methods. By sending some
predefined messages, we can request the actor to perform the actions we need.

## Design motivation

The initial naive approach would be to implement the `Checkout` type without
the presence of the pricing classes. Perhaps introducing these rules as
conditional statements while calculating the total amount. Nevertheless, this
wouldn't be a clever solution because adding new discount forms would be a
pain in the ass. We might end with dozens of conditional statements in the
type, a very illegible code with a huge cyclomatic complexituy and really
hard to test it separated.

In constrast, the proposed solution allow us to add more pricing classes
to represent new forms of calculating the cost of a product.

In the other hand, using an actor model is a very simple way to have
thread-safety in Scala.

## Design caveats

Some complex form of pricing policies cannot be implemented with the proposed
solution. Particularly things like having two or more discount policies for
the same product or having a discount that depends on the purchased units of
another product are not possible. We would need a more complex interface in
the pricing classes to support that.

## Performance

The `Checkout` class has the following Big O performance:

* `scan()`: O(1) assuming insertions in Scala maps are constant.
* `total()`: O(n) assuming traversal of Scala maps are linear. If required,
this could be reduced to O(1) by just accumulating the current total amount
of the cart when products are added in `scan()`.

## Other considerations

* The actor under `Checkout` never ends. If necessary, we could provide
an additional message to request the actor to terminate.



[1]: https://gist.github.com/samlown/f7347775af429aaf9afb
[2]: https://en.wikipedia.org/wiki/Actor_model

