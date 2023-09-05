module github.com/vtrenton/go-website-tutorials/create-a-module/hello

go 1.21.0

// this will pull from the local module path instead of github since my local code will be more up-to-date than github
replace github.com/vtrenton/go-website-tutorials/create-a-module/greetings => ../greetings

require github.com/vtrenton/go-website-tutorials/create-a-module/greetings v0.0.0-00010101000000-000000000000
