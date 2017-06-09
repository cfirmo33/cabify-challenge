# Cabify Code Challenge - Rust

This is a implementation of the Cabify Code Challenge [you may find here][1]
for the Rust programming language. The proposed exercise is for Go, but this
has been adapted to use Rust.

Please read the challenge before reading further to understand the rest of
the document and the code.

## How to run it

First, you will Rust to be installed in your system. You can do it by just
executing `curl https://sh.rustup.rs -sSf | sh` and following the
instructions.

The code is a Rust library with some unit tests to check the goals of the
challenge are satisfied. You can execute the tests with `cargo test` (Cargo
is the Rust building tool).

## Solution and design

The original challenge was clearly focused on showing how the concurrency
primitives of Go can be used, particularly gorutines and channels. This
implementation respect the same constraints: make the `Checkout` object
thread-safe. To do that, we use [MPSC (multiple-producer single-consumer)][3]
channels provided by Rust standard library. They are very similar to
the Go channels, but they use regular threads instead of coroutines.

As the rest of solutions for this challenge, the different pricing policies
are declared separately. See `pricing.rust` to find the following types:

* The `DefaultPrice`, which represents the regular pricing
approach: each item has a fixed cost.
* The `GetTwoPayOne`, which applies a 2-for-1 discount.
* The `BulkDiscount`, which applies a bulk discount.

Each type is instantiated with the necessary information to apply the price
to an amount of product units and calculate the applied price. All them
satisfy a `Pricing` trait (a fucking awesome [typeclass mechanism][4] as the one
you have in Haskell) so they can handled polymorphically.

Appart from that, the `Checkout` type in `Checkout.scala` represents the checkout
process according to the definition in the Challenge. Since the challenge says 
it is mandatory to make it thread-safe, under the hoods it uses MPSC channels
to communicate from the `scan()` and `total()` methods with the handler thread.
The main difference here is that _select_ primitive is [still unsupported][5] in Rust.
So you have to define an algebraic abstract data type to describe the messages
that will be received by the handler thread. Thanks to the pattern matching,
this is quite easy and clear :-)

## Design motivation

The initial naive approach would be to implement the `Checkout` type without
the presence of the types declare in `pricing.rs`. Perhaps introducing these
rules as conditional statements while calculating the total amount.
Nevertheless, this wouldn't be a clever solution because adding new discount
forms would be a pain in the ass. We might end with dozens of conditional
statements in the type, a very illegible code with a huge cyclomatic
complexity and really hard to test it separated.

In constrast, the proposed solution allow us to add more pricing types
to represent new forms of calculating the cost of a product.

In the other hand, we use again channels to communicate threads so we obtain
thread-safety in the checkout object.

## Design caveats

Some complex form of pricing policies cannot be implemented with the proposed
solution. Particularly things like having two or more discount policies for
the same product or having a discount that depends on the purchased units of
another product are not possible. We would need a more complex interface in
the pricing classes to support that.

## Performance

The `Checkout` class has the following Big O performance:

* `scan()`: O(1) assuming insertions in Rust hashmaps are constant.
* `total()`: O(n) assuming traversal of Rust hashmaps are linear. If required,
this could be reduced to O(1) by just accumulating the current total amount
of the cart when products are added in `scan()`.

## Other considerations

* This time the handler thread terminates. Otherwise, the main thread would end
before it, invoking the [destructors][6] (yes, Rust have destructors as C++) that
cause the sender part of the channel to be closed. That causes an unexpected
error in the handler thread while receiving from the channel, with an ugly
message in the console. So in this implementation we just close the handler
thread once `total()` is invoked.



[1]: https://gist.github.com/samlown/f7347775af429aaf9afb
[2]: https://www.rust-lang.org/en-US/
[3]: https://doc.rust-lang.org/std/sync/mpsc/
[4]: https://en.wikipedia.org/wiki/Type_class
[5]: https://doc.rust-lang.org/std/macro.select.html
[6]: https://doc.rust-lang.org/std/ops/trait.Drop.html
