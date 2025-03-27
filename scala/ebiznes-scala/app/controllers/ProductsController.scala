package controllers

import models.Product
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class ProductsController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val productFormat: OFormat[Product] = Json.format[Product]

  // Tymczasowa lista produktów
  private val products = ListBuffer(
    Product(1, "Produkt A", 100.0),
    Product(2, "Produkt B", 200.0)
  )

  // Endpoint: Listowanie wszystkich produktów
  def listProducts: Action[AnyContent] = Action {
    Ok(Json.toJson(products))
  }

  // Endpoint: Wyświetlenie produktu po ID
  def getProduct(id: Int): Action[AnyContent] = Action {
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(s"Produkt o ID $id nie znaleziony.")
    }
  }

  // Endpoint: Dodanie nowego produktu
  def addProduct: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest,
      product => {
        products += product
        Created(Json.toJson(product))
      }
    )
  }

  // Endpoint: Aktualizacja produktu
  def updateProduct(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest,
      updatedProduct => {
        products.indexWhere(_.id == id) match {
          case -1 => NotFound(s"Produkt o ID $id nie znaleziony.")
          case index =>
            products.update(index, updatedProduct)
            Ok(Json.toJson(updatedProduct))
        }
      }
    )
  }

  // Endpoint: Usunięcie produktu
  def deleteProduct(id: Int): Action[AnyContent] = Action {
    products.indexWhere(_.id == id) match {
      case -1 => NotFound(s"Produkt o ID $id nie znaleziony.")
      case index =>
        products.remove(index)
        NoContent
    }
  }
}
