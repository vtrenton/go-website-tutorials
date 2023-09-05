module github.com/vtrenton/go-website-tutorials/create-a-module/hello

go 1.21.0

replace github.com/vtrenton/go-website-tutorials/create-a-module/greetings => ../greetings

require github.com/vtrenton/go-website-tutorials/create-a-module/greetings v0.0.0-00010101000000-000000000000
