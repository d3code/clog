# CLog: Colorful Console Logging for Go

![Workflow](https://github.com/d3code/clog/actions/workflows/run-tests.yaml/badge.svg)

Welcome to CLog, the go-to logging library for adding a touch of flair and color to your console logs in Go. Say goodbye to dull and monotonous logs and bring life to your terminal with CLog. With its simplicity and elegance, CLog will make your logging experience a delightful one.

## Features

### 1. Colorful Logs 🌈

CLog allows you to add vibrant colors to your logs, making them visually appealing and easy to read. Differentiate between log levels, highlight important information, and add a personal touch to your console output. Your logs will never be the same again!

### 2. Flexible Log Levels ⚙️

With CLog, you have granular control over log levels. Set the desired level and filter out logs that are below the threshold, so you can focus on what matters most. Whether it's debugging, info, warning, or error, CLog has got you covered.

### 3. Custom Formatting ✍️

CLog understands that developers have different preferences when it comes to log formatting. Customize the log format according to your needs. Add timestamps, include or exclude log levels, and tailor the log output to match your unique style.

### 4. Easy Integration 🚀

Integrating CLog into your Go projects is a breeze. With just a few lines of code, you'll be up and running in no time. CLog seamlessly integrates with your existing logging infrastructure, ensuring a smooth transition and hassle-free logging.

## Installation

Installing CLog is as easy as running a simple command:

```shell
go get github.com/d3code/clog
```

## Usage

Using CLog is straightforward. Here's a quick example to get you started:

```go
package main

import (
	"github.com/d3code/clog"
)

func main() {
    clog.Debug("Debug message")
    clog.Info("Info message")
    clog.Warn("Warning message")
    clog.Error("Error message")
}
```

For more advanced usage and customization options, please refer to the [CLog documentation](https://github.com/yourusername/CLog/wiki).

## Log Levels

CLog supports the following log levels, in order of severity:

- Debug
- Info
- Warn
- Error

By default, CLog logs messages of all levels. However, you can set the log level threshold to filter out less important logs.

## Contributing

We welcome contributions to CLog! If you'd like to contribute, please follow these steps:

1. Fork the CLog repository.
2. Create a new branch: `git checkout -b feature/my-awesome-feature`
3. Make your changes and commit them: `git commit -m "Add my awesome feature"`
4. Push your changes to your forked repository: `git push origin feature/my-awesome-feature`
5. Open a pull request in the CLog repository. We'll review your changes and merge them if they align with the project's goals.

## Support

If you have any questions, encounter issues, or want to share your feedback, please reach out to us:

- GitHub Issues: [CLog Issues](https://github.com/d3code/clog/issues)

We're here to help and make your logging experience a colorful one!