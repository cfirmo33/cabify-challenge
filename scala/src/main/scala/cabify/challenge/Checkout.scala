package cabify.challenge

import scala.concurrent.Await
import scala.concurrent.duration._

import akka.actor.{Actor, ActorRef, ActorSystem, Props}
import akka.pattern.ask
import akka.util.Timeout

class Checkout private (actor: ActorRef) {

  private implicit val timeout = Timeout(1.second)

  def scan(productName: String): Unit = {
    actor ! Checkout.ScanMessage(productName)
  }

  def total: Double = {
    Await.result((actor ? Checkout.TotalMessage).mapTo[Double], Duration.Inf)
  }
}

object Checkout {

  def apply(rules: Pricing.Rules = Pricing.DefaultRules)
           (implicit system: ActorSystem): Checkout = {
    val props = Props(new HandlerActor(rules))
    val actor = system.actorOf(props)
    new Checkout(actor)
  }

  case class ScanMessage(productName: String)
  case object TotalMessage

  class HandlerActor(rules: Pricing.Rules) extends Actor {

    private var cart: Map[String, Int] = Map.empty.withDefaultValue(0)

    override def receive: Receive = {
      case ScanMessage(productName) =>
        cart += productName -> (cart(productName) + 1)
      case TotalMessage =>
        sender ! cart.foldLeft(0.0) {
          case (total, (product, units)) =>
            total + rules(product).calculate(units)
        }
    }
  }
}
