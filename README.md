# devices

Utility to map device models to marketing friendly device name.

## Installation

```bash
npm i --save @customerio/devices
```

## Usage

```ts
import { getMarketingName } from '@customerio/devices';

console.log(getMarketingName('iPhone14,2', 'ios'));
```
Output:

```bash
iPhone 13 Pro
``` 



## Development

- The data is collected from the following sources:
  - Android: https://storage.googleapis.com/play_public/supported_devices.html (Automated)
  - iOS: https://www.theiphonewiki.com/wiki/Models (Manual)
    - Apple TV, Watch, iPads, iPhones

## BundlePhobia

https://bundlephobia.com/package/@customerio/devices
