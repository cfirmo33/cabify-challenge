package cabify.challenge

import org.scalatest.{FlatSpec, Matchers}

class PricingTest extends FlatSpec with Matchers {

  "Default pricing" should "calculate the price for some amount" in {
    val pricing = Pricing.Default(1.95)

    pricing.calculate(0) shouldBe 0.0
    pricing.calculate(1) shouldBe 1.95
    pricing.calculate(2) shouldBe 3.90
  }
}
