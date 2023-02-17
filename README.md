<h1>Go Timezone CLI</h1>

<p><strong>A CLI Tool for timezones made using golang :zap:</strong></p>

## What can `go-timezones-cli` do? :sparkles:
- Search for date and time based on country, or timezones.
- Manage list for timezones you frequently view.
- Get UTC date and time based on your local timezone or any timezones.
- Convert time across timezones

## Credits
- Special thanks to @MitaliBhalla for the idea & feedback on the project. 
- Also thanks to @supreeth7 for providing code review & feedback on the project

## Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the CLI](#running-the-cli)
- [Commands](#commands)
  - [Search for local date time](#search-for-local-date-time)
  - [Search based on timezone abbreviations](#search-based-on-timezone-abbreviations)
  - [Add timezones](#add-timezones)
  - [Remove timezones](#remove-timezones)
  - [Show & Convert between timezones](#show-and-convert-between-timezones)
  - [Select a single timezone from defaults](#select-a-single-timezone-from-defaults)
  - [Get UTC time](#get-utc-time)
  - [Use Interactive UI](#get-interactive-ui)
- [Future Work](#future-work)
- [Contributing and Feedback](#contribution)


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
> **NOTE:** [List of country codes or timezone names](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) :earth_asia:

## Commands

### Search for local date time

Search can be done via country code or country name or city name that is included in the timezone list

#### Usage 
```
tmz search <search-term>
```


```shell
$ tmz search "us"

$ tmz search "kol"
```
<details><summary><strong>Demo</strong></summary>

<img src = "demos/search.gif" width="700" alt="demo of search search" />
</details>

---
### Search based on timezone abbreviations
To search for a timezone based on abbreviation its abbreviations e.g IST

### Usage
```
tmz get <abbreviation> [time]
```

```bash
$ tmz get "ist"

$ tmz get "est"
```
To get the converted time at the zone provide the local time as one of the arguments

```bash
$ tmz get "jst" 10:56
```
<details><summary><strong>Demo</strong></summary>
<img src = "demos/get.gif" width="700" alt="demo of search search" />
</details>

---

### Add timezones

Timezones added to the config file are treated as the default timezones which is triggered by the `tmz show --all` command.

If the timezone already exists, the CLI throws an error.

> config file is stored at ~/.tmz.list
#### Usage 
```
tmz add <timezone>
```

```bash
$ tmz add "Asia/Kolkata"
```
<details><summary><strong>Demo</strong></summary>
<img src = "demos/add.gif" width="700" alt="demo of search search" />
</details>

---

### Remove timezones

Timezones are removed from the config file
#### Usage 
```
tmz remove <timezone>
```

```bash
$ tmz remove "Asia/Kolkata"
```
<details><summary><strong>Demo</strong></summary>
<img src = "demos/remove.gif" width="700" alt="demo of search search" />
</details>

---

### Show and Convert between timezones

- This comand will help you to convert between timezones. 
- The first timezone passed as argument is set as the default timezone. Other time are calculated based on it. 

#### Usage
```
tmz show <timezone1> [timezone2, timezone3 ...] [--all] [--time time]

Flags:
--all           Load all Timezones from local config
-t, --time string   Show the list at a different time
```
- The `--time` flag will accept an time which will be used to show the list based on that time
- The `--all` flag will load all the timezones saved on the config file

```bash
$ tmz show --all

$ tmz show "Asia/Jakarta" -t=10:56
```
---

### Select a single timezone from defaults

To get a list of all local timezones that are saved in the config file. 
#### Usage
```
tmz select
```
On selection it will provide an option to show the current time or convert the time based on provided local time

<details><summary><strong>Demo</strong></summary>
<img src = "demos/select.gif" width="700" alt="demo of search search" />
</details>

---

### Get UTC time

Get UTC time based on current system time.
#### Usage
```
tmz utc
```
<details><summary><strong>Demo</strong></summary>
<img src = "demos/utc.gif" width="700" alt="demo of search search" />
</details>

---
### Get Interactive UI

Get an Interactive UI to see time across timezones. 
#### Usage
```
tmz visual [timezone1, timezone2, ...] [--all]

Flags:
--all    Load all Timezones from locally saved timezones
```

```bash
$ tmz visual --all
$ tmz visual "Asia/Kolkata" "Europe/Berlin"
```
<details><summary><strong>Demo</strong></summary>
<img src = "demos/visual.gif" width="700" alt="demo of search search" />
</details>

---
## Future Work
- Redefined search
- Time Format Configuration 
- Work with Daylight Savings

## Feedback and Contribution
The project is open for feedback & contribution from all talented people across the globe. Feel free to reach out to me for any suggestions or improvemnts. 

If there is an issue do open up a PR for the same