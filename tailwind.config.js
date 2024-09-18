const { addIconSelectors } = require("@iconify/tailwind");

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./**/*.{templ,go,html}"],
  // safelist: ["a4-container", "content-wrapper", "left-section", "main-section"],
  theme: {
    extend: {
      width: {
        "a4-width": "210mm",
      },
      height: {
        "a4-height": "297mm",
      },
    },
  },
  plugins: [
    addIconSelectors(["mdi", "mdi-light", "simple-icons", "logos"]),
    require("@tailwindcss/typography"),
  ],
};
