const esbuild = require("esbuild");

esbuild
  .serve(
    {
      servedir: "public",
      port: 8000,
    },
    {
      entryPoints: ["./app.jsx"],
      outfile: "./public/bundle.js",
      bundle: true,
      loader: {
        ".js": "jsx",
      },
      plugins: [],
    }
  )
  .catch(() => process.exit());