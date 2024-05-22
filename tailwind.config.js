const { addIconSelectors } = require('@iconify/tailwind');

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: [
    "./**/*.{templ,go,html}",
  ],
  theme: {
    screens: {
      'xs': '475px',
    },
    extend: {},
  },
  plugins: [
    addIconSelectors(['mdi', 'mdi-light', 'simple-icons', 'logos']),
    require("@tailwindcss/typography"),
    require('daisyui'),
  ],
}
