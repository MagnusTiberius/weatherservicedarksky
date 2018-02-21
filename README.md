# weatherservicedarksky

**Build Status**

[![CircleCI](https://circleci.com/gh/MagnusTiberius/weatherservicedarksky.svg?style=svg)](https://circleci.com/gh/MagnusTiberius/weatherservicedarksky)

Please see .circleci/config.yml for details about the build.

**Design**
1. Receive address data
2. Pass address data to Google Geo and receive a lat/lng information.
3. Call DarkSky and pass lat/lng and(or) datetime data.
4. Receive a json string and pass this back to caller in json format.

**Docker Container**

Each build will create a docker container which is packaged and sent to a Google Cloud project repo.


**Google Cloud**
1. A config.yml script sets up a gcloud environment during build which then creates and pushes the docker container to the repo.
2. A create pod command is entered in a cluster console
   ```
   kubectl run api-darksky --image=us.gcr.io/weatherservice-195512/apidarksky --port 8090
   ```
3. A create service command is then used to expose it.
   ```
   kubectl expose deployment api-darksky --type=LoadBalancer --port 8090 --target-port 8090
   ```


**Dependent Microservices**

1. https://github.com/MagnusTiberius/weatherservice
