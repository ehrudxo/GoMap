var map
$( document ).ready(function() {
  console.log( "ready!" );
  map= L.map('map').setView([37.519459, 126.983113], 13);
  var osmUrl='http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
  L.tileLayer(osmUrl, {
    attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="http://mapbox.com">Mapbox</a>',
    maxZoom: 18
  }).addTo(map);
  var toilet = L.tileLayer.wms("http://localhost:9000/toilet", {
    layers: 'toilet',
    format: 'image/png',
    transparent: true,
      attribution: "public toilet"
  });
  map.addLayer(toilet);
});
