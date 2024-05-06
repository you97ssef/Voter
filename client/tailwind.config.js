/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,ts}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["abeezee", "sans-serif"],
        lobster: ["lobster-two", "sans-serif"],
        courgette: ["courgette", "sans-serif"],
      },
    },
  },
  daisyui: {
    themes: [
      {
        darko: {
          primary: "#66347F",
          "primary-content": "#0C060E",
          secondary: "#0D7377",
          "secondary-content": "#021212",
          accent: "#37306B",
          "accent-content": "#07060e",
          neutral: "#181611",
          "base-100": "#2B283E",
          "base-200": "#222032",
          "base-300": "#191825",
          info: "#3876BF",
          "info-content": "#0D1929",
          success: "#00f08c",
          "success-content": "#002918",
          warning: "#ffa600",
          "warning-content": "#3D2800",
          error: "#ff0037",
          "error-content": "#290008",

          "--rounded-box": "1.5rem",
          "--rounded-btn": "1rem",
          "--rounded-badge": "2rem",

          "--animation-btn": ".25s",
          "--animation-input": ".2s",

          "--btn-text-case": "uppercase",
          "--navbar-padding": "1rem",
          "--border-btn": "2px",
        },
      },
      {
        lighto: {
          primary: "#66347F",
          "primary-content": "#EDE2F3",
          secondary: "#0D7377",
          "secondary-content": "#DAF9FB",
          accent: "#37306B",
          "accent-content": "#E5E2F3",
          neutral: "#F7F6F3",
          "base-100": "#F3F3F7",
          "base-200": "#E7E7EF",
          "base-300": "#DADAE7",
          info: "#3876BF",
          "info-content": "#EFF4FA",
          success: "#00f08c",
          "success-content": "#EBFFF7",
          warning: "#ffa600",
          "warning-content": "#FFF8EB",
          error: "#ff0037",
          "error-content": "#FFEBEF",

          "--rounded-box": "1.5rem",
          "--rounded-btn": "1rem",
          "--rounded-badge": "2rem",

          "--animation-btn": ".25s",
          "--animation-input": ".2s",

          "--btn-text-case": "uppercase",
          "--navbar-padding": "1rem",
          "--border-btn": "2px",
        },
      },
    ],
  },
  plugins: [
    require('daisyui'),
  ],
}

