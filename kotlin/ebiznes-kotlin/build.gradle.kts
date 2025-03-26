plugins {
    kotlin("jvm") version "2.1.10"
    kotlin("plugin.serialization") version "1.8.21"
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-server-netty:2.3.1")
    implementation("io.ktor:ktor-server-core-jvm:2.3.1")

    implementation("io.ktor:ktor-server-core:2.3.1")
    implementation("io.ktor:ktor-server-content-negotiation-jvm:2.3.1")
    implementation("io.ktor:ktor-serialization-kotlinx-json-jvm:2.3.1")
    implementation("io.ktor:ktor-server-netty-jvm:2.3.1")
    implementation("dev.kord:kord-core:0.8.2")

    implementation("io.ktor:ktor-serialization-kotlinx-json:2.3.1")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.5.0")

    implementation("io.ktor:ktor-client-core:2.3.1")
    implementation("io.ktor:ktor-client-cio:2.3.1")

    implementation("ch.qos.logback:logback-classic:1.4.0")
}

tasks.test {
    useJUnitPlatform()
}
kotlin {
    jvmToolchain(23)
}