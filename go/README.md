# Cabify Code Challenge - Go

This is a implementation of the Cabify Code Challenge [you may find here][1]
for the Go programming language.

Please read the challenge before reading further to understand the rest of
the document and the code.

## Disclaimer

Last time I wrote Go code was in ~2012. And it was just a little prototype I
never ended because Go runtime didn't met the requirements (I needed to embed
it in a C++ process as a DLL plugin, and that was not possible at that time).
So, I have almost no experience in the language. Fortunately, Go is so simple
and the concepts behind it ([corutines][2], [CSP][4], [structural typing][3])
are so familiar to me that I could complete the challenge in less than I had
estimated. But... do not expect perfect code here.

## How to run it

The code is organized in as single package. There are some unit tests to check the
goals of the challenge are satisfied. The code depends on `testify` package. You
can install it with `go get github.com/stretchr/testify`. After that, you can
execute the tests with `go test`.

## Solution and design

The proposed challenge was clearly focused on showing how the concurrency
primitives of Go can be used, particularly gorutines and channels. This
solution is based on modelling the different pricing policies defined in
the `pricing.go` source file. First, a common interface `PriceCalculator`
is defined to handle the price policies polymorphically. After that, a set
of policy types are declared that conforms that interface:

* The `DefaultPrice`, which represents the regular pricing
approach: each item has a fixed cost.
* The `GetTwoPayOnePrice`, which applies a 2-for-1 discount.
* The `BulkDiscountPrice`, which applies a bulk discount.

Each type is instantiated with the necessary information to apply the price
to an amount of product units and calculate the applied price.

Appart from that, the `Checkout` type in `checkout.go` represents the checkout 
process according to the definition in the Challenge. Since the challenge says 
it is mandatory to make it thread-safe, under the hoods it uses a gorutine and
some channels to communicate with the client code that requests `Scan()` and
`Total()` operations. Two channels are used:
* A `chan string` to submit scan actions
* A `chan chan float64` to request total calculation. The inner channel passed
to the outer channel is used to receive a response from the gorutine with the
computed result.

The gorutine selects events from both channels and manages the requests 
accordingly.

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

In the other hand, using channels is the simplest and idiomatic way to have
thread-safety in Go. Nothing new under the sun.

## Design caveats

Some complex form of pricing policies cannot be implemented with the proposed
solution. Particularly things like having two or more discount policies for
the same product or having a discount that depends on the purchased units of
another product are not possible. We would need a more complex interface in
the pricing classes to support that.

## Performance

The `Checkout` class has the following [Big O][3] performance:

* `Scan`: O(1) assuming insertions in Go maps are constant.
* `Total`: O(n) assuming traversal of Go maps are linear. If required,
this could be reduced to O(1) by just accumulating the current total amount
of the cart when products are added in `Scan()`.

## Other consideratings

* The gorutine under `Checkout` never ends. If necessary, we could provide
an additional channel to request the goroutine to terminate.
* The package structure could be lightly inconventional. As disclaimed
above, I'm not a Go expert at all.


[1]: https://gist.github.com/samlown/f7347775af429aaf9afb
[2]: https://en.wikipedia.org/wiki/Coroutine
[3]: https://en.wikipedia.org/wiki/Structural_type_system
[4]: https://en.wikipedia.org/wiki/Communicating_sequential_processes

