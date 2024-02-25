# LinkVigil

## Description

LinkVigil is a simple tool to monitor the status of API endpoints and posts their status to Atlassian Statuspage via its API.

## Setup

### Environment Variables

STATUSPAGE_API_KEY - The API key for the Statuspage API.

### Settings File

Please see `settings.yaml.example` for an example of the settings file. The settings file should be named `settings.yaml` and should be placed in the root directory of the project.

## Metrics and Thresholds

* **Response Time**: The time it takes for the API to respond to a request. The threshold is 500ms.
* **Status Code**: The status code returned by the API. The threshold is 200.
