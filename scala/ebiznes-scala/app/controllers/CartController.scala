package controllers

import models.CartItem
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val cartItemFormat: OFormat[CartItem] = Json.format[CartItem]

  // Tymczasowa lista koszyka
  private val cart = ListBuffer(
    CartItem(1, 2),
    CartItem(2, 1)
  )

  // Endpoint: Listowanie zawartości koszyka
  def listCartItems: Action[AnyContent] = Action {
    Ok(Json.toJson(cart))
  }

  // Endpoint: Dodanie produktu do koszyka
  def addCartItem: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest,
      cartItem => {
        cart += cartItem
        Created(Json.toJson(cartItem))
      }
    )
  }

  // Endpoint: Aktualizacja ilości produktu w koszyku
  def updateCartItem(productId: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest,
      updatedCartItem => {
        cart.indexWhere(_.productId == productId) match {
          case -1 => NotFound(s"Produkt o ID $productId nie znaleziony w koszyku.")
          case index =>
            cart.update(index, updatedCartItem)
            Ok(Json.toJson(updatedCartItem))
        }
      }
    )
  }

  // Endpoint: Usunięcie produktu z koszyka
  def deleteCartItem(productId: Int): Action[AnyContent] = Action {
    cart.indexWhere(_.productId == productId) match {
      case -1 => NotFound(s"Produkt o ID $productId nie znaleziony w koszyku.")
      case index =>
        cart.remove(index)
        NoContent
    }
  }
}
