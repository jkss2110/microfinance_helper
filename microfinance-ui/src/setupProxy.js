const proxy = require("http-proxy-middleware");

module.exports = app => {
  app.use(
    "/loandetails",
    proxy({
      target: "http://localhost:5000",
      changeOrigin: true,
      "secure": false
    })
  );
};