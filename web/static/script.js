function getStatus() {
  fetch("/weather")
    .then((response) => response.json())
    .then((data) => {
      document.getElementById("water").innerText = data.status.water;
      document.getElementById("wind").innerText = data.status.wind;

      // Determine water status
      let waterStatus = "";
      if (data.status.water === "aman") {
        waterStatus = "safe";
      } else if (data.status.water === "siaga") {
        waterStatus = "alert";
      } else {
        waterStatus = "alert";
      }
      document.getElementById("water-status").className = waterStatus;

      // Determine wind status
      let windStatus = "";
      if (data.status.wind === "aman") {
        windStatus = "safe";
      } else if (data.status.wind === "siaga") {
        windStatus = "alert";
      } else {
        windStatus = "alert";
      }
      document.getElementById("wind-status").className = windStatus;
    })
    .catch((error) => {
      console.error("Error fetching data:", error);
    });
}

// Auto reload status every 15 seconds
setInterval(getStatus, 15000);

// Initial load of status
getStatus();
