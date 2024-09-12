import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    backgroundImage: {
      'background_estacionamento' : "url('../src/assets/images/background_estacionamento.jpeg')"
    },
    extend: {
      colors: {'primary-blue' : '#106E80',},
    },
  },
  plugins: [],
};
export default config;
