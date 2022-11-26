/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,ts,jsx,tsx}"
  ],
  theme: {
    extend: {
      animation: {
        'move':' moveForever 25s cubic-bezier(.55,.5,.45,.5) infinite'
      },
      keyframes: {
        moveForever: {
          '0%': {
           transform: 'translate3d(-90px,0,0)'
          },
          '100%': {
            transform: 'translate3d(85px,0,0)'
          }
        }
      }
    },
    animationDelay: {
      1: "-2000ms",
      2: "-3000ms",
      3: "-4000ms",
      4: "-5000ms",
    },
  },
  plugins: [],
}
