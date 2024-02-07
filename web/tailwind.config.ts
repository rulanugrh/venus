import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        darkColor: "#1F2544",
        semiDark: "#474F7A",
        semiLight: "#81689D",
        lightColor: "#FFD0EC"
      },

      container: {
        center: true,
        padding: "16px"
      },

      fontFamily: {
        jetBrains: "'JetBrains Mono', monospace"
      }
    },
  },
  plugins: [],
};
export default config;
