package cabify.challenge

trait Pricing {
  def calculate(amount: Int): Double
}

object Pricing {

  type Rules = Map[String, Pricing]

  val DefaultRules = Map(
    "VOUCHER" -> GetTwoPayOne(pricePerUnit = 5.00),
    "TSHIRT" -> BulkDiscount(pricePerUnit = 20.00,
                             discountPricePerUnit = 19.00,
                             minBulkUnits = 3),
    "MUG" -> Default(pricePerUnit = 7.50)
  )

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
