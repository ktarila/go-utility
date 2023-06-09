module utilityapp.ktarila/go-stimulus

go 1.18

require (
	github.com/gofiber/fiber/v2 v2.32.0
	github.com/gofiber/template v1.6.26
	utilityapp.ktarila/internal/view v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.35.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)

replace utilityapp.ktarila/internal/view => ./internal/view
