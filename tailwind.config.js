module.exports = {
  content: ["./pkg/web/templates/**/*.templ"],
  theme: {
    extend: {
      colors: {
        ["hc-bg"]: {
          light: "#f6f8fb",
          dark: "#141e2c",
        },
        ["hc-bg-alt"]: {
          light: "#ffffff",
          dark: "#172333",
        },
        ["hc-border"]: {
          light: "#dce1e7",
          dark: "#24384f",
        },
        ["hc-text"]: {
          light: "#172333",
          dark: "#dce1e7",
        },
        ["hc-text-alt"]: {
          light: "#6b7991",
          dark: "#6b7991",
        },
        ["hc-primary"]: "#056ed0",
        ["hc-alert"]: "#d63939",
        ["hc-warning"]: "#d99005",
        ["hc-success"]: "#236b3b",
      },
    },
  },
  plugins: [],
};
