# Recruitment project - Webscraper

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
    - [Usage](#usage)
    - [Test](#test)

## Introduction

This tool is a website word frequency analyzer built using Go (Golang). It leverages goroutines, channels, and the standard Go library to combine text from multiple websites and generate a list of the most frequently appearing words.

## Features

- Fetch text content from multiple websites.
- Analyze and generate word frequency statistics.
- Concurrent using goroutines and channels.
- Basic cache mechanism.
- Limit the number of concurrently running goroutines.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Go (Golang)](https://golang.org/doc/install)
- Optional: [Docker](https://www.docker.com/)
## Getting Started

### Usage

1. Clone this repository & change into the project directory:
   ```shell
   git clone https://github.com/raj3k/webscraper.git
   cd webscraper
   ```
2. Run project using **Makefile**:
    ```shell
   make run
    ```
3. Run project using **Docker**:
    ```shell
   docker build -t webscraper .
   docker build -e URLS="https://example.com/,https://example2.com/" webscraper
    ```
   or
    ```shell
   docker build -t webscraper .
   docker build webscraper
    ```

### Test
1. Clone this repository & change into the project directory:
   ```shell
   git clone https://github.com/raj3k/webscraper.git
   cd webscraper
   ```
2. Test project using **Makefile**:
    ```shell
   make test
    ```