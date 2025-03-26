import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

@Serializable
data class Product(val name: String, val category: String, val price: Int)

val categories = listOf("electronics", "furniture", "toys", "clothing")
val products = listOf(
    Product("Smartphone", "electronics", 800),
    Product("Laptop", "electronics", 1200),
    Product("Sofa", "furniture", 500),
    Product("Dining Table", "furniture", 750),
    Product("Action Figure", "toys", 30),
    Product("Board Game", "toys", 45),
    Product("Jeans", "clothing", 40),
    Product("T-shirt", "clothing", 20)
)

suspend fun main() {
    val kord = Kord("secret_token")
    val channelId = "1354426572140056749"

    setupEventListeners(kord)
    startServer(kord, channelId)

    kord.login {
        @OptIn(PrivilegedIntent::class)
        intents += Intent.MessageContent
    }
}

fun setupEventListeners(kord: Kord) {
    kord.on<MessageCreateEvent> {
        if (message.author?.isBot != false) return@on

        val command = message.content.takeIf { it.startsWith('!') }?.substringAfter('!')
        when (command) {
            "categories" -> message.channel.createMessage(Json.encodeToString(categories))
            in categories -> {
                val filteredProducts = products.filter { it.category == command }
                message.channel.createMessage(Json.encodeToString(filteredProducts))
            }
            else -> message.channel.createMessage("Category not found")
        }

        println("Message received from ${message.author?.tag}: ${message.content}")
    }
}

fun startServer(kord: Kord, channelId: String) {
    embeddedServer(Netty, port = 8080) {
        messageSenderModule(kord, channelId)
    }.start(wait = false)
}

fun Application.messageSenderModule(kord: Kord, channelId: String) {
    routing {
        post("/send") {
            val message = call.receive<String>()
            kord.rest.channel.createMessage(Snowflake(channelId)) {
                content = message
            }
            call.respond("Message was sent to channel with ID $channelId.")
        }
    }
}
