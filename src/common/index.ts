import { getIosMarketingName } from '../ios';
import { getAndroidMarketingName } from '../android';

export const getMarketingName = (deviceModel: string, os: string): string => {
  if (os === 'ios') {
    return getIosMarketingName(deviceModel);
  } else if (os === 'android') {
    return getAndroidMarketingName(deviceModel);
  } else {
    return deviceModel;
  }
};
