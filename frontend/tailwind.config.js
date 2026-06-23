/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      borderRadius: {
        'xl': '20px',
      },
      backdropBlur: {
        'glass': '20px',
        'glass-lg': '24px',
      }
    },
  },
  plugins: [],
}