package cabify.challenge

import akka.actor.ActorSystem
import org.scalatest.{FlatSpec, Matchers}

class CheckoutTest extends FlatSpec with Matchers {

  implicit val system = ActorSystem("test")

  "Checkout" should "meet expectations of scenario #1" in {
    val co = Checkout()
    Seq("VOUCHER", "TSHIRT", "MUG").foreach(co.scan)
    co.total shouldBe 32.50
  }

  it should "meet expectations of scenario #2" in {
    val co = Checkout()
    Seq("VOUCHER", "TSHIRT", "VOUCHER").foreach(co.scan)
    co.total shouldBe 25.00
  }

  it should "meet expectations of scenario #3" in {
    val co = Checkout()
    Seq("TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT").foreach(co.scan)
    co.total shouldBe 81.00
  }

  it should "meet expectations of scenario #4" in {
    val co = Checkout()
    Seq("VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT").foreach(co.scan)
    co.total shouldBe 74.50
  }
}
