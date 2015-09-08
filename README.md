# maps

## dependency
* gorm
* github.com/llgcode/draw2d
-- draw2d's dependency has some problem with google code
* github.com/nferruzzi/gormGIS

## map settings
download shape files from
http://data.seoul.go.kr/openinf/mapview.jsp?infId=OA-365
It's Korean Gov. open data.
License is under BY(CC)
Author is Seoul Si(서울시)
### shp2 -> postgis

## go Lang coding
### go install
### revel install
### GoLang code
#### Test code authoring
#### Test URL authoring
##### test URL
  http://localhost:9000/toilet?BBOX=177000%2C437000%2C219000%2C466000&SRS=EPSG:2097&WIDTH=1024&HEIGHT=768
#### Real code
### Leaflet code


### EPSG2097 settings
EPSG
2097	EPSG	2097	PROJCS["Korean 1985 / Korea Central Belt",GEOGCS["Korean 1985",DATUM["Korean_Datum_1985",SPHEROID["Bessel 1841",6377397.155,299.1528128,AUTHORITY["EPSG","7004"]],AUTHORITY["EPSG","6162"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4162"]],UNIT["metre",1,AUTHORITY["EPSG","9001"]],PROJECTION["Transverse_Mercator"],PARAMETER["latitude_of_origin",38],PARAMETER["central_meridian",127],PARAMETER["scale_factor",1],PARAMETER["false_easting",200000],PARAMETER["false_northing",500000],AUTHORITY["EPSG","2097"],AXIS["X",NORTH],AXIS["Y",EAST]]	+proj=tmerc +lat_0=38 +lon_0=127 +k=1 +x_0=200000 +y_0=500000 +ellps=bessel +units=m +no_defs
