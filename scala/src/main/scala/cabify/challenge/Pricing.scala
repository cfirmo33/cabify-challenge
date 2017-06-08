package cabify.challenge

trait Pricing {
  def calculate(amount: Int): Double
}

object Pricing {

  case class Default(pricePerUnit: Double) extends Pricing {
    override def calculate(amount: Int): Double = amount * pricePerUnit
  }

  case class GetTwoPayOne(pricePerUnit: Double) extends Pricing {
    override def calculate(amount: Int): Double =
      pricePerUnit * (amount - amount / 2)
  }

  case class BulkDiscount(pricePerUnit: Double,
                          discountPricePerUnit: Double,
                          minBulkUnits: Int) extends Pricing {
    override def calculate(amount: Int): Double =
      if (amount < minBulkUnits) pricePerUnit * amount
      else discountPricePerUnit * amount
  }
}
