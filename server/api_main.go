package main

import (
    "final-project/kelas-beta-golang/proto"
    "context"
    "github.com/gofiber/fiber/v2"
    "google.golang.org/grpc"
    "net/http"
)

func main() {
    app := fiber.New()

    // Menghubungkan ke gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        app.SendStatus(http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    // Menginisialisasi client gRPC
    client := proto.NewGreetServiceClient(conn)

    // Membuat endpoint HTTP
    app.Post("/greet", func(c *fiber.Ctx) error {
        // Mendapatkan data dari request HTTP
        var request proto.GreetRequest
        if err := c.BodyParser(&request); err != nil {
            return err
        }

        // Memanggil gRPC server menggunakan client
        response, err := client.Greet(context.Background(), &request)
        if err != nil {
            return err
        }

        // Mengirimkan response ke client HTTP
        return c.JSON(response)
    })

    // Menjalankan aplikasi Fiber
    app.Listen(":3000")
}
