# Web fetcher
## Assumption and Notes
* the uri used does not contain query or path, so it will be google.com, not google.com/something?a=b
* the uri is full path includes scheme,like http,https
* the links printed out is the anchor count
* the images printed out is the img tag count
* the assets downloaded include the img, anchor, js, css
* to change the requested uri, change the dockerfile cmd
* there are many libraries already does it, but using that defeats the purpose, so I chose to use the most fundamental packages
