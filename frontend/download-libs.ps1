# Download CSS & JS lib

## CSS libs
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" -OutFile "wwwroot/css/bootstrap.min.css"
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css" -OutFile "wwwroot/css/bootstrap-icons.min.css"

## Fonts
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/twbs/icons/main/font/fonts/bootstrap-icons.woff2" -OutFile "wwwroot/css/fonts/bootstrap-icons.woff2"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/twbs/icons/main/font/fonts/bootstrap-icons.woff" -OutFile "wwwroot/css/fonts/bootstrap-icons.woff"

## JS libs
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" -OutFile "wwwroot/js/bootstrap.bundle.min.js"
Invoke-WebRequest -Uri "https://cdnjs.cloudflare.com/ajax/libs/Chart.js/4.0.1/chart.umd.js" -OutFile "wwwroot/js/chart.umd.js"
Invoke-WebRequest -Uri "https://cdnjs.cloudflare.com/ajax/libs/chartjs-plugin-datalabels/2.2.0/chartjs-plugin-datalabels.min.js" -OutFile "wwwroot/js/chartjs-plugin-datalabels.min.js"
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js" -OutFile "wwwroot/js/Sortable.min.js"