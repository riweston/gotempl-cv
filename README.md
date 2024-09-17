# README

## What is this?

This is my online CV, open sourced here on GitHub for anyone to see. It's a work in progress, and I'm always looking for ways to improve it.

This is based on the GoTH stack (minus HTMX as I'm not using it here).

## Why did I make this?

I wanted to have a place where I could showcase my achievements and skills in a way that was more interactive than a traditional CV. I also wanted to learn how to use the GoTH stack, and this seemed like a good way to do it. All credit to the current design and layout goes to https://github.com/KonradSzwarc/devscard - I've made some changes, but the core design is his.

## How can I use this?

You can fork this repo and use it as a template for your own CV. You can also use it as a learning resource to see how the GoTH stack works.

All the data used to generate the CV is stored the yaml files in the `./public/data` directory. You can edit these files to add your own data.

> [!NOTE]
> Please make sure you adhere to the types defined in `./pkgs/types` when adding data to the yaml files. This is to ensure that the data is displayed correctly on the CV.

### Running the server locally

For local development, you can run either `go run main.go` or you can use the `air` tool to run the server. You can install `air` by running `go install github.com/air-verse/air@latest`. Once you have `air` installed, you can run `air` in the root directory of the project to start the server.

> [!NOTE]
> This project contains several `entrypoint_*.go` files. These files are used for different build contexts using Go build tags. The `main.go` file is used for the default build context. If you want to run the server in a different build context, you can use the `entrypoint_<context>.go` file.
>
> - The `dev` build context you can run `go run -tags dev main.go` or `air` in the root directory of the project.
> - The `prod` build context you can run `go run main.go`, this will run a live server with the production build which is handy for specific configuration if deploying to Cloud Run or similar.
> - The `static` build context you can run `go run -tags test main.go`, this will vend out a rendered html page to the root directory of the project.

### Updating the certificate data

The certificate data is stored in the `certificates_gen.yml` file in the `./public/data` directory. You can add your own certificates to this file to display them on the CV or you can run the helper command to generate the certificates data for you including the certificate images. To do this, run `go run ./cmd/generate-certs/main.go <credly username>` in the root directory of the project. This will generate the certificates data for you from the credly API.

### PDF generation

This project includes a separate view for a PDF version of the CV better suited for printing. You can access this view by going to `/pdf` on the server. This uses it's own tailwindcss file (`tailwind.pdf.css`) to ensure that the layout is correct for printing.

## How can I contribute?

If you see any issues with the code, or have any suggestions for improvements, feel free to open an issue or a pull request.
