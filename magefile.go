//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/riweston/gotempl-cv/views/layout"
)

// BuildCSS builds the CSS.
func BuildCSS(inputFileName string, outputFileName string) error {
	mg.Deps(InstallNodeDeps)
	if err := sh.Run("npx", "tailwindcss",
		"-i", "assets/css/"+inputFileName,
		"-o", "public/css/"+outputFileName); err != nil {
		log.Error("Failed to build CSS")
		return err
	}
	return nil
}

// BuildSiteCSS builds the CSS for the site.
func BuildSiteCSS() error {
	return BuildCSS("tailwind.css", "style.css")
}

// BuildPrintableCSS builds the CSS for the printable CV.
func BuildPrintableCSS() error {
	return BuildCSS("tailwind.pdf.css", "style.pdf.css")
}

// InstallDeps installs the dependencies.
func InstallNodeDeps() error {
	log.Info("Installing Node dependencies")
	if err := sh.Run("npm", "install"); err != nil {
		log.Error("Failed to install Node dependencies")
		return err
	}
	log.Info("Installed Node dependencies")
	return nil
}

// GenerateTempl generates Go code from the templates.
func GenerateTempl() error {
	log.Info("Generating templ")
	if err := sh.Run("templ", "generate"); err != nil {
		log.Error("Failed to generate templ")
		return err
	}
	log.Info("Generated templ")
	return nil
}

// BuildPrintableCV builds the pdf view of the CV.
func BuildPrintableCVAsHTML() error {
	mg.Deps(GenerateTempl, BuildPrintableCSS)
	log.Info("Building 1 page CV as HTML")
	f, err := os.Create("pdf.html")
	if err != nil {
		log.Error("Failed to create pdf.html file")
		return err
	}
	defer f.Close()

	if err := layout.PDFLayout().Render(context.Background(), f); err != nil {
		log.Error("Failed to render pdf.html")
		return err
	}
	log.Info("Built 1 page CV as HTML")

	return nil
}

// ConvertHTMLToPDF converts an HTML file to a PDF.
func ConvertHTMLToPDF(inputFileName string, outputFileName string) error {
	log.Info(fmt.Sprintf("Converting %s to PDF", inputFileName))
	if err := sh.Run("node", "html-to-pdf.js", inputFileName, outputFileName); err != nil {
		log.Error(fmt.Sprintf("Failed to convert %s to PDF", inputFileName))
		return err
	}
	log.Info(fmt.Sprintf("Converted %s to %s", inputFileName, outputFileName))
	return nil
}

// GenerateCVAsPDF generates the CV as a PDF.
func GenerateCVAsPDF() error {
	mg.Deps(BuildPrintableCVAsHTML)
	log.Info("Generating CV as PDF")
	if err := ConvertHTMLToPDF("pdf.html", "cv.pdf"); err != nil {
		log.Error("Failed to generate CV as PDF")
		return err
	}

	log.Info("Generated CV as PDF")
	return nil
}
