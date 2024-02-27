# LinkVigil

## Description

LinkVigil is a simple tool to monitor the status of API endpoints and posts their status to Atlassian Statuspage via its API.

## Setup

### Environment Variables

STATUSPAGE_API_KEY - The API key for the Statuspage API.

### Settings File

Please see `endpoints.example.yaml` for an example of the settings file.

## Building

```bash
go build .
```

## Usage

Run the binary with the path to endpoints yaml file as argument, and "debug" as optional argument.
Example:

```bash
./linkvigil /path/to/endpoints.yaml
```

(Debug mode, i.e report only, no API call to Statuspage)

```bash
./linkvigil /path/to/endpoints.yaml debug
```

## Metrics and Thresholds

| Event    | Action   |
| -------- | -------- |
| Endpoint Respond Status 200 | Post Status *Operational* |
| Endpoint Response Time Above 500ms | Post Status *Degraded* |
| Endpoint Respond Status 500 | Post Status *Major Outage* |
| Endpoint Not Responding | Post Status *Major Outage* |
| Any Other Response Status | Info Log Endpoint And HTTP Status |
