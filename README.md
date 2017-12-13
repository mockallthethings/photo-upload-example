# Setup

## Install Prerequisites

You will need [docker-compose](https://docs.docker.com/compose/) to run
this tutorial - if you are on a Mac, it should already be installed.
If not, follow the instructions on Docker's website to install it.

## OPTIONAL: Creating a Bucket for this Tutorial

If you do not have an existing S3 bucket to use for this tutorial,
you can create one like this (you will need the [AWS CLI](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)
set up first):

```
aws s3 mb s3://YOUR_BUCKET_NAME
```

# Starting the App

Start the app with this command (must be run from the repository root):

```
docker-compose up -d
```

Then visit the web page at `http://localhost:5000` in your browser.

# Cleaning Up

Stop the app with this command (must be run from the repository root):

```
docker-compose up -d
```

```
docker-compose down
```
