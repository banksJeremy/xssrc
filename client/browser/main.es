navigator.serviceWorker.register('compiled.js').then(registration => {
  console.log("Registered service worker", registration);
}, err => {
  console.log("Failed to register service worker", err);
});

