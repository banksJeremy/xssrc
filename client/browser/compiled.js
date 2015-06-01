"use strict";

navigator.serviceWorker.register("compiled.js").then(function (registration) {
  console.log("Registered service worker", registration);
}, function (err) {
  console.log("Failed to register service worker", err);
});
