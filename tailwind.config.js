/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./**/*.{templ,go,html}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui'),
  ],
}
