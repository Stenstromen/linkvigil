# LinkVigil

![Logo](logo.webp)

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

## Linux Systemd Service

1. Download the binary from the latest release.
2. Create a settings file and place it in `/etc/linkvigil.yaml`.
3. Create a systemd service file in `/etc/systemd/system/linkvigil.service`:

```ini
[Unit]
Description=LinkVigil
Wants=network-online.target
After=network-online.target
AssertFileIsExecutable=/usr/local/bin/linkvigil

[Service]
WorkingDirectory=/usr/local/
User=linkvigil
Group=linkvigil
EnvironmentFile=/etc/default/linkvigil
ExecStart=/usr/local/bin/linkvigil /etc/linkvigil.yaml 2>&1 | logger -t linkvigil
Restart=always
TasksMax=infinity
TimeoutStopSec=infinity
SendSIGKILL=no

[Install]
WantedBy=multi-user.target
```

4. Create a user and group for the service:

```bash
sudo useradd -r -s /bin/false linkvigil
```

5. Change ownership of binary:

```bash
chown linkvigil: /usr/local/bin/linkvigil
```

6. Add the API key to `/etc/default/linkvigil`:

```bash
STATUSPAGE_API_KEY=your-api-key
```

7. Enable and start the service:

```bash
sudo systemctl enable linkvigil --now
```

## Metrics and Thresholds

| Event                              | Action                            |
| ---------------------------------- | --------------------------------- |
| Endpoint Respond Status 200        | Post Status _Operational_         |
| Endpoint Response Time Above 500ms | Post Status _Degraded_            |
| Endpoint Respond Status 500        | Post Status _Major Outage_        |
| Endpoint Not Responding            | Post Status _Major Outage_        |
| Any Other Response Status          | Info Log Endpoint And HTTP Status |
