import { getIosDevices } from '../ios';
import { getAndroidDevices } from '../android';
import { Platform } from '../types';

/**
 * Get marketing name for a given device model or identifier.
 * If no device is found, returns the model or identifier.
 *
 * @param deviceModel - case insensitive
 * @param os
 */
export const getMarketingName = (deviceModel: string, os: Platform): string => {
  if (os === Platform.IOS) {
    const devices = getIosDevices(deviceModel);
    if (devices.length > 0) {
      return devices[0].generation;
    }
  } else {
    const devices = getAndroidDevices(deviceModel);
    if (devices.length > 0) {
      return devices[0].marketingName;
    }
  }
  return deviceModel;
};
