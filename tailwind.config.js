module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        transparent: "transparent",
        current: "currentColor",
        googred: "#ea4335",
        googblue: "#4285f4",
        googgreen: "#0f9d58",
        googyellow: "#fbbc04",
      },
      backgroundImage: {
        "nm-down": "linear-gradient(225deg, #e1e2e6, #f5f7fb)",
        "nm-up": "linear-gradient(225deg, #f5f7fb, #e1e2e6);",
      },
      transitionDuration: {
        25: "25ms",
      },
      boxShadow: {
        "nm-sm": "-5px 5px 10px #b9bbbe, 5px -5px 10px #ffffff",
        "nm-sm-inset": "inset -4px 4px 8px #d0d2d5, inset 4px -4px 8px #ffffff",
        "nm-xs": "-3px 3px 5px #b9bbbe, 3px -3px 5px #ffffff;",
      },
    },
    fontFamily: {
      sans: ["Zen Maru Gothic", "sans-serif"],
      serif: ["Zen Antique", "serif"],
    },
  },
  plugins: [],
};
