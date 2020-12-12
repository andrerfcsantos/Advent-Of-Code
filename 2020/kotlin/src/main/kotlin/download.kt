import java.io.File
import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse
import java.nio.file.Paths




fun downloadInput(session: String, file: File, year: Int, day: Int): String {
    val input = getInputFile(session, year, day)

    file.bufferedWriter().use { out ->
        out.write(input)
    }

    return input
}


fun loadInputFile(file: File): String {
    file.bufferedReader().use {
        return it.readText()
    }
}

fun getInputFile(session: String, year: Int, day: Int): String {

    val client: HttpClient = HttpClient.newHttpClient()
    val requestUri = "https://adventofcode.com/${year}/day/${day}/input"

    val request: HttpRequest = HttpRequest.newBuilder()
        .uri(URI.create(requestUri))
        .header("cookie", "session=${session}")
        .GET()
        .build()
    val response = client.send(request, HttpResponse.BodyHandlers.ofString())

    return response.body()
}
