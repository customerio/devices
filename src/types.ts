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

export enum OsType {
  IOS = 'ios',
  ANDROID = 'android',
}
