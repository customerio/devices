# @customerio/devices

Utility to map device models to marketing friendly device name.

## Installation

```bash
npm i --save @customerio/devices
```

## Usage

```ts
import { getMarketingName, OsType } from '@customerio/devices';

console.log(getMarketingName('iPhone14,2', OsType.IOS));
```
Output:

```bash
iPhone 13 Pro
``` 

## Development

#### How to generate android devices json?

Simple golang script fetches data from https://storage.googleapis.com/play_public/supported_devices.html and generates json file.

```bash
go run 
```

TODO

#### How to generate ios devices json?

- Manual 

- The data is collected from the following sources:
  - Android: https://storage.googleapis.com/play_public/supported_devices.html (Automated)
  - iOS: https://www.theiphonewiki.com/wiki/Models (Manual)
    - Apple TV, Watch, iPads, iPhones

## Size details

[Size Limit](https://github.com/ai/size-limit) controls the size.
Small Libraries

```bash
$ npm run build && npm run size

✔ Adding to empty webpack project
✔ Running JS in headless Chrome

  Size:         419.8 kB with all dependencies, minified and gzipped
  Loading time: 8.2 s    on slow 3G
  Running time: 1.2 s    on Snapdragon 410
  Total time:   9.4 s
  
```

## Todo

- [ ] Protect main branch from direct commits
- [ ] Add a PR template
- [ ] Add development instructions
- [ ] Discuss the need of all 37,708 android devices
- [ ] Automate iOS data collection (far fetched)
