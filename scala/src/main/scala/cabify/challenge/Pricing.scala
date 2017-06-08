package cabify.challenge

trait Pricing {
  def calculate(amount: Int): Double
}

object Pricing {

  case class Default(pricePerUnit: Double) extends Pricing {
    override def calculate(amount: Int): Double = amount * pricePerUnit
  }
}
