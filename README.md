# Country search

## Settinup Locally

### Using docker. Docker must be installed.
The initial plan was to use the multistage docker build to build the react frontend and golang backend. Then copy the frontend's static build and serve it under golang backend.
Docker file is kinda complete but a new api endpoint needs to be created in the backend that hosts the static build folder.
### running the backend
- Clone the repo
- cd into the country-search-react-golang folder. There should be  dockerfile.dev
- type "docker build -t country-search -f Dockerfile.dev ." in the terminal to build the docker file.
- type "docker run --rm -p 2020:2020 -it country-search" to run the docker container
- If everything goes fine the url http://localhost:2020 should be up. 

### running the frontend Node should be installed
- cd into country-search-react-golang.
- type "yarn install" and wait for dependencies to install
- type "yarn start"
- http://localhost:3000 should be up and running. That uses the http://localhost:2020/search get endpoint to load the filtered country.

## Features
- When the length of searchstring (i.e http://localhost:2020/search?text=searchString) is less than 2 characters it should list all the countries and the matching algorithm is not performed.

- If the searchstring is 2 or more characters long then a full search of nested elements is performed and the countries with the matching criteria is returned.

- The search is performed on the master json file /backend/datasource/countries.json. Search is also performed on most of the nested elements. 

## What's left
- When clicked on the list of countries there should be a page where all the information available about a single country is shown.