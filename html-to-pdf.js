const puppeteer = require("puppeteer");
const path = require("path");

console.log("Starting the PDF generation process...");

// Function to get the input and output filenames from command-line arguments
function getFileNames() {
  const args = process.argv.slice(2); // Remove the first two elements (node and script name)
  if (args.length < 1) {
    console.error("Usage: node script.js <input_html_file> [output_pdf_file]");
    process.exit(1);
  }
  const inputFile = args[0];
  const outputFile = args[1] || inputFile.replace(/\.html$/, ".pdf");
  return { inputFile, outputFile };
}

(async () => {
  try {
    const { inputFile, outputFile } = getFileNames();
    console.log(`Input file: ${inputFile}`);
    console.log(`Output file: ${outputFile}`);

    console.log("Launching browser...");
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    // Get the current directory
    console.log("Current directory:", process.cwd());
    const currentDir = process.cwd();

    // Load the HTML file
    console.log("Loading HTML file...");
    await page.goto(`file:${path.join(currentDir, inputFile)}`, {
      waitUntil: "networkidle0",
    });

    // Generate PDF
    console.log("Generating PDF...");
    await page.pdf({
      path: outputFile,
      format: "A4",
      printBackground: true,
      margin: {
        top: "20mm",
        right: "20mm",
        bottom: "20mm",
        left: "20mm",
      },
    });

    await browser.close();
    console.log(`PDF has been generated successfully: ${outputFile}`);
  } catch (error) {
    console.error("An error occurred:", error);
    process.exit(1);
  }
})();
