package controllers

import models.Category
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class CategoriesController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val categoryFormat: OFormat[Category] = Json.format[Category]

  // Tymczasowa lista kategorii
  private val categories = ListBuffer(
    Category(1, "Kategoria A"),
    Category(2, "Kategoria B")
  )

  // Endpoint: Listowanie wszystkich kategorii
  def listCategories: Action[AnyContent] = Action {
    Ok(Json.toJson(categories))
  }

  // Endpoint: Wyświetlenie kategorii po ID
  def getCategory(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(s"Kategoria o ID $id nie znaleziono.")
    }
  }

  // Endpoint: Dodanie nowej kategorii
  def addCategory: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest,
      category => {
        categories += category
        Created(Json.toJson(category))
      }
    )
  }

  // Endpoint: Aktualizacja kategorii
  def updateCategory(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest,
      updatedCategory => {
        categories.indexWhere(_.id == id) match {
          case -1 => NotFound(s"Kategoria o ID $id nie znaleziono.")
          case index =>
            categories.update(index, updatedCategory)
            Ok(Json.toJson(updatedCategory))
        }
      }
    )
  }

  // Endpoint: Usunięcie kategorii
  def deleteCategory(id: Int): Action[AnyContent] = Action {
    categories.indexWhere(_.id == id) match {
      case -1 => NotFound(s"Kategoria o ID $id nie znaleziono.")
      case index =>
        categories.remove(index)
        NoContent
    }
  }
}
