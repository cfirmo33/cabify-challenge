package cabify.challenge

import org.scalatest.{FlatSpec, Matchers}

class PricingTest extends FlatSpec with Matchers {

  "Default pricing" should "calculate the price for some amount" in {
    val pricing = Pricing.Default(1.95)

    pricing.calculate(0) shouldBe 0.0
    pricing.calculate(1) shouldBe 1.95
    pricing.calculate(2) shouldBe 3.90
  }

  "Get-two-pay-one" should "calculate the price for some amount" in {
    val pricing = Pricing.GetTwoPayOne(1.95)

    pricing.calculate(0) shouldBe 0.0
    pricing.calculate(1) shouldBe 1.95
    pricing.calculate(2) shouldBe 1.95
    pricing.calculate(3) shouldBe 3.90
    pricing.calculate(4) shouldBe 3.90
  }

  "Bulk discount" should "calculate the price for some amount" in {
    val pricing = Pricing.BulkDiscount(1.95, 1.75, 3)

    pricing.calculate(0) shouldBe 0.0
    pricing.calculate(1) shouldBe 1.95
    pricing.calculate(2) shouldBe 3.90
    pricing.calculate(3) shouldBe 5.25
    pricing.calculate(4) shouldBe 7.00
  }
}
