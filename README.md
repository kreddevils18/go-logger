# **Go Logger**

A powerful and extensible logging library for Go, designed for high-performance and flexibility. Go Logger supports multiple logging strategies, including seamless integration with Fluentd for centralized log aggregation.

---

## **Table of Contents**

- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Log Strategies](#log-strategies)
  - [StandardLogStrategy](#standardlogstrategy)
  - [FluentLogStrategy](#fluentlogstrategy)
- [Configuration](#configuration)
- [License](#license)

---

## **Features**

- **Fluentd Integration:** Send logs to Fluentd for centralized storage and processing.
- **Strategy Pattern:** Easily switch between logging strategies (e.g., standard logging, Fluentd).
- **High Performance:** Built on Uber's [zap](https://github.com/uber-go/zap) for fast, structured logging.
- **Environment Awareness:** Adapts to production or development environments.
- **Thread-Safe:** Safe for use in multi-threaded applications.
- **Extensibility:** Create custom logging strategies by implementing the `LogStrategy` interface.

---

## **Getting Started**

### **Installation**

To install the library and its dependencies, run:

```bash
go get -u github.com/kreddevsil18/go-logger
go get -u github.com/fluent/fluent-logger-golang
go get -u go.uber.org/zap
```

## **Log strategies**
### StandardLogStrategy
The `StandardLogStrategy` provides a simple logging mechanism based on zap. It supports JSON-encoded or human-readable logs, configurable through environment settings.

### FluentLogStrategy
The FluentLogStrategy integrates with Fluentd to send structured logs for centralized log aggregation.

#### Features
- Sends logs to Fluentd in structured JSON format.
- Dynamically configurable log levels and output destinations.
- Production-ready with minimal setup.

## Configuration
The Config struct defines global configurations for all logging strategies

## License
This project is licensed under the MIT License. See the LICENSE file for details.
