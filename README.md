# Setup

## Install Prerequisites

You will need [docker-compose](https://docs.docker.com/compose/) to run
this tutorial - if you are on a Mac, it should already be installed.
If not, follow the instructions on Docker's website to install it.

# Starting the App

Start the app with this command (must be run from the repository root):

```
docker-compose up -d
```

Then visit the web page at `http://localhost:5000` in your browser. If you see
a loading message, wait a few seconds and then refresh the page. It should not
take more than one minute to load.

# Cleaning Up

Stop the app with this command (must be run from the repository root):

```
docker-compose down
```
