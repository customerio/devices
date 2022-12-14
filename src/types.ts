export interface AndroidDevice {
  retailBranding: string;
  device: string;
  model: string;
  marketingName: string;
}

export interface IosDevice {
  identifier: string;
  generation: string;
}

export enum Platform {
  IOS = 'ios',
  ANDROID = 'android',
}
