# HealthCheckService

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/healthcheckservice)](https://goreportcard.com/report/github.com/yourusername/healthcheckservice)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

HealthCheckService is a lightweight, production‑ready Go service that monitors the health of a configurable list of HTTP endpoints. 
It periodically checks each target, stores the latest status, and exposes a JSON API (`/health`) that reports the current health of all monitored services.

Features:

- **Configurable** via environment variables (`.env` file support)
- Periodic health checks with configurable interval
- Structured logging with Logrus
- Simple HTTP API powered by Gorilla Mux
- Graceful shutdown handling
- Unit tests covering core logic

## Getting Started

### Prerequisites

- Go 1.22 or newer
- Git (to clone the repository)

### Installation

<!-- tiny readability tweak -->
