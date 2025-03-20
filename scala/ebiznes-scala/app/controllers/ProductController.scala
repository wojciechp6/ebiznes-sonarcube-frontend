package controllers

import play.api.libs.json._
import play.api.mvc._
import javax.inject._
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents, Result}

case class Product(id: Int, name: String, price: Double)
object Product {
  implicit val productFormat: OFormat[Product] = Json.format[Product]
}

@Singleton
class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
  private var products = List(
    Product(1, "Laptop", 2000.0),
    Product(2, "Phone", 800.0)
  )

  def getAll = Action {
    Ok(Json.toJson(products))
  }

  def getById(id: Int) = Action {
    products.find(_.id == id)
      .map(product => Ok(Json.toJson(product)))
      .getOrElse(NotFound("Product not found"))
  }

def add: Action = Action { request =>
  val newProduct = Product(3, "da", 7.0)
  products = products :+ newProduct
  Created(Json.toJson(newProduct))
}




  def delete(id: Int) = Action {
    products = products.filterNot(_.id == id)
    NoContent
  }
}

