<img src="https://i.imgur.com/ZebplfT.png" width="110" align="left"/><h1>Go Timezone CLI</h1>

<p><strong>A CLI Tool for timezones :zap:</strong></p>

## What can `go-timezones-cli` do? :sparkles:
- Search for date and time based on country, or timezones.
- Manage list for timezones you frequently view.
- Get UTC date and time based on your local timezone or any timezones.
- Convert time across timezones


## Contents

- [Installation](#installation)
- [Commands](#commands)
  - [Search for local date time](#search-for-local-date-time)
  - [Search based on timezone abbreviations](#search-based-on-timezone-abbreviations)
  - [Add timezones](#add-timezones)
  - [Remove timezones](#remove-timezones)
  - [Show local datetime of all saved timezones](#show-local-datetime-of-all-saved-timezones)
  - [Select a single timezone from defaults](#select-a-single-timezone-from-defaults)
  - [Get UTC time](#get-utc-time)


## Prerequisites

You will need to `Go >= 1.18` installed on your computer

## Installation
First clone the repository somewhere in your $PATH. A common place would be within your $GOPATH. <br>

### Option 1: Install binary

Build and copy `tmz` to your $GOPATH/bin:

```
$ make install
```
### Option 2: Build from source

This command will build the Timezone CLI binary, named `tmz`. This binary will be created in the root of your project folder.

```
$ make build
```

## Running the CLI
To list all commands:
```shell
$ tmz help
```

To run command with the full name:
```shell
$ tmz <command> [flags] <string>
```
> **NOTE:** [List of country codes or timezone names](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) :earth:

## Commands

### Search for local date time

Search can be done via country code or country name or city name that is included in the timezone list


```shell
$ tmz search "us"
```
---
### Search based on timezone abbreviations

```bash
$ tmz get "ist"

$ tmz get "est"
```

---

### Add timezones

Timezones added to the config file are treated as the default timezones which is triggered by the `tmz show` command.

> file is stored at ~/.tmz.list

```bash
$ tmz add "Asia/Kolkata"
```
---

### Remove timezones

Timezones are removed from the config file
```bash
$ tmz remove "Asia/Kolkata"
```
---

### Show local datetime of all saved timezones

```bash
$ tmz show
```

To get the time based on a different local time use
```
tmz show <time>

$ tmz show "21:45"
```

---

### Select a single timezone from defaults

```bash
$ tmz select
```
---

### Get UTC time

Get UTC time based on current system time.
```bash
$ tmz utc
```

## Planned Features, Updates & Future Work
- Search for timezones & display local time
- Manage list of timezones
- UTC & Local Time Conversion
- Time Conversion from TimeZone to Local Time
- Interactive Time View Across Time Zones
- Redefined search
- Time Format Configuration 
- Support for 24 Hrs instead of 12 Hrs
- Code Modularity & Improvements